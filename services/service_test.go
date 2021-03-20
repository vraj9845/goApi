package services

import (
	"goApi/models"
	"goApi/stores"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	product = models.Product{Id: 1, Name: "French knife", Price: "Rs500"}
)

func initializeTest(t *testing.T) (*stores.MockProduct, productService) {
	mockCtrl := gomock.NewController(t)
	mockStore := stores.NewMockProduct(mockCtrl)
	service := New(mockStore)
	return mockStore, service
}

func TestServiceCreate(t *testing.T) {
	mockStore, service := initializeTest(t)
	testcases := []struct {
		id          int
		input       models.Product
		call        *gomock.Call
		expectedErr error
	}{
		{
			id:          1,
			input:       product,
			call:        mockStore.EXPECT().Create(gomock.Any()).Return(nil),
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		err := service.Create(tc.input)

		if !reflect.DeepEqual(err, tc.expectedErr) {
			t.Errorf("Testcase: %v\nExpected: %v\nGot: %v\n", tc.id, tc.expectedErr, err)
		}
	}
}

func TestServiceRead(t *testing.T) {
	mockStore, service := initializeTest(t)
	testcases := []struct {
		id          int
		call        *gomock.Call
		expectedOut []models.Product
		expectedErr error
	}{
		{
			id:          1,
			call:        mockStore.EXPECT().Read().Return([]models.Product{product}, nil),
			expectedOut: []models.Product{product},
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		resp, err := service.Read()

		if !reflect.DeepEqual(resp, tc.expectedOut) {
			t.Errorf("Testcase: %v\nExpected: %v\nGot: %v\n", tc.id, tc.expectedOut, resp)
		}

		if !reflect.DeepEqual(err, tc.expectedErr) {
			t.Errorf("Testcase: %v\nExpected: %v\nGot: %v\n", tc.id, tc.expectedErr, err)
		}
	}
}

func TestServiceUpdate(t *testing.T) {
	mockStore, service := initializeTest(t)
	testcases := []struct {
		id          int
		inputPrice  string
		inputId     int
		call        *gomock.Call
		expectedErr error
	}{
		{
			id:          1,
			inputPrice:  "Rs500",
			inputId:     1,
			call:        mockStore.EXPECT().Update("Rs500", 1).Return(nil),
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		err := service.Update(tc.inputPrice, tc.inputId)

		if !reflect.DeepEqual(err, tc.expectedErr) {
			t.Errorf("Testcase: %v\nExpected: %v\nGot: %v\n", tc.id, tc.expectedErr, err)
		}
	}
}

func TestServiceDelete(t *testing.T) {
	mockStore, service := initializeTest(t)
	testcases := []struct {
		id          int
		inputId     int
		call        *gomock.Call
		expectedErr error
	}{
		{
			id:          1,
			inputId:     1,
			call:        mockStore.EXPECT().Delete(1).Return(nil),
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		err := service.Delete(tc.inputId)

		if !reflect.DeepEqual(err, tc.expectedErr) {
			t.Errorf("Testcase: %v\nExpected: %v\nGot: %v\n", tc.id, tc.expectedErr, err)
		}
	}
}
