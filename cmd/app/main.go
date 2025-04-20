package main

import (
	"context"
	"fmt"
	"library/internal/book"
	"library/internal/database"
	"library/internal/show"
	"library/internal/utils"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rdb := database.InitDB()
	defer rdb.Close()

	fmt.Fprintln(os.Stdout, "\nДобро пожаловать в библиотеку. Просьба не шуметь")
	fmt.Fprintln(os.Stdout, "Выберите один из предложенных вариантов:")
	fmt.Fprintln(os.Stdout, "1 - ознакомиться со всеми книгами в библотеке")
	fmt.Fprintln(os.Stdout, "2 - ознакомиться с определенной книгой")
	fmt.Fprintln(os.Stdout, "3 - добавить новую книгу")
	fmt.Fprintln(os.Stdout, "4 - изменить данные книги")
	fmt.Fprintln(os.Stdout, "5 - удалить книгу")
	fmt.Fprintln(os.Stdout, "6 - уйти")
	fmt.Fprintln(os.Stdout)

	for {
		numberOption := utils.ChooseOption()

		switch numberOption {
		case 1:
			show.ShowAll(ctx, rdb)
		case 2:
			fmt.Fprintln(os.Stdout, "Укажите название той книги, которая вас интересует")
			fmt.Fprintln(os.Stdout)
			title := utils.ChooseTitleBook()
			show.ShowOne(ctx, rdb, title)
		case 3:
			fmt.Fprintln(os.Stdout)
			fmt.Fprintln(os.Stdout, "Этап добовления книги в общий список")

			newBook := book.FillInFields(3)
			book.CreateBook(ctx, rdb, newBook)

			fmt.Fprintln(os.Stdout)
		case 4:
			fmt.Fprintln(os.Stdout)
			fmt.Fprintln(os.Stdout, "Этап изменения книги (введите \"-\", чтобы оставить предыдущие данные)")
			fmt.Fprintf(os.Stdout, "Введите название той книги, которую хотите изменить: ")
			title := utils.GetString(4)
			fmt.Fprintln(os.Stdout)

			book.UpdateBook(ctx, rdb, title, 4)
		case 5:
			fmt.Fprintln(os.Stdout)
			fmt.Fprintln(os.Stdout, "Этап удаления книги из списка")
			fmt.Fprint(os.Stdout, "Введите название той книги, которую хотите удалить из списка: ")
			title := utils.GetString(5)
			fmt.Fprintln(os.Stdout)

			book.RemoveBook(ctx, rdb, title)
		case 6:
			return
		}
	}
}
