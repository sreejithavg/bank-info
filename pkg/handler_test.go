package pkg_test

import (
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("non-cellular users", func() {
	router:=mux.NewRouter()


	It("return success msg", func() {
		w:=performRequest(router,"GET","/user/non_cellular",nil)
		Expect(w.Code).To(BeEquivalentTo(http.StatusOK))
	})
})

var _= Describe("user details", func() {
	router:=mux.NewRouter()
	It("return success msg", func() {
		w:=performRequest(router,"GET","/user/get",nil)
		Expect(w.Code).To(BeEquivalentTo(http.StatusOK))
	})
})
var _= Describe("filter user details", func() {
	router:=mux.NewRouter()

	It("return success msg", func() {
		body:=[]byte(`{
				"start_age":20,
				"end_age":30,
				"martial_status":"married"
		}`)
		w:=performRequest(router,"POST","/user/filter",body)
		Expect(w.Code).To(BeEquivalentTo(http.StatusOK))
	})
})
