package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	rom "github.com/brandenc40/romannumeral"
)

// Функция для проверки, является ли строка римским числом
func isRomanNumeral(s string) bool {
	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for _, r := range s {
		if _, ok := romanNumerals[string(r)]; !ok {
			return false
		}
	}

	return true
}


func main() {
	fmt.Println("Введите выражение:")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	parts := strings.Split(text, " ")
	if len(parts) != 3 {
		fmt.Println("Ошибка: неверный формат выражения")
		return
	}

	a := parts[0]
	operator := parts[1]
	b := parts[2]

	// Проверка, являются ли числа римскими или арабскими
	isRomanA := isRomanNumeral(a)
	isRomanB := isRomanNumeral(b)

	var result int
	if isRomanA && isRomanB {
		// Если оба числа римские, выполним операцию в римской системе
		arabicA, err := rom.StringToInt(a)
		if err != nil {
			panic(err)
		}
		if arabicA < 1 || arabicA > 10 {
			fmt.Println("Ошибка, калькулятор принимает числа от 1 до 10 включительно")
			return
		}
		arabicB, err := rom.StringToInt(b)
		if err != nil {
			panic(err)
		}
		if arabicB < 1 || arabicB > 10 {
			fmt.Println("Ошибка, калькулятор принимает числа от 1 до 10 включительно")
			return
		}
		switch operator {
		case "+":
			result = arabicA + arabicB
		case "-":
			result = arabicA - arabicB
		case "*":
			result = arabicA * arabicB
		case "/":
			if arabicB != 0 {
				result = arabicA / arabicB
			} else {
				fmt.Println("Ошибка: деление на ноль")
				return
			}
		default:
			fmt.Println("Ошибка: неподдерживаемый оператор")
			return
		}

		// Преобразуем результат обратно в римское число
		romanResult, err := rom.IntToString(result)
		if err != nil {
			fmt.Println("Ошибка, в римской системе нет отрицательных чисел.")
			return
		}
		fmt.Println("Результат:", romanResult)

	} else if !isRomanA && !isRomanB {
		// Если оба числа арабские, выполним операцию в арабской системе
		arabicA, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("Ошибка при преобразовании первого операнда:", err)
			return
		}
		if arabicA < 1 || arabicA > 10 {
			fmt.Println("Ошибка, калькулятор принимает числа от 1 до 10 включительно")
			return
		}

		arabicB, err := strconv.Atoi(b)
		if err != nil {
			fmt.Println("Ошибка при преобразовании второго операнда:", err)
			return
		}
		if arabicB < 1 || arabicB > 10 {
			fmt.Println("Ошибка, калькулятор принимает числа от 1 до 10 включительно")
			return
		}

		switch operator {
		case "+":
			result = arabicA + arabicB
		case "-":
			result = arabicA - arabicB
		case "*":
			result = arabicA * arabicB
		case "/":
			if arabicB != 0 {
				result = arabicA / arabicB
			} else {
				fmt.Println("Ошибка: деление на ноль")
				return
			}
		default:
			fmt.Println("Ошибка: неподдерживаемый оператор")
			return
		}

		fmt.Println("Результат:", result)

	} else {
		fmt.Println("Ошибка: операнды должны быть либо оба арабскими, либо оба римскими числами")
		return
	}
}

