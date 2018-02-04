package internal

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler(e Endpoints) http.Handler {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{}
	r.Methods("GET").Path("/tasks").Handler(
		httptransport.NewServer(
			e.GetTasksEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
	)
	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}
