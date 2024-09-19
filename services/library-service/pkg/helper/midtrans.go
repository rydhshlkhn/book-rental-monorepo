package helper

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"monorepo/services/library-service/pkg/shared"
	"strconv"
	"strings"
)

func GetIDFromOrderID(str string, prefix string) (int, error) {
	if !strings.HasPrefix(str, prefix) {
		return 0, fmt.Errorf("string does not start with %s", prefix)
	}

	idStr := strings.TrimPrefix(str, prefix)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("error converting transactionID to integer: %w", err)
	}

	return id, nil
}

func VerifyingSignatureKey(signatureKey, orderID, statusCode, grossAmount string) (isValid bool) {
	input := fmt.Sprintf("%s%s%s%s", orderID, statusCode, grossAmount, shared.GetEnv().MidtransServerKey)
	hash := sha512.New()
	hash.Write([]byte(input))

	// Mengambil hasil hash dalam bentuk byte
	hashedBytes := hash.Sum(nil)

	// Mengonversi hasil hash ke format hex string
	hashedString := hex.EncodeToString(hashedBytes)
	fmt.Println(hashedString, orderID, statusCode, grossAmount, shared.GetEnv().MidtransServerKey)

	return hashedString == signatureKey
}
