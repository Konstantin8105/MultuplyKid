package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var amountTasks = 20

type numbers struct {
	values [2]int
}

func calculuteNumber(n numbers) int {
	return n.values[0] * n.values[1]
}

var other []numbers

func generateNumber() (n numbers) {
	n.values[0] = 7 + rand.Intn(3)
	n.values[1] = rand.Intn(11)
	for _, nn := range other {
		if nn.values[0] == n.values[0] && nn.values[1] == n.values[1] {
			return generateNumber()
		}
	}
	other = append(other, n)
	return
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	var summaryResult []bool

	for i := 0; i < amountTasks; i++ {
		n := generateNumber()
	AGAIN:
		// out
		fmt.Printf("\n\n_______________\n\n")
		fmt.Println("Пример №", i+1, " из ", amountTasks)
		fmt.Printf("Дан следующий пример:\n")
		fmt.Printf("%v * %v\n", n.values[0], n.values[1])
		fmt.Println("Реши и запиши ответ:")
		// get result
		var result int
		var line string
		_, err := fmt.Scanf("%s", &line)
		if err != nil {
			fmt.Println("\nНе понятный ответ")
			//fmt.Println("Неверный ответ. Попробуй снова")
			//fmt.Println("t = ", t)
			//fmt.Println("err1 = ", err)
			goto AGAIN
		}
		r, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("\nНе понятный ответ")
			//fmt.Println("Неверный ответ. Попробуй снова")
			//fmt.Println("err2 = ", err)
			goto AGAIN
		}
		result = int(r)
		// output fail or ok
		if result != calculuteNumber(n) {
			fmt.Println("Неверный ответ. Попробуй снова")
			summaryResult = append(summaryResult, false)
			goto AGAIN
		}
		fmt.Println("Правильно")
		summaryResult = append(summaryResult, true)
	}

	// results
	positive := 0
	for i := range summaryResult {
		if summaryResult[i] {
			positive++
		}
	}
	fmt.Println("\n\n_______________")
	fmt.Println("Количество правильных ответов : ", positive)
	fmt.Println("Количество неправильных ответов : ", len(summaryResult)-positive)

}
