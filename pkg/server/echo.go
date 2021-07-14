package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// ErrBadRequest is the error message we reply with when unable to parse the
	// request's body.
	ErrBadRequest = map[string]string{"error": "could not parse request payload"}
	// ErrTrueEchoPResent is the error message we reply with when "echoed" is
	// present and equalto true.
	ErrTrueEchoPResent = map[string]string{"error": "request already had 'echoed: true'"}
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	// Keep track of our metrics.
	requestsTotal.With(prometheus.Labels{"code": strconv.Itoa(code)}).Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response) // Lint: We assign 2 empty values to resolve an errcheck error.
}

// echoHandler echoes back the JSON the client sends us except in the case in
// which the request's body already has the top-level field "echoed" set to
// true.
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// We will parse the request's body into a map so we can just add the value
	// "echoed".
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, ErrBadRequest)
		return
	}
	defer r.Body.Close()

	// However, before we add "echoed: true", let's check that it isn't already
	// set as such. If it is, then let's return a 400.
	echoedVal, exists := body["echoed"]
	if exists {
		echoed, ok := echoedVal.(bool)
		if ok {
			if echoed {
				respondWithJSON(w, http.StatusBadRequest, ErrTrueEchoPResent)
				return
			}
		}
	}

	// Finally, set the value for "echoed".
	body["echoed"] = true
	respondWithJSON(w, http.StatusOK, body)
}
