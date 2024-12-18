package main

import (
	"fmt"
	"library/internal/book"
	"library/internal/show"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "\nДобро пожаловать в библиотеку. Просьба не шуметь")
	fmt.Fprintln(os.Stdout, "Выберите один из предложенных вариантов:")
	fmt.Fprintln(os.Stdout, "1 - ознакомиться со всеми книгами в библотеке")
	fmt.Fprintln(os.Stdout, "2 - ознакомиться с определенной книгой")
	fmt.Fprintln(os.Stdout, "3 - добавить новую книгу")
	fmt.Fprintln(os.Stdout, "4 - удалить книгу")
	fmt.Fprintln(os.Stdout, "5 - уйти")
	fmt.Fprintln(os.Stdout)

	for {
		numberOption := show.ChooseOption()

		switch numberOption {
		case 1:
			show.ShowAll()
		case 2:
			fmt.Fprintln(os.Stdout, "Укажите название той книги, которая вас интересует")
			fmt.Fprintln(os.Stdout)
			title := show.ChooseTitleBook()
			show.ShowOne(title)
		case 3:
			fmt.Fprintln(os.Stdout)
			fmt.Fprintln(os.Stdout, "Этап добовления книги в общий список")

			fmt.Fprint(os.Stdout, "Введите название книги: ")
			name := show.GetString()
			fmt.Fprintln(os.Stdout, "Название успешно сохранено!")
			fmt.Fprintln(os.Stdout)

			fmt.Fprint(os.Stdout, "Введите имя автора книги: ")
			autor := show.GetString()
			fmt.Fprintln(os.Stdout, "Автор книги успешно сохранен!")
			fmt.Fprintln(os.Stdout)

			fmt.Fprint(os.Stdout, "Введите год издания книги: ")
			year := show.GetInt(numberOption)
			fmt.Fprintln(os.Stdout, "Дата успешно сохранена!")
			fmt.Fprintln(os.Stdout)

			fmt.Fprint(os.Stdout, "Введите цену книги (в рублях): ")
			price := show.GetInt(numberOption)
			fmt.Fprintln(os.Stdout, "Цена успешно сохранена!")
			fmt.Fprintln(os.Stdout)

			newBook := book.NewBook(name, autor, year, price)
			book.MyBooks.CreateBook(newBook)

			fmt.Fprintln(os.Stdout)
		case 4:
			fmt.Fprintln(os.Stdout)
			fmt.Fprintln(os.Stdout, "Этап удаления книги из списка")
			fmt.Fprint(os.Stdout, "Введите название той книги, которую хотите удалить из списка: ")
			title := show.GetString()
			book.MyBooks.RemoveBook(title)
		case 5:
			return
		}
	}
}
