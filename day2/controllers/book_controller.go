package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/muhammad-zulfa/alterra-agmc/models"
	"net/http"
	"strconv"
)

var books []models.Book

type BookController interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type BookControllerImpl struct{}

func NewBookController() BookController {
	return &BookControllerImpl{}
}

func (controller *BookControllerImpl) Delete(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if len(books) < bookId {
		return c.NoContent(http.StatusNotFound)
	}

	var newBooks []models.Book

	for i, book := range books {
		if i+1 == bookId {
			continue
		}

		newBooks = append(newBooks, book)
	}

	books = newBooks

	return c.JSON(http.StatusOK, books)
}

func (controller *BookControllerImpl) Update(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if len(books) < bookId {
		return c.NoContent(http.StatusNotFound)
	}

	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	books[bookId-1] = book

	return c.NoContent(http.StatusOK)
}

func (controller *BookControllerImpl) FindById(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if len(books) < bookId {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, books[bookId-1])
}

func (controller *BookControllerImpl) Create(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	books = append(books, book)
	return c.NoContent(http.StatusCreated)
}

func (controller *BookControllerImpl) FindAll(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}
