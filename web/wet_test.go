package web

import (
	"net/http"
	"testing"
)

func TestConfigParse(t *testing.T) {
	http.HandleFunc("/fuck", ConfigParse)
	http.HandleFunc("/conf", ConfigStr)
	_ = http.ListenAndServe(":9999",nil)
}