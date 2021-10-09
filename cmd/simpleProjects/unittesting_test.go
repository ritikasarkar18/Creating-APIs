package main

// tried to write unit tests, but found out it works differently for mongo-driver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// credits : https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d

func TestcreateUser(t *testing.T) {

	var jsonStr = []byte(`{"Name": "Mashi",
    "Email": "tanya@gmail.com",
    "Password": "bhved3"}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	/* cannot write the following code, as our returned response will contain mongo generated id, as well as hashed password*/
	// expected := `{{"Name": "Mashi", "Email": "tanya@gmail.com", "Password": "bhved3"}}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}
