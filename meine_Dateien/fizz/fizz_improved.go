package fizz

import "fmt"

//Fizz gibt die Zahlen von 1 bis 30 aus und ersetzt dabei
//jede durch 3 teilbare Zahl durch "fizz" und jede
//durch 5 teilbare Zahl durch "buzz"
//bei Zahlen, die durch beide Zahlen teilbar sind wird "fizzbuzz" ausgegeben
func FizzImproved(x, y, n int) {
	for i := 1; i <= n; i++ {
		//wenn i durch 3 und 5 teilbar ist, gib "fizzbuzz" aus
		if i%y == 0 && i%x == 0 {
			fmt.Println("fizzbuzz")
		} else
		//wenn i durch 3 teilbar ist, gib "fizz" aus
		if i%x == 0 {
			fmt.Println("fizz")
			//fmt.Printf("%10d:%v\n", i, "fizz")
		} else
		//wenn i durch 5 teilbar ist, gib "buzz" aus
		if i%y == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
		//sonst gib i aus
	}
}
