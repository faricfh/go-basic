package main

import "fmt"

func main() {
	// thisArray := [...]int{1, 2, 3}
	// thisSlice := []int{1, 2, 3}

	names := []string{"John", "Jane", "Doe", "Smith", "Emily", "Michael"}
	slice := names[3:5]

	fmt.Println("Names:", names)
	fmt.Println("Slice:", slice)

	slice2 := names[:4]
	fmt.Println("Slice2:", slice2)

	slice3 := names[2:]
	fmt.Println("Slice3:", slice3)

	slice4 := names[:]
	fmt.Println("Slice4:", slice4)

	var slice5 []string = names[1:4]
	fmt.Println("Slice5:", slice5)

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	daysSlice1 := days[5:] // Fri, Sat

	fmt.Println("Days Slice1:", daysSlice1)

	daysSlice1[0] = "New Fri"
	daysSlice1[1] = "New Sat"

	fmt.Println("Days Slice1:", daysSlice1)
	fmt.Println("Days:", days)

	daysSlice2 := append(daysSlice1, "Holiday")
	daysSlice2[0] = "Fri Again"

	fmt.Println("Days Slice1:", daysSlice1)
	fmt.Println("Days Slice2:", daysSlice2)
	fmt.Println("Days:", days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "A"
	newSlice[1] = "B"

	fmt.Println("New Slice:", newSlice)
	fmt.Println("Length:", len(newSlice))
	fmt.Println("Capacity:", cap(newSlice))

	newSlice2 := append(newSlice, "C")
	// newSlice2 := append(newSlice, "C", "D", "E", "F")

	fmt.Println("New Slice2:", newSlice2)
	fmt.Println("Length:", len(newSlice2))
	fmt.Println("Capacity:", cap(newSlice2))

	newSlice2[0] = "Changed A"

	fmt.Println("New Slice:", newSlice)
	fmt.Println("New Slice2:", newSlice2)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println("From Slice:", fromSlice)
	fmt.Println("To Slice:", toSlice)
}
