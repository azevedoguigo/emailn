package endpoint

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_Endpoint_Returns_Internal_Error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerros.ErrInternal
	}
	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerros.ErrInternal.Error())
}

func Test_HandlerError_Endpoint_Returns_Domain_Error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("domain error")
	}
	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func Test_HandlerError_Endpoint_Returns_Obj_And__Status(t *testing.T) {
	assert := assert.New(t)

	type ResponseBody struct {
		ID int
	}
	expectedResponseBody := ResponseBody{
		ID: 2,
	}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return expectedResponseBody, http.StatusOK, nil
	}
	handlerFunc := HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	responseBody := ResponseBody{}
	json.Unmarshal(res.Body.Bytes(), &responseBody)
	assert.Equal(expectedResponseBody, responseBody)
}
