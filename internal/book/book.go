package book

import (
	"encoding/json"
	"fmt"
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

func (books *BookStore) CreateBook(fileName string, book Book) {
	ReadJsonFile(fileName, books)

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

	dataJson, err := json.Marshal(*books)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при записи в JSON формат")
		return
	}
	err = os.WriteFile(fileName, dataJson, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при записи данных в файл")
	}

	fmt.Fprintf(os.Stdout, "Книга успешно сохранена под номером %d и добавлена в общий список!\n", book.Id)
}

func (books *BookStore) RemoveBook(fileName string, title string) {
	ReadJsonFile(fileName, books)

	for i, val := range books.Books {
		if val.Name == title {
			books.Books = append(books.Books[:i], books.Books[i+1:]...)

			for i := 0; i < len(books.Books); i++ {
				books.Books[i].Id = i + 1
			}

			dataJson, err := json.Marshal(*books)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при записи в JSON формат")
			}
			err = os.WriteFile(fileName, dataJson, 0666)
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

func CreatFile(filePath string) (*os.File, error) {
	var file *os.File

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при создании файла")
			return nil, err
		}
	} else {
		file, err = os.OpenFile(filePath, os.O_RDWR, 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при открытии файла")
			return nil, err
		}
	}
	defer file.Close()

	return file, nil
}

func ReadJsonFile(fileName string, bookStore *BookStore) {
	dataJson, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении файла")
		return
	}

	if len(dataJson) == 0 {
		return
	}

	err = json.Unmarshal(dataJson, bookStore)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении файла: %v\n", err)
		return
	}
}
