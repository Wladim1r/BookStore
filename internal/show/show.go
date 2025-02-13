package show

import (
	"bufio"
	"fmt"
	"library/internal/book"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const NewLine = '\n'

func localShow(id int, name, autor string, year, price int) {
	fmt.Fprintf(os.Stdout, "\nКНИГА №%d\n", id)
	fmt.Fprintf(os.Stdout, "Название: %s\n", name)
	fmt.Fprintf(os.Stdout, "Автор: %s\n", autor)
	fmt.Fprintf(os.Stdout, "Год издания: %d\n", year)
	fmt.Fprintf(os.Stdout, "Цена (в рублях): %d\n\n", price)
}

func ShowAll(fileName string) {
	var bookStore book.BookStore
	book.ReadJsonFile(fileName, &bookStore)
	books := bookStore.Books

	sort.Slice(books, func(i, j int) bool {
		return books[i].Id < books[j].Id
	})

	if len(books) == 0 {
		fmt.Fprintf(os.Stdout, "\nКниг нет\n\n")
	} else {
		fmt.Fprintf(os.Stdout, "\nСПИСОК ВСЕХ КНИГ ИЗ СПИСКА\n")
		for _, val := range books {
			localShow(val.Id, val.Name, val.Author, val.Year, val.Price)
		}
	}
}

func ShowOne(fileName string, title string) {
	var bookStore book.BookStore
	book.ReadJsonFile(fileName, &bookStore)
	books := bookStore.Books

	for _, val := range books {
		if val.Name == title {
			localShow(val.Id, val.Name, val.Author, val.Year, val.Price)
			return
		}
	}
	fmt.Fprintln(os.Stderr, "Такой книги нет")
	fmt.Fprintln(os.Stdout)
}

func ChooseOption() int {
	var numberOption int

	for {
		fmt.Fprint(os.Stdout, "Поле для ввода действия над библиотекой: ")

		ui := bufio.NewReader(os.Stdin)
		str, err := ui.ReadString(NewLine)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода\nПопробуйте еще раз")
			fmt.Fprintln(os.Stdout)
			continue
		}

		str = strings.TrimSpace(str)

		if len(str) == 0 {
			fmt.Fprintln(os.Stderr, "Нельзя оставлять поле пустым")
			fmt.Fprintln(os.Stdout)
			continue
		}

		numberOption, err = strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ввод не может содержать какие-либо символы, кроме цифр")
			fmt.Fprintln(os.Stdout)
			continue
		}

		if numberOption > 5 {
			fmt.Fprintln(os.Stderr, "Число слишком большое")
			fmt.Fprintln(os.Stdout)
			continue
		}
		if numberOption < 1 {
			fmt.Fprintln(os.Stderr, "Число слишком маленькое")
			fmt.Fprintln(os.Stdout)
			continue
		}
		break
	}
	return numberOption
}

func ChooseTitleBook() string {
	for {
		fmt.Fprint(os.Stdout, "Поле для ввода названия книги: ")

		ui := bufio.NewReader(os.Stdin)
		str, err := ui.ReadString(NewLine)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода\nПопробуйте еще раз")
			fmt.Fprintln(os.Stdout)
			continue
		}
		str = strings.TrimSpace(str)

		if len(str) == 0 {
			fmt.Fprintln(os.Stderr, "Ввод не должен быть пустым")
			fmt.Fprintln(os.Stdout)
			continue
		}
		return str
	}
}

func GetInt(someone string) int {
	for {
		ui := bufio.NewReader(os.Stdin)
		str, err := ui.ReadString(NewLine)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода\nПопробуйте еще раз")
			continue
		}
		str = strings.TrimSpace(str)

		number, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprint(os.Stdout, "Ошибка ввода\nПопробуйте еще раз: ")
			continue
		}

		if number < 1 {
			fmt.Fprint(os.Stderr, "Число должно быть положительным\nПопробуйте еще раз: ")
			continue
		}

		if someone == "year" && number > time.Now().Year() {
			fmt.Fprint(os.Stderr, "Указанный вами год не совпадает с текущим\nПопробуйте еще раз: ")
			continue
		}

		return number
	}
}

func GetString() string {
	for {
		ui := bufio.NewReader(os.Stdin)
		str, err := ui.ReadString(NewLine)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода")
			continue
		}
		str = strings.TrimSpace(str)

		if len(str) == 0 {
			fmt.Fprint(os.Stderr, "Поле не может оставаться пустым\nПопробйте еще раз: ")
			continue
		}

		return str
	}
}
