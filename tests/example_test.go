package tests

import (
	"net/http/httptest"

	_ "github.com/gtforge/{|PROJECTNAME|}/routers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	initBeego()
}

// A sample endpoint test
// Use `Expect` instead of `Ω` (⌥+Z) if you don't like Greek alphabet.

var _ = Describe("/alive", func() {

	var response *httptest.ResponseRecorder

	BeforeEach(func() {
		response = MakeTestRequest("GET", "/alive", nil)
	})

	It("should respond with 200 OK", func() {
		Ω(response.Code).To(Equal(200))
	})

	It("should respond with non-empty body", func() {
		Ω(response.Body.Len()).NotTo(BeZero())
	})
})
