package usecase

import (
	"context"
	"fmt"
	"time"

	"monorepo/services/library-service/internal/modules/lending/domain"
	"monorepo/services/library-service/pkg/helper"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"monorepo/services/library-service/pkg/shared/domain/constant"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (uc *lendingUsecaseImpl) ReturnLending(ctx context.Context, id int) (result domain.ReturnLending, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "LendingUsecase:UpdateLending")
	defer trace.Finish()

	repoFilter := shareddomain.LendingParamGet{ID: &id}
	existing, err := uc.repoSQL.LendingRepo().Find(ctx, &repoFilter)
	if err != nil {
		return
	}

	now := time.Now()
	existing.ReturnDate = &now

	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) (err error) {
		err = <-uc.repoSQL.LendingRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("ReturnDate"))
		if err != nil {
			return
		}

		truncNow := helper.TruncateToDate(time.Now())
		truncdueDate := helper.TruncateToDate(existing.DueDate)
		if !truncNow.After(truncdueDate) {
			return nil
		}

		duration := truncNow.Sub(truncdueDate)
		daysLate := int(duration.Hours() / 24)
		amount := daysLate * 1000

		// 1. Initiate Snap client
		var s = snap.Client{}
		s.New("SB-Mid-server-WmTgfO1rHzGtcbt-oZc0ACaL", midtrans.Sandbox)

		// 2. Initiate Snap request
		req := &snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  fmt.Sprintf("%s%v", constant.PREFIX_LENDING, existing.ID),
				GrossAmt: int64(amount),
			},
			CreditCard: &snap.CreditCardDetails{
				Secure: true,
			},
		}

		// 3. Request create Snap transaction to Midtrans
		snapResp, midErr := s.CreateTransaction(req)
		if midErr != nil {
			return midErr.RawError
		}

		fine := shareddomain.Fine{LendingID: existing.ID, Amount: amount, SnanpURL: snapResp.RedirectURL, PaymentStatus: constant.Pending}
		err = uc.repoSQL.FineRepo().Save(ctx, &fine)
		if err != nil {
			return
		}

		result.Serialize(&fine)

		return nil
	})
	return
}
