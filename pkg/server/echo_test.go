package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho(t *testing.T) {
	// Begin by preparing the request.
	var jsonStr = []byte(`{"name":"user"}`)
	req, _ := http.NewRequest("POST", "/api/echo", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// Make the request and get a response out.
	w := httptest.NewRecorder()
	echoHandler(w, req)
	response := w.Result()

	// Check the status code was good.
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.StatusCode)
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Errorf("error decoding response responseBody: %v", err)
	}

	var expected map[string]interface{}
	if err := json.Unmarshal(jsonStr, &expected); err != nil {
		t.Errorf("error converting request payload into map: %v", err)
	}

	// Check that all the keys and values from our request payload are present
	// in the response.
	for k, v := range expected {
		// We will test the "echoed" key separately after his.
		if k == "echoed" {
			continue
		}

		respVal, exists := responseBody[k]
		if !exists {
			t.Errorf("%v doesn't exist in response", k)
		}

		if v != respVal {
			t.Errorf("expected '%v' and got '%v'", v, respVal)
		}
	}

	// Finally, let's check that "echoed" is set to true.
	echoedVal, exists := responseBody["echoed"]
	if !exists {
		t.Errorf("'echoed' is not set")
	}

	echoed, ok := echoedVal.(bool)
	if !ok {
		t.Errorf("error casting 'echoed' into a boolean")
	}
	if !echoed {
		t.Errorf("'echoed' is not set")
	}
}

func TestEchoWithEchoedField(t *testing.T) {
	// Begin by preparing the request.
	var jsonStr = []byte(`{"name":"user", "val": 123, "echoed": false}`)
	req, _ := http.NewRequest("POST", "/api/echo", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// Make the request and get a response out.
	w := httptest.NewRecorder()
	echoHandler(w, req)
	response := w.Result()

	// Check the status code was good.
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.StatusCode)
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Errorf("error decoding response responseBody: %v", err)
	}

	var expected map[string]interface{}
	if err := json.Unmarshal(jsonStr, &expected); err != nil {
		t.Errorf("error converting request payload into map: %v", err)
	}

	// Check that all the keys and values from our request payload are present
	// in the response.
	for k, v := range expected {
		// We will test the "echoed" key separately after his.
		if k == "echoed" {
			continue
		}

		respVal, exists := responseBody[k]
		if !exists {
			t.Errorf("%v doesn't exist in response", k)
		}

		if v != respVal {
			t.Errorf("expected '%v' and got '%v'", v, respVal)
		}
	}

	// Finally, let's check that "echoed" is set to true.
	echoedVal, exists := responseBody["echoed"]
	if !exists {
		t.Errorf("'echoed' is not set")
	}

	echoed, ok := echoedVal.(bool)
	if !ok {
		t.Errorf("error casting 'echoed' into a boolean")
	}
	if !echoed {
		t.Errorf("'echoed' is not set")
	}
}

func TestEchoWithEchoedFieldTrue(t *testing.T) {
	// Begin by preparing the request.
	var jsonStr = []byte(`{"name":"user", "val": 123, "echoed": true}`)
	req, _ := http.NewRequest("POST", "/api/echo", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// Make the request and get a response out.
	w := httptest.NewRecorder()
	echoHandler(w, req)
	response := w.Result()

	// Check the status code was good.
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusBadRequest, response.StatusCode)
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Errorf("error decoding response responseBody: %v", err)
	}

	// Finally, let's check that the error message is what we expected.
	errorVal, exists := responseBody["error"]
	if !exists {
		t.Error("expected 'error' field to be present when a bad request is sent")
	}

	errVal, ok := errorVal.(string)
	if !ok {
		t.Error("error casting error message into string")
	}

	if errVal != "request already had 'echoed: true'" {
		t.Errorf("unexpected error message: %v", errVal)
	}
}
