package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Простой калькулятор на Go")
	fmt.Println("Доступные операции: +, -, *, /")
	fmt.Println("Введите выражение (например: 2 + 3 или 5 * 4.2):")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Выход из калькулятора")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Ошибка: введите выражение в формате 'число оператор число'")
			continue
		}

		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Ошибка: первое значение не является числом")
			continue
		}

		operator := parts[1]

		num2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("Ошибка: второе значение не является числом")
			continue
		}

		var result float64
		var errorMessage string

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				errorMessage = "Ошибка: деление на ноль"
			} else {
				result = num1 / num2
			}
		default:
			errorMessage = "Ошибка: неизвестный оператор"
		}

		if errorMessage != "" {
			fmt.Println(errorMessage)
		} else {
			fmt.Printf("Результат: %.2f\n", result)
		}
	}
}