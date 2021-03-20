package stores

import (
	"database/sql"
	"goApi/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	product = models.Product{Id: 1, Name: "French knife", Price: "Rs500"}
)

func newMockAndHandler(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *productStore) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	store := New(db)
	return db, mock, store
}

func TestStoreCreate(t *testing.T) {
	_, mock, store := newMockAndHandler(t)
	testcases := []struct {
		id          int
		input       models.Product
		call        *sqlmock.ExpectedExec
		expectedErr error
	}{
		{
			id:          1,
			input:       product,
			call:        mock.ExpectExec(`INSERT INTO PRODUCT \(id,name,price\) VALUES \(\?,\?,\?\)`).WithArgs(1,"French knife","Rs500").WillReturnResult(sqlmock.NewResult(0, 1)),
			expectedErr: nil,
		},
	}

	for _, eachCase := range testcases {
		err := store.Create(eachCase.input)

		if !reflect.DeepEqual(err, eachCase.expectedErr) {
			t.Errorf("Expected: %v\n Got: %v\n", eachCase.expectedErr, err)
		}
	}
}

func TestStoreRead(t *testing.T) {
	_, mock, store := newMockAndHandler(t)
	testcases := []struct {
		id          int
		call        *sqlmock.ExpectedQuery
		expectedOut []models.Product
		expectedErr error
	}{
		{
			id:          1,
			call:        mock.ExpectQuery(`SELECT id,name,price FROM PRODUCT`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).AddRow(product.Id, product.Name, product.Price)),
			expectedOut: []models.Product{product},
			expectedErr: nil,
		},
	}

	for _, eachCase := range testcases {
		resp, err := store.Read()

		if !reflect.DeepEqual(resp, eachCase.expectedOut) {
			t.Errorf("Testcase: %v\nExpected: %v\n Got: %v\n", eachCase.id, eachCase.expectedOut, resp)
		}
		if !reflect.DeepEqual(err, eachCase.expectedErr) {
			t.Errorf("Testcase: %v\nExpected: %v\n Got: %v\n", eachCase.id, eachCase.expectedErr, err)
		}
	}
}

func TestStoreUpdate(t *testing.T) {
	_, mock, store := newMockAndHandler(t)
	testcases := []struct {
		id          int
		inputId     int
		inputPrice  string
		call        *sqlmock.ExpectedExec
		expectedErr error
	}{
		{
			id:          1,
			inputId:     1,
			inputPrice:  "Rs499",
			call:        mock.ExpectExec(`UPDATE PRODUCT SET price=\? WHERE id=\?`).WithArgs("Rs499", 1).WillReturnResult(sqlmock.NewResult(0, 1)),
			expectedErr: nil,
		},
	}

	for _, eachCase := range testcases {
		err := store.Update(eachCase.inputPrice, eachCase.inputId)

		if !reflect.DeepEqual(err, eachCase.expectedErr) {
			t.Errorf("Testcase: %v\nExpected: %v\n Got: %v\n", eachCase.id, eachCase.expectedErr, err)
		}
	}
}
func TestStoreDelete(t *testing.T) {
	_, mock, store := newMockAndHandler(t)
	testcases := []struct {
		id          int
		input       int
		call        *sqlmock.ExpectedExec
		expectedErr error
	}{
		{
			id:          1,
			input:       1,
			call:        mock.ExpectExec(`DELETE FROM PRODUCT WHERE id=\?`).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1)),
			expectedErr: nil,
		},
	}

	for _, eachCase := range testcases {
		err := store.Delete(eachCase.input)
		if !reflect.DeepEqual(err, eachCase.expectedErr) {
			t.Errorf("Expected: %v\n Got: %v\n", eachCase.expectedErr, err)
		}
	}
}
