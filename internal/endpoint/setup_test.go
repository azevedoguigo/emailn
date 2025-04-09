package endpoint

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/azevedoguigo/emailn/internal/test/internalmock"
	"github.com/go-chi/chi/v5"
)

var (
	service *internalmock.CampaignServiceMock
	handler = Handler{}
)

func Setup() {
	service = new(internalmock.CampaignServiceMock)
	handler.CampaignService = service
}

func NewRequestAndRecorder(method, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
	var buffer bytes.Buffer
	if body != nil {
		json.NewEncoder(&buffer).Encode(body)
	}

	request, _ := http.NewRequest(method, url, &buffer)
	recorder := httptest.NewRecorder()

	return request, recorder
}

func AddParameter(req *http.Request, key, value string) *http.Request {
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add(key, value)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

	return req
}
