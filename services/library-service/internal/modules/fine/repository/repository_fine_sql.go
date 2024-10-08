// Code generated by candi v1.17.15.

package repository

import (
	"context"

	"strings"
	"time"

	"monorepo/services/library-service/internal/modules/fine/domain"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"

	"monorepo/globalshared"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type fineRepoSQL struct {
	readDB, writeDB *gorm.DB
	updateTools     *candishared.DBUpdateTools
}

// NewFineRepoSQL mongo repo constructor
func NewFineRepoSQL(readDB, writeDB *gorm.DB) FineRepository {
	return &fineRepoSQL{
		readDB: readDB, writeDB: writeDB,
		updateTools: &candishared.DBUpdateTools{
			KeyExtractorFunc: candishared.DBUpdateGORMExtractorKey, IgnoredFields: []string{"id"},
		},
	}
}

func (r *fineRepoSQL) FetchAll(ctx context.Context, filter *domain.FilterFine) (data []shareddomain.Fine, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineRepoSQL:FetchAll")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	if filter.OrderBy == "" {
		filter.OrderBy = "updated_at"
	}

	db := r.setFilterFine(globalshared.SetSpanToGorm(ctx, r.readDB), filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	})
	if filter.Limit > 0 || !filter.ShowAll {
		db = db.Limit(filter.Limit).Offset(filter.CalculateOffset())
	}
	err = db.Find(&data).Error
	return
}

func (r *fineRepoSQL) Count(ctx context.Context, filter *domain.FilterFine) (count int) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineRepoSQL:Count")
	defer trace.Finish()

	var total int64
	r.setFilterFine(globalshared.SetSpanToGorm(ctx, r.readDB), filter).Model(&shareddomain.Fine{}).Count(&total)
	count = int(total)

	trace.Log("count", count)
	return
}

func (r *fineRepoSQL) Find(ctx context.Context, filter *domain.FilterFine) (result shareddomain.Fine, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineRepoSQL:Find")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	err = r.setFilterFine(globalshared.SetSpanToGorm(ctx, r.readDB), filter).First(&result).Error
	return
}

func (r *fineRepoSQL) Save(ctx context.Context, data *shareddomain.Fine, updateOptions ...candishared.DBUpdateOptionFunc) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineRepoSQL:Save")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	db := r.writeDB
	if tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB); ok {
		db = tx
	}
	data.UpdatedAt = time.Now()
	if data.CreatedAt.IsZero() {
		data.CreatedAt = time.Now()
	}
	if data.ID == 0 {
		err = globalshared.SetSpanToGorm(ctx, db).Omit(clause.Associations).Create(data).Error
	} else {
		err = globalshared.SetSpanToGorm(ctx, db).Model(data).Omit(clause.Associations).Updates(r.updateTools.ToMap(data, updateOptions...)).Error
	}
	return
}

func (r *fineRepoSQL) Delete(ctx context.Context, filter *domain.FilterFine) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineRepoSQL:Delete")
	defer func() { trace.Finish(tracer.FinishWithError(err)) }()

	db := r.writeDB
	if tx, ok := candishared.GetValueFromContext(ctx, candishared.ContextKeySQLTransaction).(*gorm.DB); ok {
		db = tx
	}
	err = r.setFilterFine(globalshared.SetSpanToGorm(ctx, db), filter).Delete(&shareddomain.Fine{}).Error
	return
}

func (r *fineRepoSQL) setFilterFine(db *gorm.DB, filter *domain.FilterFine) *gorm.DB {

	if filter.ID != nil {
		db = db.Where("id = ?", *filter.ID)
	}
	if filter.Search != "" {
		db = db.Where("(field ILIKE '%%' || ? || '%%')", filter.Search)
	}

	for _, preload := range filter.Preloads {
		db = db.Preload(preload)
	}

	return db
}
