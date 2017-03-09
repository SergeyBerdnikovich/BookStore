package tests

import (
	"net/http/httptest"

	_ "github.com/gtforge/{|PROJECTNAME|}/routers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// A sample endpoint test
// Use `Ω` (⌥+Z) instead of `Expect` if you like Greek alphabet.

var _ = Describe("/alive", func() {

	var response *httptest.ResponseRecorder

	BeforeEach(func() {
		response = MakeTestRequest("GET", "/alive", nil)
	})

	It("should respond with 200 OK", func() {
		Expect(response.Code).To(Equal(200))
	})

	It("should respond with non-empty body", func() {
		Expect(response.Body.Len()).NotTo(BeZero())
	})
})
