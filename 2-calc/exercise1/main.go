package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

var menu = map[string]func([]float64){
	"AVG": calculateAVG,
	"SUM": calculateSUM,
	"MED": calculateMED,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		operation, err := askOperation(reader) 
		if err != nil {
			fmt.Printf("Ошибка при вводе операции: %v\n", err)
			continue
		}
		var numbers []float64
		for {
			numbers = askNumbers(reader)
			if numbers != nil {
				break
			}
			fmt.Println("Попробуйте ввести числа снова.")
		}
		menu[operation](numbers)
		fmt.Print("Хотите выполнить еше одну операцию? (y/n): ")
		answer, err := reader.ReadString('\n')
		if err != nil || strings.TrimSpace(strings.ToLower(answer)) != "y" {
			fmt.Print("Спасибо за использование!")
			break
		}
	}
}

func askOperation(reader *bufio.Reader) (string, error) {
	fmt.Print("Введите операцию AVG, SUM, MED: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	operation := strings.TrimSpace(strings.ToUpper(input))
	if operation == "AVG" || operation == "SUM" || operation == "MED" {
		return operation, nil
	}
	fmt.Println("Ошибка: неизвестная операция. Доступны: AVG, SUM, MED")
	return askOperation(reader)
}

func askNumbers(reader *bufio.Reader) []float64 {
	fmt.Print("Введите числа через запятую: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	strNumbers := strings.Split(input, ",")
	var numbers []float64
	for _, str := range strNumbers {
		num, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
		if err != nil {
			fmt.Printf("Ошибка: '%s' не является числом\n", str)
			return nil
		}
		numbers = append(numbers, num)
	}
	if len(numbers) == 0 {
		fmt.Println("Ошибка: не введены числа")
		return nil
	}
	return numbers
}

func calculateAVG(numbers []float64) {
	if (len(numbers) == 0) {
		fmt.Println("Ошибка: невозможно вычислить среднее для пустого списка чисел")
		return 
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Среднее: %.2f\n", sum/float64(len(numbers)))
}

func calculateSUM(numbers []float64) {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	 fmt.Printf("Сумма: %.2f\n", sum)
}

func calculateMED(numbers []float64) {
	if len(numbers) == 0 {
		fmt.Println("Ошибка: невозможно вычислить медиану для пустого списка чисел")
        return
	}
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)
	n := len(sorted)
	if n%2 == 1 {
        fmt.Printf("Медиана: %.2f\n", sorted[n/2])
    } else {
        fmt.Printf("Медиана: %.2f\n", (sorted[n/2-1]+sorted[n/2])/2)
    }
}