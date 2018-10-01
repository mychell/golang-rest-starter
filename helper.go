package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Logger return log message
// func Logger() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(time.Now(), r.Method, r.URL)
// 		router.ServeHTTP(w, r) // dispatch the request
// 	})
// }
