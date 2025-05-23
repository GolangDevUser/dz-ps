package main

import (
	"fmt" 
	"strings"
)

var exchangeRates = map[string]map[string]float64 {
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
		from, err := inputCurrency()
		if err != nil {
			fmt.Println("Ошибка ввода исходной валюты", err)
			return
		}
	
		amount := getUserInput()
	
		to, err := inputCurrency()
		if err != nil {
			fmt.Println("Ошибка ввода целевой валюты:", err)
			return
		}
		result, err := calculate(amount, from, to)
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

func inputCurrency() (string, error) {
	var currency string
	for {
		fmt.Print("Введите валюту usd, eur, rub: ")
		_, err := fmt.Scan(&currency)
		if err != nil {
			return "", err
		}
		currency = strings.TrimSpace(strings.ToLower(currency))

		if _, exists := exchangeRates[currency]; exists {
			return currency, nil
		}
		fmt.Println("Неверная валюта. Попробуйте ещё раз")
	}
}

func calculate(amount int, from string, to string) (float64, error) {
    if from == to {
        return float64(amount), nil 
    }
    
    // Проверяем, поддерживается ли исходная валюта
    rates, ok := exchangeRates[from]
    if !ok {
        return 0, fmt.Errorf("исходная валюта '%s' не поддерживается", from)
    }
    
    // Проверяем, поддерживается ли целевая валюта
    rate, ok := rates[to]
    if !ok {
        return 0, fmt.Errorf("целевая валюта '%s' не поддерживается", to)
    }
    
    return float64(amount) * rate, nil
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