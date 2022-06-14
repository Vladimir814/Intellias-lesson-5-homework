package main

import  "fmt" 


// вхідні дані
var (
    myMoney int = 23 
	applePrice float64 = 5.99
	pearPrice int = 7
)

func main() {
    sum1 := (applePrice * float64(9)) + float64(pearPrice * 8) 
	howManyPear := myMoney / pearPrice 
	howManyApple := int(float64(myMoney) / applePrice )
	sum2 := (float64(pearPrice) + applePrice) * 2
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? : %.2f грн\n", sum1)
    fmt.Printf("2. Скільки груш ми можемо купити? : %d шт\n", howManyPear)
    fmt.Printf("3. Скільки яблук ми можемо купити? : %d шт\n", howManyApple)
    fmt.Printf("4. Чи ми можемо купити 2 груші та 2 яблука? : НІ, бо нам треба %.2f грн\n, а ми маємо лише %d грн\n", sum2, myMoney)
}
