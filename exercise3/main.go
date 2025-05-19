package main

import (
	"fmt" 
	"strings"
	"errors"
)

var exchangeRates = &map[string]map[string]float64 {
	"usd": {
		"eur": 0.88,
		"rub": 82.10,
	},
	"eur": {
		"usd": 1 / 0.88,
		"rub": 82.10 / 0.88,
	},
	"rub": {
		"usd": 1 / 82.10,
		"eur": 0.88 / 82.10,
	},
}

func main() {
	for {
		from, err := inputCurrency(exchangeRates)
		if err != nil {
			fmt.Println("Ошибка ввода исходной валюты", err)
			return
		}
	
		amount := getUserInput()
	
		to, err := inputCurrency(exchangeRates)
		if err != nil {
			fmt.Println("Ошибка ввода целевой валюты:", err)
			return
		}
		result, err := calculate(amount, from, to, exchangeRates)
		if err != nil {
			fmt.Println("Ошибка конвертации:", err)
			return
		}
	
		fmt.Printf("%d %s = %.2f %s\n", amount, from, result, to)
		
		if !askToContinue() {
			break
		}
	}
}

func getUserInput() int {
	var userEntered int
	for {
		fmt.Print("Введите количество: ")
		_, err := fmt.Scan(&userEntered)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Ошибка: нужно ввести число: ")
			continue
		}
		if userEntered < 0 {
			fmt.Println("Ошибка: число не может быть отрицательным: ")
			continue
		}
		return userEntered
	}
}

func inputCurrency(rates *map[string]map[string]float64) (string, error) {
	var currency string
	for {
		fmt.Print("Введите валюту usd, eur, rub: ")
		_, err := fmt.Scan(&currency)
		if err != nil {
			return "", err
		}
		currency = strings.TrimSpace(strings.ToLower(currency))

		if _,exists := (*rates)[currency]; exists {
			return currency, nil
		}
		fmt.Println("Неверная валюта, поддерживаемая валюта eur, usd, rub. Попробуйте ещё раз")
	}
}

func calculate(amount int, from string, to string, rates *map[string]map[string]float64) (float64, error) {
	if from == to {
		return float64(amount), nil 
	}
	if fromRates, ok := (*rates)[from]; ok {
		if rate, ok := fromRates[to]; ok {
			return float64(amount) * rate, nil
		}
	}
	return 0,errors.New("конвертация не поддерживается")
}

func askToContinue() bool {
    for {
        fmt.Print("Хотите выполнить еше одну операцию? (y/n): ")
        var answer string
        _, err := fmt.Scan(&answer)
        if err != nil {
            fmt.Println("Ошибка ввода. Попробуйте снова.")
            continue
        }
        answer = strings.TrimSpace(strings.ToLower(answer))
        if answer == "y" {
            return true
        }
        if answer == "n" {
            return false
        }
        fmt.Println("Ошибка ввода. Введите 'y' для продолжения или 'n' для выхода.")
    }
}