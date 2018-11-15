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
			booksStub []models.Book
			errorStub error
		)

		Context("when we fetch all books from DB without error", func() {
			BeforeEach(func() {
				booksStub = []models.Book{
					{ID: 1, Name: "Test 1", Quantity: 1},
					{ID: 2, Name: "Test 2", Quantity: 2},
				}
				errorStub = nil

				mockDB.EXPECT().GetAll().Return(booksStub, errorStub).Times(1)
			})

			It("should return slice of books without error", func() {
				books, err := subject.GetAll()

				Expect(err).To(Succeed())
				Expect(books).To(Equal(booksStub))
			})
		})

		Context("when there is some error during books fetching", func() {
			BeforeEach(func() {
				booksStub = []models.Book{}
				errorStub = errors.New("Some error")

				mockDB.EXPECT().GetAll().Return(booksStub, errorStub).Times(1)
			})

			It("should return empty slice of books with error", func() {
				books, err := subject.GetAll()

				Expect(err).To(HaveOccurred())
				Expect(books).To(Equal(booksStub))
			})
		})

	})

	Describe(".GetPerPage", func() {
		var (
			booksStub  []models.Book
			errorStub  error
			pageNumber int
		)

		Context("when we fetch all books from DB without error", func() {
			BeforeEach(func() {
				booksStub = []models.Book{
					{ID: 1, Name: "Test 1", Quantity: 1},
					{ID: 2, Name: "Test 2", Quantity: 2},
				}
				errorStub = nil
				pageNumber = 1

				mockDB.EXPECT().GetPerPage(pageNumber).Return(booksStub, errorStub).Times(1)
			})

			It("should return slice of books without error", func() {
				books, err := subject.GetPerPage(pageNumber)

				Expect(err).To(Succeed())
				Expect(books).To(Equal(booksStub))
			})
		})

		Context("when there is some error during books fetching", func() {
			BeforeEach(func() {
				booksStub = []models.Book{}
				errorStub = errors.New("Some error")

				mockDB.EXPECT().GetPerPage(pageNumber).Return(booksStub, errorStub).Times(1)
			})

			It("should return empty slice of books with error", func() {
				books, err := subject.GetPerPage(pageNumber)

				Expect(err).To(HaveOccurred())
				Expect(books).To(Equal(booksStub))
			})
		})
	})

	Describe(".GetById", func() {
		var (
			bookStub  models.Book
			errorStub error
			bookId    uint
		)

		Context("when we fetch a book from DB without error", func() {
			BeforeEach(func() {
				bookStub = models.Book{ID: 1, Name: "Test 1", Quantity: 1}
				errorStub = nil
				bookId = 1

				mockDB.EXPECT().GetById(bookId).Return(bookStub, errorStub).Times(1)
			})

			It("should return a book without error", func() {
				book, err := subject.GetById(bookId)

				Expect(err).To(Succeed())
				Expect(book).To(Equal(bookStub))
			})
		})

		Context("when there is some error during books fetching", func() {
			BeforeEach(func() {
				bookStub = models.Book{}
				errorStub = errors.New("Some error")

				mockDB.EXPECT().GetById(bookId).Return(bookStub, errorStub).Times(1)
			})

			It("should return empty Book with error", func() {
				book, err := subject.GetById(bookId)

				Expect(err).To(HaveOccurred())
				Expect(book).To(Equal(bookStub))
			})
		})
	})

	Describe(".Create", func() {
		var (
			errorStub error
			name      string
			quantity  int
		)

		Context("when we create a book without error", func() {
			BeforeEach(func() {
				errorStub = nil
				name = "Test 1"
				quantity = 1

				mockDB.EXPECT().Create(name, quantity).Return(errorStub).Times(1)
			})

			It("should not return an error", func() {
				err := subject.Create(name, quantity)

				Expect(err).To(Succeed())
			})
		})

		Context("when there is some error during book creation", func() {
			BeforeEach(func() {
				errorStub = errors.New("Some error")

				mockDB.EXPECT().Create(name, quantity).Return(errorStub).Times(1)
			})

			It("should return an error", func() {
				err := subject.Create(name, quantity)

				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe(".Update", func() {
		var (
			errorStub error
			bookId    uint
			name      string
			quantity  int
		)

		Context("when we update a book without error", func() {
			BeforeEach(func() {
				errorStub = nil
				bookId = 1
				name = "Test 1"
				quantity = 1

				mockDB.EXPECT().Update(bookId, name, quantity).Return(errorStub).Times(1)
			})

			It("should not return an error", func() {
				err := subject.Update(bookId, name, quantity)

				Expect(err).To(Succeed())
			})
		})

		Context("when there is some error during book updating", func() {
			BeforeEach(func() {
				errorStub = errors.New("Some error")

				mockDB.EXPECT().Update(bookId, name, quantity).Return(errorStub).Times(1)
			})

			It("should return an error", func() {
				err := subject.Update(bookId, name, quantity)

				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe(".Delete", func() {
		var (
			errorStub error
			bookId    uint
		)

		Context("when we delete a book without error", func() {
			BeforeEach(func() {
				errorStub = nil
				bookId = 1

				mockDB.EXPECT().Delete(bookId).Return(errorStub).Times(1)
			})

			It("should not return an error", func() {
				err := subject.Delete(bookId)

				Expect(err).To(Succeed())
			})
		})

		Context("when there is some error during book deletion", func() {
			BeforeEach(func() {
				errorStub = errors.New("Some error")

				mockDB.EXPECT().Delete(bookId).Return(errorStub).Times(1)
			})

			It("should return an error", func() {
				err := subject.Delete(bookId)

				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe(".DecreaseQuantity", func() {
		var (
			errorStub error
			bookId    uint
		)

		Context("when we decrease a quantity of a book without error", func() {
			BeforeEach(func() {
				errorStub = nil
				bookId = 1

				mockDB.EXPECT().DecreaseQuantity(bookId).Return(errorStub).Times(1)
			})

			It("should not return an error", func() {
				err := subject.DecreaseQuantity(bookId)

				Expect(err).To(Succeed())
			})
		})

		Context("when there is some error during book quantity decreasing", func() {
			BeforeEach(func() {
				errorStub = errors.New("Some error")

				mockDB.EXPECT().DecreaseQuantity(bookId).Return(errorStub).Times(1)
			})

			It("should return an error", func() {
				err := subject.DecreaseQuantity(bookId)

				Expect(err).To(HaveOccurred())
			})
		})
	})
})
