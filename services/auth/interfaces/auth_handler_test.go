package interfaces

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func TestSignUpOk(t *testing.T) {
	url := "http://localhost:3000/api/users/signup"
	statusCode := 201
	credentials := map[string]string{"email": "test_email@email.com", "password": "password"}

	payload, _ := json.Marshal(credentials)
	reader := bytes.NewReader(payload)

	res, err := http.Post(url, "application/json", reader)
	if err != nil {
		t.Fatalf("Could not make the get call, %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode == statusCode {
		t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)
	} else {
		t.Errorf("\t%s\tShould receive a %d status code : %v", failed, statusCode, res.StatusCode)
	}

}

func TestSignUpBad(t *testing.T) {
	url := "http://localhost:3000/api/users/signup"
	statusCode := 400
	credentials := map[string]string{"email": "test_email@email.com", "password": "password"}

	payload, _ := json.Marshal(credentials)
	reader := bytes.NewReader(payload)

	res, err := http.Post(url, "application/json", reader)
	if err != nil {
		t.Fatalf("Could not make the get call, %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode == statusCode {
		t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)
	} else {
		t.Errorf("\t%s\tShould receive a %d status code : %v", failed, statusCode, res.StatusCode)
	}

}
