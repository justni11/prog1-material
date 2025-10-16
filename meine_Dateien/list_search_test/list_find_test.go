package listsearch

import "fmt"

func FindAll(l []int, x int) []int {
	result := []int{}

	for i := 0; i < len(l); i++ {
		if l[i] == x {
			result = append(result, i)
		}
	}
	return result
}

func Find(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}

func ExampleFind() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23, 5}

	pos1 := Find(l1, 42)
	pos2 := Find(l1, 200)

	fmt.Println(pos1)
	fmt.Println(pos2)

	//Output:
	// 2
	// -1

}

func ExampleFindAll() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23, 5}

	pos1 := FindAll(l1, 42)
	pos2 := FindAll(l1, 200)

	fmt.Println(pos1)
	fmt.Println(pos2)

	//Output:
	// [2]
	// []
}
