package tests

import (
	"net/http/httptest"

	_ "github.com/gtforge/BookStore/routers"

	"github.com/gtforge/services_common_go/gett-tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// A sample endpoint test
// Use `Ω` (⌥+Z) instead of `Expect` if you like Greek alphabet.
var _ = Describe("/alive", func() {
	var response *httptest.ResponseRecorder

	BeforeEach(func() {
		response = gettTests.MakeRequest("GET", "/alive")
	})

	It("should respond with 200 OK", func() {
		Expect(response.Code).To(Equal(200))
	})

	It("should respond with non-empty body", func() {
		Expect(response.Body.Len()).NotTo(BeZero())
	})
})
