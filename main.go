package main

import (
	"fmt"
	"math"
)

func main(){

	var (
	allMoney float64 = 23
	pearPrice float64 = 7
	applePrice float64 = 5.99
)
	
	fmt.Println("1.Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println("Відповідь: ", math.Ceil(((float64(9) * applePrice + float64(8) * pearPrice)) * 100) / 100)

	fmt.Println("2.Скільки груш ми можемо купити?")
	fmt.Println("Відповідь: ", math.Floor(allMoney/pearPrice))

	fmt.Println("3.Скільки яблук ми можемо купити?")
	fmt.Println("Відповідь: ", math.Floor(allMoney/applePrice))

	fmt.Println("4.Чи ми можемо купити 2 груші та 2 яблука?")

	if (2 * pearPrice + 2 * applePrice <= allMoney){
		fmt.Println("Відповідь: ", "Так")	
	} else {
		fmt.Println("Відповідь: ", "Ні")
	}

}