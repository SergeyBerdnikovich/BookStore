package pg_store

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gtforge/BookStore/models"

	monkeyExtensions "github.com/artiomgiza/monkey-flexible-extensions"
	"github.com/bouk/monkey"
	"github.com/jinzhu/gorm"

	"github.com/gtforge/services_common_go/gett-storages"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPgBooksStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PgBooksStore suite")
}

var _ = Describe("Testing PgBooksStore", func() {
	var subject *pgBooksStore

	BeforeEach(func() {
		subject = &pgBooksStore{db: gettStorages.DB.New()}
	})

	AfterEach(func() {
		Expect(cleanDB()).To(Succeed())
		monkey.UnpatchAll()
	})

	Describe(".GetAll", func() {
		var (
			books []models.Book
			err   error
		)

		Context("when there is some error", func() {
			BeforeEach(func() {
				monkeyPatchDBError("Find")
			})

			It("should return an error", func() {
				books, err = subject.GetAll()

				Expect(err).To(HaveOccurred())
				Expect(books).To(BeZero())
			})
		})

		Context("when there are no books in DB", func() {
			It("should return empty slice without error", func() {
				books, err = subject.GetAll()

				Expect(err).To(Succeed())
				Expect(books).To(BeEmpty())
			})
		})

		Context("when there are books in DB", func() {
			BeforeEach(func() {
				books = []models.Book{
					{Name: "Test 1", Quantity: 1},
					{Name: "Test 2", Quantity: 2},
				}

				for _, book := range books {
					Expect(subject.Create(book.Name, book.Quantity)).To(Succeed())
				}
			})

			It("should return a slice of books without error", func() {
				books, err = subject.GetAll()

				Expect(err).To(Succeed())
				Expect(books[0].Name).To(BeEquivalentTo("Test 1"))
				Expect(books[0].Quantity).To(BeEquivalentTo(1))
				Expect(books[1].Name).To(BeEquivalentTo("Test 2"))
				Expect(books[1].Quantity).To(BeEquivalentTo(2))
			})
		})
	})

	Describe(".GetPerPage", func() {
		var (
			books      []models.Book
			err        error
			pageNumber int
		)

		Context("when there is some error", func() {
			BeforeEach(func() {
				monkeyPatchDBError("Find")
			})

			It("should return an error", func() {
				books, err = subject.GetPerPage(pageNumber)

				Expect(err).To(HaveOccurred())
				Expect(books).To(BeZero())
			})
		})

		Context("when there are no books in DB", func() {
			It("should return empty slice without error", func() {
				books, err = subject.GetPerPage(pageNumber)

				Expect(err).To(Succeed())
				Expect(books).To(BeEmpty())
			})
		})

		Context("when there are books in DB", func() {
			BeforeEach(func() {
				pageNumber = 1

				for i := 1; i < 20; i++ {
					Expect(subject.Create(fmt.Sprintf("Test %d", i), i)).To(Succeed())
				}
			})

			It("should return a slice of 10 books without error", func() {
				books, err = subject.GetPerPage(pageNumber)

				Expect(err).To(Succeed())
				Expect(len(books)).To(BeEquivalentTo(10))
			})
		})
	})
})

func cleanDB() error {
	return gettStorages.DB.Exec(`TRUNCATE TABLE public.books CONTINUE IDENTITY RESTRICT;`).Error
}

func monkeyPatchDBError(methodName string) {
	monkeyExtensions.PatchInstanceMethodFlexible(reflect.TypeOf(gettStorages.DB), methodName, func() *gorm.DB {
		return &gorm.DB{Error: errors.New(methodName + " failed")}
	})
}
