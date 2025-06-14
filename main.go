package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Difficulty struct {
	Mode      string
	MaxNumber int
	Attempts  int
}

func inputNumber(attempt, maxNumber int) (int, bool) {
	fmt.Printf("Попытка %d: Введите число: ", attempt)
	var userNumber int
	_, err := fmt.Scanln(&userNumber)
	if err != nil {
		fmt.Printf("❌ Введите число от 1 до %d.\n", maxNumber)
		var clear string
		fmt.Scanln(&clear)
		return 0, false
	}

	if userNumber < 1 || userNumber > maxNumber {
		fmt.Printf("❌ Введите число от 1 до %d.\n", maxNumber)
		return 0, false
	}

	return userNumber, true
}

func generateRandomNumber(maxNumber int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNumber) + 1
}

func compareAndHint(userNumber, randomNumber int) bool {
	if userNumber == randomNumber {
		return true
	}

	if randomNumber > userNumber {
		if randomNumber-userNumber <= 5 {
			fmt.Println("🔥 Горячо")
		} else if randomNumber-userNumber <= 15 {
			fmt.Println("🙂 Тепло")
		} else {
			fmt.Println("❄️ Холодно")
		}
		fmt.Println("Секретное число больше 👆")
	} else {
		if userNumber-randomNumber <= 5 {
			fmt.Println("🔥 Горячо")
		} else if userNumber-randomNumber <= 15 {
			fmt.Println("🙂 Тепло")
		} else {
			fmt.Println("❄️ Холодно")
		}
		fmt.Println("Секретное число меньше 👇")
	}
	return false
}

func selectDifficulty(difficulties []Difficulty) (Difficulty, bool) {
	fmt.Println("Выберите сложность:")
	for i, d := range difficulties {
		fmt.Printf("%d. %s (1–%d, %d попыток)\n", i+1, d.Mode, d.MaxNumber, d.Attempts)
	}
	fmt.Print("Введите номер сложности (1–3): ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > 3 {
		fmt.Println("❌ Введите число от 1 до 3.")
		var clear string
		fmt.Scanln(&clear)
		return Difficulty{}, false
	}

	return difficulties[choice-1], true
}

func playGame(difficulty Difficulty) {
	fmt.Printf("Игра 'Угадай число' — от 1 до %d началась!\n", difficulty.MaxNumber)
	fmt.Printf("Угадайте число за %d попыток!\n", difficulty.Attempts)

	randomNumber := generateRandomNumber(difficulty.MaxNumber)
	attempt := 1
	rememberNum := make([]int, 0, difficulty.Attempts)

	for attempt <= difficulty.Attempts {
		userNumber, ok := inputNumber(attempt, difficulty.MaxNumber)
		if !ok {
			continue
		}

		rememberNum = append(rememberNum, userNumber)
		fmt.Printf("Введенные ранее вами числа: %v\n", rememberNum)

		if compareAndHint(userNumber, randomNumber) {
			if attempt == 1 {
				fmt.Printf("🎉 Вы победили! Использовав всего %d попытку\n", attempt)
			} else if attempt <= 4 {
				fmt.Printf("🎉 Вы победили! Использовав всего %d попытки\n", attempt)
			} else {
				fmt.Printf("🎉 Вы победили! Использовав %d попыток\n", attempt)
			}
			return
		}

		attempt++
	}

	fmt.Printf("😢 Ваши попытки закончились! Секретное число было: %d\n", randomNumber)
}

func main() {

	difficulties := []Difficulty{
		{Mode: "Easy", MaxNumber: 50, Attempts: 15},
		{Mode: "Medium", MaxNumber: 100, Attempts: 10},
		{Mode: "Hard", MaxNumber: 200, Attempts: 5},
	}

	for {

		difficulty, ok := selectDifficulty(difficulties)
		if !ok {
			continue
		}

		playGame(difficulty)

		fmt.Println("Игра закончена!")

		fmt.Print("Хотите сыграть ещё раз? (y/n): ")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			var clear string
			fmt.Scanln(&clear)
			continue
		}

		if response != "y" && response != "Y" {
			break
		}
	}

	fmt.Println("Спасибо за игру!")
}
