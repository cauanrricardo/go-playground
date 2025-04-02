package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 20, 50}
	fmt.Println(intArr)
	fmt.Println(intArr[2:3])

	//pode usar das duas formas (:=)
	frutas := [3]string{"maca", "banana", "pera"}
	fmt.Println(frutas)

	//pode usar [...] que o propio compialdor define a quantidade do array
	num := [...]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(num)
	fmt.Println(len(num))

}
