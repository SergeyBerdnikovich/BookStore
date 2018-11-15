package books_store

import (
	"errors"
	"testing"

	"github.com/gtforge/BookStore/models"
	"github.com/gtforge/BookStore/store/books-store/pg-store/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooksStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BooksStore suite")
}

var _ = Describe("Testing BooksStore", func() {
	var (
		subject  *booksStore
		mockCtrl *gomock.Controller
		mockDB   *mocks.MockPgBooksStoreInterface
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockDB = mocks.NewMockPgBooksStoreInterface(mockCtrl)

		subject = &booksStore{db: mockDB}
	})

	Describe(".GetAll", func() {
		var (
			books_stub []models.Book
			err_stub   error
		)



		Context("when we fetch all books from DB without error", func() {
			BeforeEach(func() {
				// You should do mock with right return value, without it here it will do mock with empty value
				books_stub = []models.Book{
					{ID: 1, Name: "Test 1", Quantity: 1},
					{ID: 2, Name: "Test 2", Quantity: 2},
				}

				mockDB.EXPECT().GetAll().Return(books_stub, err_stub).Times(1)
			})

			It("should return slice of books without error", func() {
				books, err := subject.GetAll()

				Expect(err).To(Succeed())
				// Expect(err).To(Equal(err_stub)) // <-- Refusing to compare <nil> to <nil>. We don't need this, prev check ensure it
				Expect(books).To(Equal(books_stub))
			})
		})

		Context("when there is some error during books fetching", func() {
			BeforeEach(func() {
				// define values BEFORE mock:
				books_stub = []models.Book{}
				err_stub = errors.New("Some error")

				mockDB.EXPECT().GetAll().Return(books_stub, err_stub).Times(1)
			})

			It("should return empty slice of books with error", func() {
				books, err := subject.GetAll()

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(err_stub))
				Expect(books).To(Equal(books_stub))
			})
		})
	})
})
