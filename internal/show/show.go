package show

import (
	"context"
	"encoding/json"
	"fmt"
	"library/internal/book"
	"library/internal/utils"
	"os"
	"sort"

	"github.com/redis/go-redis/v9"
)

func localShow(b book.Book) {
	fmt.Fprintf(os.Stdout, "\nКНИГА №%d\n", b.Id)
	fmt.Fprintf(os.Stdout, "Название: %s\n", b.Name)
	fmt.Fprintf(os.Stdout, "Автор: %s\n", b.Author)
	fmt.Fprintf(os.Stdout, "Год издания: %d\n", b.Year)
	fmt.Fprintf(os.Stdout, "Цена (в рублях): %d\n\n", b.Price)
}

func ShowAll(ctx context.Context, r *redis.Client) {

	ids, err := r.SMembers(ctx, "books:index").Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка получения индексов", err)
		return
	}

	booksJSON, err := r.MGet(ctx, utils.IdToKey(ids)...).Result()
	if err != nil {
		fmt.Fprint(os.Stderr, "\nКниг нет\n\n")
		return
	}

	var bookList []book.Book
	for _, bookJSON := range booksJSON {
		var b book.Book
		if bookJSON == nil {
			continue
		}
		if err := json.Unmarshal([]byte(bookJSON.(string)), &b); err != nil {
			fmt.Fprintln(os.Stdout, "Ошибка декодирования книги", err)
			continue
		}
		bookList = append(bookList, b)
	}

	sort.Slice(bookList, func(i, j int) bool {
		return bookList[i].Id < bookList[j].Id
	})

	fmt.Fprintln(os.Stdout, "\nСПИСОК ВСЕХ КНИГ ИЗ БИБЛИОТЕКИ")
	for _, b := range bookList {
		localShow(b)
	}
	fmt.Fprintln(os.Stdout)
}

func ShowOne(ctx context.Context, r *redis.Client, title string) {
	id, err := r.Get(ctx, "book:title:"+title).Result()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения ID книги\n%v\n\n", err)
		return
	}

	bookJSON, err := r.Get(ctx, "book:id:"+id).Result()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения книги\n%v\n\n", err)
		return
	}

	var book book.Book
	if err := json.Unmarshal([]byte(bookJSON), &book); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка десериализации книги\n%v\n\n", err)
		return
	}

	localShow(book)
}
