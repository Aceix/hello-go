package main

import "fmt"

func main() {
	var num int

start:
	fmt.Printf("Enter number of scores (min of 3): ")
	fmt.Scanf("%d", &num)

	if num < 3 {
		goto start
	}

	var scores []float64 = make([]float64, num)
	var sum float64 = 0.0

	fmt.Println()

	for i := 0; i < len(scores); i++ {
		fmt.Println("Enter your scores below...\n")
		fmt.Printf("Score #%d: ", i+1)
		fmt.Scanf("%f", &scores[i])
		sum += scores[i]
	}

	fmt.Println("Your average score is:", fmt.Sprintf("%.2f", sum/float64(len(scores))))

	fmt.Println("Scores array has a capacity of:", cap(scores))
	fmt.Println("Scores array has a lenght of:", len(scores))
	fmt.Println("First 3 elements are:", scores[:3]) // slices are similar to python's

	scores = append(scores, 100)
	fmt.Println("New slice:", scores)

	var copySlice []float64
	copySlice = make([]float64, num)
	copy(copySlice, scores)

	fmt.Println("Copy slice:", copySlice)

	for i, x := range copySlice {
		fmt.Printf("Score #%d: %.2f\n", i+1, x)
	}

	// MAPS in Go
	var name string

	fmt.Println()
	fmt.Println("========================")
	fmt.Println("Save scores under a name\n")
	fmt.Printf("Please enter your name: ")
	fmt.Scanf("%s", &name)

	var db map[string][]float64
	var dbPtr *map[string][]float64

	db = make(map[string][]float64)
	db[name] = scores
	dbPtr = &db

	(*dbPtr)[name][0] = 0

	fmt.Println("Modded db: ", *dbPtr)
	
	delete(db, name)

	fmt.Println("Deleted db entry: ", *dbPtr)
}
