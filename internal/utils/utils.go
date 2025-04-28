package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const NewLine = '\n'

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

		if numberOption > 6 {
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

func GetInt(num int, someone string) int {
	for {
		ui := bufio.NewReader(os.Stdin)
		str, err := ui.ReadString(NewLine)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода\nПопробуйте еще раз")
			continue
		}
		str = strings.TrimSpace(str)

		if (someone == "year" || someone == "price") && str == "-" && num == 4 {
			return 0
		}

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

func GetString(num int) string {
	for {
		ui := bufio.NewReader(os.Stdin)
		str, err := ui.ReadString(NewLine)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода")
			continue
		}
		str = strings.TrimSpace(str)

		if len(str) == 0 {
			fmt.Fprint(os.Stderr, "Поле не может оставаться пустым\nПопробуйте еще раз: ")
			continue
		}

		if str == "-" && num != 4 {
			fmt.Fprint(os.Stdout, "Введено некорректное значение\nПопробуйте еще раз: ")
			continue
		}

		return str
	}
}

func IdToKey(ids []string) []string {
	if len(ids) == 0 {
		return nil
	}

	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = "book:id:" + id
	}

	return keys
}
