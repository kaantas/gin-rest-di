package repository

import (
	"encoding/json"
	"gin-rest-di/model"
	"github.com/emirpasic/gods/lists/arraylist"
)

func GetBooks() *[]model.Book {
	list := arraylist.New()

	book := model.Book{
		ID:    114241,
		Title: "Inferno",
		Author: &model.Author{
			FirstName: "Dan",
			LastName:  "Brown",
		},
	}

	book2 := model.Book{
		ID:    5646654,
		Title: "Harry Potter",
		Author: &model.Author{
			FirstName: "J.K.",
			LastName:  "Rowling",
		},
	}

	list.Add(book)
	list.Add(book2)

	json2, err := list.ToJSON()
	if err != nil {
	}

	bookListJson := new([]model.Book)
	err2 := json.Unmarshal(json2, &bookListJson)

	if err2 != nil {
	}
	return bookListJson
}

func GetBooksList() *arraylist.List {
	list := arraylist.New()

	book := model.Book{
		ID:    114241,
		Title: "Inferno",
		Author: &model.Author{
			FirstName: "Dan",
			LastName:  "Brown",
		},
	}

	book2 := model.Book{
		ID:    5646654,
		Title: "Harry Potter",
		Author: &model.Author{
			FirstName: "J.K.",
			LastName:  "Rowling",
		},
	}

	list.Add(book)
	list.Add(book2)

	return list
}
