package main

import (
	"github.com/unrolled/render"
	"net/http"
)

// API Ping Handler
func getpayments(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Payment alive!"})
	}
}