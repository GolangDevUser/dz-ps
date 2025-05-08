package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		operation, err := askOperation(reader) 
		if err != nil {
			fmt.Printf("Ошибка при вводе операции: %v\n", err)
			continue
		}
		numbers := askNumbers(reader) 
		if numbers == nil {
			continue
		}
		executeOperation(operation, numbers)
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

func executeOperation(operation string, numbers []float64) {
	switch operation {
	case "AVG":
		fmt.Printf("Среднее: %.2f\n", calculateAVG(numbers))
	case "SUM":
		fmt.Printf("Сумма: %.2f\n", calculateSUM(numbers))
	case "MED":
		fmt.Printf("Медиана: %.2f\n", calculateMED(numbers))
	}
}

func calculateAVG(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func calculateSUM(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateMED(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if n % 2 == 1 {
		return numbers[n/2]
	}
	return (numbers[n / 2 - 1] + numbers[n / 2]) / 2
}