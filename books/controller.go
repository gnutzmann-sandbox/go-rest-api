package books

import (
	"github.com/kataras/iris/v12"
	"github.com/samber/lo"
	"go-rest-api/errors"
)

func ServeRoutes(app *iris.Application) {
	booksApi := app.Party("/books")
	{
		booksApi.Get("/", getAllBooks)
		booksApi.Get("/{title}", getBook)
		booksApi.Post("/", create)
	}
}

var books = []book{
	{"Mastering Concurrency in Go"},
	{"Go Design Patterns"},
	{"Black Cat Go"},
}

func getBook(ctx iris.Context) {
	title := ctx.Params().Get("title")

	searchedBook, ok := lo.Find(books, func(x book) bool {
		return x.Title == title
	})

	if !ok {
		ctx.StopWithProblem(iris.StatusNotFound, iris.NewProblem().Detail(errors.BOOKS_NOT_FOUND))
		return
	}

	err := ctx.JSON(searchedBook)
	if err != nil {
		println("Error when responding request")
	}
}

func getAllBooks(ctx iris.Context) {
	err := ctx.JSON(books)
	if err != nil {
		println("Error when responding request")
	}
}

func create(ctx iris.Context) {
	var newBook book
	err := ctx.ReadJSON(&newBook)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(errors.BODY_PARSER_ERROR).DetailErr(err))

		return
	}

	println("Received book: " + newBook.Title)
	books = append(books, newBook)
	ctx.StatusCode(iris.StatusCreated)
}
