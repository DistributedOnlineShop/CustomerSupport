package util

import (
	"math/big"
	"math/rand/v2"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateRandomSupportCaseType() string {
	caseTypes := []string{
		"Technical Issue",
		"Billing Inquiry",
		"Account Support",
		"Feature Request",
		"Bug Report",
		"Product Inquiry",
		"General Inquiry",
	}
	return caseTypes[rand.IntN(len(caseTypes))]
}

func GenerateRandomSubject() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.BookTitle(),
		Valid:  true,
	}
}
func GenerateRandomMessage() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.BookGenre(),
		Valid:  true,
	}
}

func GenerateRandomDate() pgtype.Timestamp {
	daysOffset := rand.IntN(365) - 180
	return pgtype.Timestamp{
		Time:  time.Now().Add(time.Duration(daysOffset) * 24 * time.Hour),
		Valid: true,
	}
}

func GenerateRandomNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}
