package main

import (
	"log"
	"net/http"

	"github.com/gomicroservice/micro/backend/rest/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &graphql.Schema,
		Pretty: true,
	})
	http.Handle("/graphql", disableCors(h))
	log.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
