package pkg

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

func performRequest(r http.Handler, method string, path string, body []byte) *httptest.ResponseRecorder {
	CsvReader(filePath)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var _ = Describe("non-cellular users", func() {
	router := InitRouter()
	It("return success msg", func() {
		w := performRequest(router, "GET", "/user/non_cellular", nil)
		Expect(w.Code).To(BeEquivalentTo(http.StatusOK))
	})

	It("return 404", func() {
		w := performRequest(router, "GET", "/user/non_cellu", nil)
		Expect(w.Code).To(BeEquivalentTo(http.StatusNotFound))
	})
})

var _ = Describe("user details", func() {
	router := InitRouter()
	It("return success msg", func() {
		w := performRequest(router, "GET", "/user/get", nil)
		Expect(w.Code).To(BeEquivalentTo(http.StatusOK))
	})
	It("return 404", func() {
		w := performRequest(router, "GET", "/user/g", nil)
		Expect(w.Code).To(BeEquivalentTo(http.StatusNotFound))
	})
})
var _ = Describe("filter user details", func() {
	router := InitRouter()

	It("return success msg", func() {
		body := []byte(`{
				"start_age":20,
				"end_age":30,
				"martial_status":"married"
		}`)
		w := performRequest(router, "POST", "/user/filter", body)
		Expect(w.Code).To(BeEquivalentTo(http.StatusOK))
	})
	It("return 400 validation failed", func() {
		body := []byte(`{
				"start_age":20,
				"end_age":40,
		}`)
		w := performRequest(router, "POST", "/user/filter", body)
		Expect(w.Code).To(BeEquivalentTo(http.StatusBadRequest))
	})
	It("return 400 validation failed for end_age", func() {
		body := []byte(`{
				"start_age":20,
				"end_age":10,
				"martial_status":"married"
		}`)
		w := performRequest(router, "POST", "/user/filter", body)
		Expect(w.Code).To(BeEquivalentTo(http.StatusBadRequest))
	})
	It("return 400 binding failed", func() {
		body := []byte(`{
				"start_age":20,
				"end_age":"10",
		}`)
		w := performRequest(router, "POST", "/user/filter", body)
		Expect(w.Code).To(BeEquivalentTo(http.StatusBadRequest))
	})
	It("return 404", func() {
		w := performRequest(router, "GET", "//user/g", nil)
		Expect(w.Code).To(BeEquivalentTo(http.StatusNotFound))
	})
})
