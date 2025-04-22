package util

import (
	"math/big"
	"math/rand/v2"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateSupportCaseType() string {
	caseTypes := []string{
		"Technical Issue",
		"Billing Inquiry",
		"Account Support",
		"Feature Request",
		"Bug Report",
		"Product Inquiry",
		"Generate Inquiry",
	}
	return caseTypes[rand.IntN(len(caseTypes))]
}

func GenerateSubject() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.BookTitle(),
		Valid:  true,
	}
}
func GenerateMessage() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.BookGenre(),
		Valid:  true,
	}
}

func GenerateDate() pgtype.Timestamp {
	daysOffset := rand.IntN(365) - 180
	return pgtype.Timestamp{
		Time:  time.Now().Add(time.Duration(daysOffset) * 24 * time.Hour),
		Valid: true,
	}
}

func GenerateNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}
