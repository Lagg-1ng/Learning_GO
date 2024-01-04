package main

import (
	"fmt"
)

type Mine struct {
	A, B, P, Q, AX, BY float64
}

func calculateMaxIncome(x, y, years float64, mines []Mine) float64 {
	if len(mines) == 0 {
		return 0
	}

	maxIncome := 0.0

	for year := 1.0; year <= years; year++ {
		for _, mine := range mines {
			chanceBreak := 1 - mine.P

			// income считается для каждой шахты и сравнивается с максимальной на текущий год и считается
			var income float64
			if (mine.A*x+(1-mine.Q)) > (mine.B*y*(1-mine.B)) {
				income = chanceBreak*mine.A*x + (1-mine.Q)*mine.B*y
				fmt.Print(" First mine\n")
			} else {
				income = chanceBreak * (mine.B*y * (1 - mine.B))
				fmt.Print(" Second mine\n")
			}

			if income > maxIncome {
				maxIncome = income
			}
		}
	}

	return maxIncome * years
}

func main() {
	var x, y, years float64
	fmt.Print("Enter the value for x: ")
	fmt.Scan(&x)
	fmt.Print("Enter the value for y: ")
	fmt.Scan(&y)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	var n int
	fmt.Print("Enter the number of mines: ")
	fmt.Scan(&n)

	var mines []Mine

	for i := 0; i < n; i++ {
		fmt.Printf("Enter the values for mine %d:\n", i+1)
		var mine Mine
		fmt.Print("  Draga b (should be in range (0, 1)): ")
		fmt.Scan(&mine.B)
		fmt.Print("  Draga a (should be in range (0, 1)): ")
		fmt.Scan(&mine.A)
		fmt.Print("  Draga q (should be in range [0, 1]): ")
		fmt.Scan(&mine.Q)
		fmt.Print("  Draga p (should be in range [0, 1]): ")
		fmt.Scan(&mine.P)
		//fmt.Print("  Draga ax: ")
		//fmt.Scan(&mine.AX)
		//fmt.Print("  Draga by: ")
		//fmt.Scan(&mine.BY)

		mines = append(mines, mine)
	}

	result := calculateMaxIncome(x, y, years, mines)
	fmt.Printf("The maximum income after %.0f years is: %.6f\n", years, result)
}


/*
Программа запрашивает входные данные, такие как значения x и y, количество лет (years), и количество шахт (n).
Далее, в цикле for i := 0; i < n; i++, программа запрашивает параметры для каждой шахты. Эти параметры включают B, A, Q, P, AX, и BY.
Внутри основной функции calculateMaxIncome, на каждом году (for year := 1.0; year <= years; year++), программа проходит по всем шахтам в цикле for _, mine := range mines.
Для каждой шахты рассчитывается доход с учетом вероятности поломки и других параметров. Это делается по формуле: income := (chanceBreak*mine.AX + (1-mine.Q)*mine.BY) * (1 - mine.B).
После рассчета дохода для каждой шахты на текущем году, программа проверяет, является ли этот доход максимальным. Если да, то обновляется переменная maxIncome.
*/
