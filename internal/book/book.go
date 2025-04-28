package book

import (
	"context"
	"encoding/json"
	"fmt"
	"library/internal/utils"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Book struct {
	Id     int64  `json:"id"`
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

func CreateBook(ctx context.Context, r *redis.Client, book Book) {
	id, err := r.Incr(ctx, "books:last_id").Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка инкрементации ID", err)
	}
	book.Id = id

	bookJSON, err := json.Marshal(book)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка сериализации книги", err)
		return
	}

	// транзакция добавления книги и ее индекса
	_, err = r.TxPipelined(ctx, func(p redis.Pipeliner) error {
		p.Set(ctx, "book:id:"+strconv.Itoa(int(book.Id)), bookJSON, 0)
		p.Set(ctx, "book:title:"+book.Name, book.Id, 0)

		p.SAdd(ctx, "books:index", book.Id)

		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка создания книги", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Книга успешно сохранена под номером %d и добавлена в общий список!\n", book.Id)
}

func RemoveBook(ctx context.Context, r *redis.Client, title string) {
	id, err := r.Get(ctx, "book:title:"+title).Result()
	if err == redis.Nil {
		fmt.Fprintln(os.Stderr, "Книга не найдена", err)
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	_, err = r.Get(ctx, "book:id:"+id).Result()
	if err == redis.Nil {
		fmt.Fprintln(os.Stderr, "Книга не найдена", err)
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	_, err = r.TxPipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, "book:title:"+title)
		p.Del(ctx, "book:id:"+id)

		p.SRem(ctx, "books:index", id)

		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка удаления книги", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Книга успешно удалена\n\n")
}

func UpdateBook(ctx context.Context, r *redis.Client, title string, num int) {
	id, err := r.Get(ctx, "book:title:"+title).Result()
	if err == redis.Nil {
		fmt.Fprintf(os.Stderr, "Книга не найдена\n%v\n\n", err)
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	bookJSON, err := r.Get(ctx, "book:id:"+id).Result()
	if err == redis.Nil {
		fmt.Fprintln(os.Stderr, "Книга не найдена", err)
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var book Book
	if err := json.Unmarshal([]byte(bookJSON), &book); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка десериализации книги", err)
		return
	}

	newBook := FillInFields(num)

	newBook.Id = book.Id
	if newBook.Name == "-" {
		newBook.Name = book.Name
	}
	if newBook.Author == "-" {
		newBook.Author = book.Author
	}
	if newBook.Year == 0 {
		newBook.Year = book.Year
	}
	if newBook.Price == 0 {
		newBook.Price = book.Price
	}

	newBookJSON, err := json.Marshal(newBook)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка сериализации книги", err)
		return
	}

	_, err = r.TxPipelined(ctx, func(p redis.Pipeliner) error {
		p.Set(ctx, "book:id:"+id, newBookJSON, 0)

		if book.Name != newBook.Name {
			p.Del(ctx, "book:title:"+book.Name)
			p.Set(ctx, "book:title:"+newBook.Name, book.Id, 0)
		}

		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка изменения книги", err)
		return
	}

	fmt.Fprintln(os.Stdout, "Книга успешно обновлена")
}

func FillInFields(num int) Book {
	fmt.Fprint(os.Stdout, "Введите название книги: ")
	name := utils.GetString(num)
	if name == "-" {
		fmt.Fprintf(os.Stdout, "Название осталось прежним!\n\n")
	} else {
		fmt.Fprintf(os.Stdout, "Название успешно сохранено!\n\n")
	}

	fmt.Fprint(os.Stdout, "Введите имя автора книги: ")
	author := utils.GetString(num)
	if author == "-" {
		fmt.Fprintf(os.Stdout, "Автор остался прежним!\n\n")
	} else {
		fmt.Fprintf(os.Stdout, "Автор книги успешно сохранен!\n\n")
	}

	fmt.Fprint(os.Stdout, "Введите год издания книги: ")
	year := utils.GetInt(num, "year")
	if year == 0 {
		fmt.Fprintf(os.Stdout, "Год издания остался прежним!\n\n")
	} else {
		fmt.Fprintf(os.Stdout, "Дата успешно сохранена!\n\n")
	}

	fmt.Fprint(os.Stdout, "Введите цену книги (в рублях): ")
	price := utils.GetInt(num, "price")
	if price == 0 {
		fmt.Fprintf(os.Stdout, "Цена осталась прежней!\n\n")
	} else {
		fmt.Fprintf(os.Stdout, "Цена успешно сохранена!\n\n")
	}

	return NewBook(name, author, year, price)
}
