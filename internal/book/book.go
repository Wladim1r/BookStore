package book

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Price  int    `json:"price"`
}

func NewBook(name, author string, year, price int) Book {
	return Book{
		Name:   name,
		Author: author,
		Year:   year,
		Price:  price,
	}
}

type BookStore struct {
	Books []Book `json:"books"`
	MaxId int    `json:"max_id"`
}

func (books *BookStore) CreateBook(book Book) {
	file, err := os.OpenFile("cmd/app/allbooks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при открытии файла")
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(books)
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении данных из файла")
		return
	}

	books.MaxId = 0

	if len(books.Books) == 0 {
		books.MaxId = 1
	}

	for _, book := range books.Books {
		if book.Id >= books.MaxId {
			books.MaxId = book.Id + 1
		}
	}

	book.Id = books.MaxId

	books.Books = append(books.Books, book)

	file.Seek(0, 0)

	if err := file.Truncate(0); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при отчистке файла")
		return
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(books)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при записи данных в файл")
		return
	}

	fmt.Fprintf(os.Stdout, "Книга успешно сохранена под номером %d и добавлена в общий список!\n", book.Id)
}

func (books *BookStore) RemoveBook(title string) {

	file, err := os.OpenFile("cmd/app/allbooks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при открытии файла")
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&books)
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении данных из файла")
	}

	for i, val := range books.Books {
		if val.Name == title {
			books.Books = append(books.Books[:i], books.Books[i+1:]...)

			for i := 0; i < len(books.Books); i++ {
				books.Books[i].Id = i + 1
			}

			file.Seek(0, 0)

			if err := file.Truncate(0); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при очистке файла")
			}

			encoder := json.NewEncoder(file)
			err = encoder.Encode(books)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при записи данных в файл")
			}

			fmt.Fprintf(os.Stdout, "Книга \"%s\" успешно удалена из списка!\n", val.Name)
			fmt.Fprintln(os.Stdout)

			return
		}
	}
	fmt.Fprintln(os.Stdout, "Книга по данному названию не найдена")
	fmt.Fprintln(os.Stdout)
}

var MyBooks = BookStore{}
