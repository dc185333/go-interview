package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DB interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

type FooRequestBody struct{ Entries []struct{ Key, Value string } }
type FooHandler struct{ DB DB }

func (f *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if keys, ok := r.URL.Query()["key"]; ok {
			for _, key := range keys {
				if val, ok := f.DB.Get(key); ok {
					fmt.Fprintf(w, "%s\n", val)
					continue
				}
				fmt.Fprintf(w, "Key %s not found\n", key)
			}
		}
		http.Error(w, "Invalid request: no keys in query", http.StatusBadRequest)
	case http.MethodPost:
		var body FooRequestBody

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, e := range body.Entries {
			f.DB.Set(e.Key, e.Value)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	myDB := &myDB{}
	myDB.Load()
	http.Handle("/foo", &FooHandler{DB: myDB})
	http.ListenAndServe(":8080", nil)
}
