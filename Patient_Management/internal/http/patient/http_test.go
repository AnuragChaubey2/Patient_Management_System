package patient

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/anurag120799/patient-management/internal/models"
	"github.com/anurag120799/patient-management/internal/service"
	"github.com/golang/mock/gomock"
)

var patient = models.Patient{
	ID:          5,
	Name:        "ZopSmart",
	Phone:       "+919172681679",
	Discharged:  true,
	CreatedAt:   "2022-02-22 13:23:22",
	UpdatedAt:   "2022-02-22 13:39:41",
	BloodGroup:  "+A",
	Description: "patient description",
}

func Test_GetByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPatientService := service.NewMockPatient(mockCtrl)
	testCases := []struct {
		id            int
		mockCall      *gomock.Call
		expectedError error
	}{
		// Success
		{
			id:            5,
			mockCall:      mockPatientService.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&patient, nil),
			expectedError: nil,
		},
		// Failure
		{
			id:            -1,
			mockCall:      mockPatientService.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, errors.InvalidParam{Param: []string{"id"}}),
			expectedError: errors.InvalidParam{Param: []string{"id"}},
		},
	}
	p := New(mockPatientService)

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "http://localhost:9000/patients/{id}", nil)
		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)

		_, err := p.GetByID(ctx)

		if !reflect.DeepEqual(testCase.expectedError, err) {
			t.Errorf("Expected error: %v Got %v", testCase.expectedError, err)
		}
	}
}

func Test_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPatientService := service.NewMockPatient(mockCtrl)
	testCases := []struct {
		body          []byte
		input         models.Patient
		mockCall      *gomock.Call
		expectedError error
	}{
		// Success
		{
			body: []byte(`{
				"name": "Zopsmart",
				"phone": "+919172681679",
				"discharged": true,
				"createdAt": "2022-02-22 13:23:22",
				"updatedAt": "2022-02-22 13:39:41",
				"bloodGroup": "+A",
				"description": "patient description"
				}`),
			input:         patient,
			mockCall:      mockPatientService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&patient, nil),
			expectedError: nil,
		},
		// Failure
		{
			body: []byte(`{
				"name": "Zopsmart",
				"phone": "+919172681679",
				"discharged": true,
				"createdAt": "2022-02-22 13:23:22",
				"updatedAt": "2022-02-22 13:39:41",
				"bloodGroup": "+A",
				"description": "patient description"
				}`),
			input:         patient,
			mockCall:      mockPatientService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.Error("invalid fileds")),
			expectedError: errors.Error("invalid fileds"),
		},
		{
			body: []byte(`{
				"name": "Zopsmart",
				"phone": "+919172681679",
				"discharged": "sdjfhje",
				"createdAt": "2022-02-22 13:23:22",
				"updatedAt": "2022-02-22 13:39:41",
				"bloodGroup": "+A",
				"description": "patient description"
				}`),
			input: patient,
			// mockCall:      mockPatientService.EXPECT().CreatePatientService(gomock.Any()).Return(models.Patient{}, errors.New("error")),
			expectedError: errors.Error("cannot read from body"),
		},
	}

	p := New(mockPatientService)

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/patients", bytes.NewReader(testCase.body))
		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)

		_, err := p.Create(ctx)

		if !reflect.DeepEqual(testCase.expectedError, err) {
			t.Errorf("Expected error: %v Got %v", testCase.expectedError, err)
		}
	}
}

func Test_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockPatientService := service.NewMockPatient(mockCtrl)
	testCases := []struct {
		mockCall      *gomock.Call
		expectedError error
	}{
		// Success
		{
			mockCall:      mockPatientService.EXPECT().Get(gomock.Any()).Return([]*models.Patient{&patient}, nil),
			expectedError: nil,
		},
		// Failure
		{
			mockCall:      mockPatientService.EXPECT().Get(gomock.Any()).Return(nil, errors.EntityNotFound{Entity: "Patient"}),
			expectedError: errors.EntityNotFound{Entity: "Patient"},
		},
	}

	p := New(mockPatientService)

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/patients", nil)
		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)

		_, err := p.Get(ctx)

		if !reflect.DeepEqual(testCase.expectedError, err) {
			t.Errorf("Expected error: %v Got %v", testCase.expectedError, err)
		}
	}
}

func Test_Update(t *testing.T) {
	var pat = models.Patient{
		Name:        "ZopSmart",
		Description: "patient description",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPatientService := service.NewMockPatient(mockCtrl)
	testCases := []struct {
		body          []byte
		mockCall      *gomock.Call
		expectedError error
	}{
		// Success
		{
			body: []byte(`{
				"name": "ZopSmart",
				"description": "patient description"
				}`),
			mockCall:      mockPatientService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(&pat, nil),
			expectedError: nil,
		},
		// Failure
		{
			body: []byte(`{
				"name": "ZopSmart",
				"description": "patient description"
				}`),
			mockCall: mockPatientService.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil, errors.EntityNotFound{Entity: "Patient", ID: "id"}),
			expectedError: errors.EntityNotFound{Entity: "Patient", ID: "id"},
		},
		// Failure
		{
			body: []byte(`{
				"name": 1,
				"description": "patient description"
				}`),
			expectedError: errors.Error("cannot read from body"),
		},
	}

	p := New(mockPatientService)

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/patients/{id}", bytes.NewReader(testCase.body))
		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)

		_, err := p.Update(ctx)

		if !reflect.DeepEqual(testCase.expectedError, err) {
			t.Errorf("Expected error: %v Got %v", testCase.expectedError, err)
		}
	}
}

func Test_DeletePatientService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPatientService := service.NewMockPatient(mockCtrl)
	testCases := []struct {
		mockCall      *gomock.Call
		expectedError error
	}{
		// Success
		{
			mockCall:      mockPatientService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil),
			expectedError: nil,
		},
		// Failure
		{
			mockCall:      mockPatientService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.InvalidParam{Param: []string{"id"}}),
			expectedError: errors.InvalidParam{Param: []string{"id"}},
		},
	}

	p := New(mockPatientService)

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "http://localhost:9000/patients/{id}", nil)
		ctx := gofr.NewContext(responder.NewContextualResponder(httptest.NewRecorder(), req), request.NewHTTPRequest(req), nil)

		_, err := p.Delete(ctx)

		if !reflect.DeepEqual(testCase.expectedError, err) {
			t.Errorf("Expected error: %v Got %v", testCase.expectedError, err)
		}
	}
}
