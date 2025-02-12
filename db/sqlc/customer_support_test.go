package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"

	"CustomerSupport/util"
)

func CreateRandomCustomerSupportCase(t *testing.T) CustomerSupport {

	data := CreateCustomerSupportCaseParams{
		CsID:    util.CreateUUID(),
		UserID:  util.CreateUUID(),
		OrderID: util.CreateUUID(),
		Subject: util.GenerateRandomSubject(),
		Message: util.GenerateRandomMessage(),
		Status:  util.GenerateRandomSupportCaseType(),
	}

	SupportCase, err := testStore.CreateCustomerSupportCase(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, SupportCase)
	require.Equal(t, SupportCase.CsID, data.CsID)
	require.Equal(t, SupportCase.UserID, data.UserID)
	require.Equal(t, SupportCase.OrderID, data.OrderID)
	require.Equal(t, SupportCase.Subject.String, data.Subject.String)
	require.Equal(t, SupportCase.Message.String, data.Message.String)
	require.Equal(t, SupportCase.Status, data.Status)
	require.NotZero(t, SupportCase.CreatedAt)

	return SupportCase
}

func TestCreateCustomerSupportCase(t *testing.T) {
	CreateRandomCustomerSupportCase(t)
}

func TestGetCustomerSupportCaseById(t *testing.T) {
	data := CreateRandomCustomerSupportCase(t)

	supportCase, err := testStore.GetCustomerSupportCaseById(context.Background(), data.CsID)
	require.NoError(t, err)
	require.NotEmpty(t, supportCase)
}
func TestGetCustomerSupportCaseList(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomCustomerSupportCase(t)
	}

	supportCase, err := testStore.GetCustomerSupportCaseList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, supportCase)
	require.GreaterOrEqual(t, len(supportCase), 10)
}

func TestUpdateCustomerSupportCaseStatus(t *testing.T) {
	data := CreateRandomCustomerSupportCase(t)

	newStatus := util.GenerateRandomSupportCaseType()
	if newStatus == data.Status {
		for {
			newStatus = util.GenerateRandomSupportCaseType()
			if newStatus != data.Status {
				break
			}
		}
	}
	newMessage := util.GenerateRandomSupportCaseType()
	if newMessage == data.Message.String {
		for {
			newMessage = util.GenerateRandomSupportCaseType()
			if newMessage != data.Message.String {
				break
			}
		}
	}

	newData := UpdateCustomerSupportCaseStatusParams{
		CsID:    data.CsID,
		Status:  newStatus,
		Message: pgtype.Text{String: newMessage, Valid: true},
	}

	updatedSupportCase, err := testStore.UpdateCustomerSupportCaseStatus(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updatedSupportCase)
	require.Equal(t, updatedSupportCase.CsID, data.CsID)
	require.Equal(t, updatedSupportCase.UserID, data.UserID)
	require.Equal(t, updatedSupportCase.OrderID, data.OrderID)
	require.Equal(t, updatedSupportCase.Subject.String, data.Subject.String)
	require.Equal(t, updatedSupportCase.Message.String, newMessage)
	require.Equal(t, updatedSupportCase.Status, newStatus)
	require.NotZero(t, updatedSupportCase.CreatedAt)
	require.NotZero(t, updatedSupportCase.UpdatedAt)
}
