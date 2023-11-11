package controller

import (
	"errors"
	"product/datastore"
	"product/internal/constants"
	"product/internal/models"
	"product/internal/models/dtos"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var productVersionData = map[uint]models.ProductVersion{
	1: {Id: 1, ProductId: 1, IsActive: false},
	2: {Id: 2, ProductId: 1, IsActive: true},
}

type addTest struct {
	updateProductDetailsRequestDTOs dtos.UpdateProductDetailsRequestDTOs
	expectedErr                     error
}

var addTests = []addTest{{
	dtos.UpdateProductDetailsRequestDTOs{
		Id:        &[]uint{uint(1)}[0],
		VersionId: &[]uint{uint(1)}[0],
		Details:   &datatypes.JSON{}}, constants.ErrProductDetailNotExists},
	{dtos.UpdateProductDetailsRequestDTOs{
		Id:        &[]uint{uint(1)}[0],
		VersionId: &[]uint{uint(2)}[0],
		Details:   &datatypes.JSON{}}, nil}}

// Successfully update product details with valid input
func TestUpdateProductDetails(t *testing.T) {
	for _, test := range addTests {
		db1, mock, err := sqlmock.New()
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn:                      db1,
			SkipInitializeWithVersion: true,
		}), &gorm.Config{})

		datastore.MySQLConn = gormDB
		defer db1.Close()
		env := &dtos.Env{
			MySQLConn: gormDB,
		}

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `product_version` WHERE `id` = ? AND `is_active` = ? FOR UPDATE")).WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "details", "is_active"}).
			AddRow(productVersionData[*test.updateProductDetailsRequestDTOs.Id].Id,
				productVersionData[*test.updateProductDetailsRequestDTOs.Id].ProductId,
				productVersionData[*test.updateProductDetailsRequestDTOs.Id].Details,
				productVersionData[*test.updateProductDetailsRequestDTOs.Id].IsActive))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `product_version` SET `is_active`=? WHERE `id` = ?")).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `product_version` (`product_id`,`details`,`is_active`,`created`,`updated`) VALUES (?,NULL,?,?,?)")).
			WillReturnResult(sqlmock.NewResult(int64(*test.updateProductDetailsRequestDTOs.VersionId+1), 1))
		mock.ExpectCommit()

		err = UpdateProductDetails(env, test.updateProductDetailsRequestDTOs)
		if err == nil && test.expectedErr == nil {
			continue
		}
		if !errors.Is(err, test.expectedErr) {
			t.Errorf("Output %v not equal to expected %v", err, test.expectedErr)
		}
	}
}
