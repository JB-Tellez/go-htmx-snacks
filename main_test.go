package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/calvinmclean/babyapi"
	babytest "github.com/calvinmclean/babyapi/test"
)

func TestAPI(t *testing.T) {
	defer os.RemoveAll("storage.json")

	os.Setenv("STORAGE_FILE", "storage.json")
	api := createAPI()

	babytest.RunTableTest(t, api, []babytest.TestCase[*babyapi.AnyResource]{
		{
			Name: "CreateSNACK",
			Test: babytest.RequestTest[*babyapi.AnyResource]{
				Method: http.MethodPost,
				Body:   `{"Title": "New SNACK"}`,
			},
			ExpectedResponse: babytest.ExpectedResponse{
				Status:     http.StatusCreated,
				BodyRegexp: `{"id":"[0-9a-v]{20}","Title":"New SNACK","Description":"","Eaten":null}`,
			},
		},
		{
			Name: "GetSNACK",
			Test: babytest.RequestTest[*babyapi.AnyResource]{
				Method: http.MethodGet,
				IDFunc: func(getResponse babytest.PreviousResponseGetter) string {
					return getResponse("CreateSNACK").Data.GetID()
				},
			},
			ExpectedResponse: babytest.ExpectedResponse{
				Status:     http.StatusOK,
				BodyRegexp: `{"id":"[0-9a-v]{20}","Title":"New SNACK","Description":"","Eaten":null}`,
			},
		},
		{
			Name: "DeleteSNACK",
			Test: babytest.RequestTest[*babyapi.AnyResource]{
				Method: http.MethodDelete,
				IDFunc: func(getResponse babytest.PreviousResponseGetter) string {
					return getResponse("CreateSNACK").Data.GetID()
				},
			},
			ExpectedResponse: babytest.ExpectedResponse{
				Status: http.StatusOK,
				NoBody: true,
			},
		},
	})
}
