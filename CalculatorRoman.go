package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	InputMessage()   // Отдельная функция создана для удобства, чтобы убрать громоздкий текст
	var input string // Задаем переменную для ввода строки
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		input = strings.ReplaceAll(input, " ", "")
		values := strings.Split(input, "")

		x := values[0] // Извлекаем числа и оператор выражения
		y := values[2] // К сожалению, не смог сделать извлечение чисел, состоящих из 2 цифр
		oper := values[1]
		FinalCalculate(x, y, oper) // Выносим все действия в другую функцию, чтобы не загромождать func main()
	}
}

func InputMessage() {
	fmt.Println("Введите арифметическое выражение, например, 'x+y'. " +
		"Цифры должны быть арабские или римские, от 1 до 10 включительно. " +
		"Также цифры в выражении должны быть целые.")
}

func FinalCalculate(x, y, oper string) {
	// Функция производит вычисление в зависимости арабской или римской цифры
	if NumInputArab(x) && NumInputArab(y) {
		CalculateArab(x, y, oper)
	} else if NumInputRoman(y) && NumInputRoman(y) {
		CalculateRoman(x, y, oper)
	} else {
		fmt.Println("В одном выражении не могут быть разные виды чисел.")
		os.Exit(0)
	}
}

func CalculateArab(x, y, oper string) {
	// Вычисляет арабские цифры
	num1, err := strconv.Atoi(x)
	num2, err := strconv.Atoi(y)
	if err != nil {
		fmt.Println(err)
	}
	if (num1 < 0 || num1 >= 10) || (num2 < 0 || num2 >= 10) {
		fmt.Println("Число должно быть в пределах от 1 до 10 включительно.")
		os.Exit(0)
	}
	result := Calculate(x, y, oper)
	fmt.Println("Полученный результат: " + strconv.Itoa(result))
}

func CalculateRoman(x, y, oper string) {
	// Вычисляет римские цифры
	num1 := strconv.Itoa(RomanToArab(x)) // RomanToArab возвращает int, а Calculate требует string
	num2 := strconv.Itoa(RomanToArab(y))
	res := Calculate(num1, num2, oper)

	if res <= 0 {
		fmt.Println("Не существует римских чисел меньше нуля.")
		os.Exit(0)
	}
	result := ArabToRoman(res)
	fmt.Println("Полученный результат: " + result)
}

func NumInputArab(numTaken string) (numCheck bool) {
	// Функция определяет, является ли введенное число арабским.
	numTakenInt, _ := strconv.Atoi(numTaken)
	if numTakenInt > 0 {
		numCheck = true
	} else {
		numCheck = false
	}
	return numCheck
}

func NumInputRoman(numTaken string) (numCheck bool) {
	for s := range romanArab {
		if s == numTaken {
			numCheck = true
			break
		}
	}
	return numCheck
}

var romanArab = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4,
	"V": 5, "VI": 6, "VII": 7, "VIII": 8,
	"IX": 9, "X": 10,
}

func RomanToArab(numRoman string) (numArab int) {
	// Конвертирует римские цифры в арабские, используя хеш-таблицу
	for key, value := range romanArab {
		if numRoman == key {
			numArab = value
		}
	}
	return numArab
}

func ArabToRoman(numArab int) string {
	// Функция принимает числовое значение и возвращает римский эквивалент, используется для вывода результата.
	roman := []string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV",
		"XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII",
		"XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI",
		"XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV",
		"LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX",
		"LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI",
		"LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII",
		"XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
	}
	numRoman := roman[numArab]
	return numRoman
}

func Calculate(x, y, oper string) int {
	// Функция для вычисления арифметических операций
	num1, _ := strconv.Atoi(x)
	num2, _ := strconv.Atoi(y)
	var res int
	if oper == "+" {
		res = num1 + num2
	} else if oper == "-" {
		res = num1 - num2
	} else if oper == "*" {
		res = num1 * num2
	} else if oper == "/" {
		res = num1 / num2
	} else {
		fmt.Println("Неверно введены данные")
	}
	return res
}
