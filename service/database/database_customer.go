package database

import (
	"context"

	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	dberrors "mylearning.com/golang/micro/ln/pharmacy/service/dberrors"
	models "mylearning.com/golang/micro/ln/pharmacy/service/models"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	if emailAddress == "" {
		result := c.DB.WithContext(ctx).
			Find(&customers)
		return customers, result.Error

	} else {
		result := c.DB.WithContext(ctx).
			Where(models.Customer{Email: emailAddress}).
			Find(&customers)
		return customers, result.Error

	}

}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error

	}
	return customer, nil
}
