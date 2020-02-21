package pkg_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	."bankInfo/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPkg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pkg Suite")
}

const filepath  = "../bank.csv"
func performRequest(r http.Handler, method string, path string, body []byte) *httptest.ResponseRecorder {
	BankDetails=CsvReader(filepath)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
