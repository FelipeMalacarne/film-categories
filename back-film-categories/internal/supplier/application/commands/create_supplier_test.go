package commands

import (
	"testing"

	"github.com/felipemalacarne/back-film-categories/internal/supplier/domain"
	"github.com/felipemalacarne/back-film-categories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandleSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)
	handler := NewCreateSupplierHandler(mockRepo)

	command := CreateSupplierCommand{
		Name:  "John Doe",
		Email: "john@example.com",
		Phone: "+55549917158012",
	}

	expectedSupplier := domain.Supplier{
		Name:  command.Name,
		Email: command.Email,
		Phone: command.Phone,
	}

	mockRepo.EXPECT().Create(gomock.Any()).Return(expectedSupplier, nil)

	supplier, err := handler.Handle(command)

	assert.NoError(t, err)
	assert.Equal(t, expectedSupplier, supplier)
}

func TestHandleFailureOnDomainCreation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockSupplierRepository(mockCtrl)
	handler := NewCreateSupplierHandler(mockRepo)

	command := CreateSupplierCommand{
		Name:  "", // invalid input to trigger domain creation failure
		Email: "john@example.com",
		Phone: "1234567890",
	}

	supplier, err := handler.Handle(command)

	assert.Error(t, err)
	assert.Equal(t, domain.Supplier{}, supplier)
}
