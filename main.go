package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	var name string = "John doe"
	var age int = 25
	city := "New York"
	year := 2026
	description := `This is a multi-line
string that spans
several lines.`

	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("City : ", city)
	fmt.Println("Year : ", year)
	fmt.Println("Description : ", description)

	const pi = 3.14

	fmt.Println("Value of pi : ", pi)

	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = -100

	fmt.Println("\nSigned Integers:")
	fmt.Printf("int8 : %v\n", i8)
	fmt.Printf("int16 : %v\n", i16)
	fmt.Printf("int32 : %v\n", i32)
	fmt.Printf("int64 : %v\n", i64)
	fmt.Printf("int : %v\n", i)

	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 100

	fmt.Println("\nUnsigned Integers:")
	fmt.Printf("uint8 : %v\n", u8)
	fmt.Printf("uint16 : %v\n", u16)
	fmt.Printf("uint32 : %v\n", u32)
	fmt.Printf("uint64 : %v\n", u64)
	fmt.Printf("uint : %v\n", u)

	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	fmt.Println("\nFloating-Point Numbers:")
	fmt.Printf("float32 : %v\n", f32)
	fmt.Printf("float64 : %v\n", f64)

	var isTrue bool = true
	var isFalse bool = false

	fmt.Println("\nBoolean Values:")
	fmt.Printf("isTrue : %v\n", isTrue)
	fmt.Printf("isFalse : %v\n", isFalse)

	text := "Go is awesome!"

	fmt.Println("\nOperations String:")
	fmt.Println("Lowercase: ", strings.ToLower(text))
	fmt.Println("Uppercase: ", strings.ToUpper(text))
	fmt.Println("Stars with 'Go': ", strings.HasPrefix(text, "Go"))
	fmt.Println("Constains 'awesome': ", strings.Contains(text, "awesome"))

	parts := strings.Split("apple,banana,cherry", ",")
	fmt.Println("Split: ", parts)

	newText := strings.ReplaceAll(text, "awesome", "fantastic")
	fmt.Println("Replace: ", newText)

	fmt.Println("\nConvertion:")

	var a int = 42
	var b float64 = float64(a)

	fmt.Println("Value of a (int): ", a)
	fmt.Println("Converted b (float64): ", b)

	var score int = 85
	var scoreStr string = strconv.Itoa(score)

	fmt.Println("Value of score (int): ", score)
	fmt.Println("Converted scoreStr (string): ", scoreStr)

	var str string = "123"
	number, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Value of str (string): ", str)
		fmt.Println("Converted number (int): ", number)
	}

	truth := true
	truthStr := strconv.FormatBool(truth)

	fmt.Println("Value of truth (boolean to string): ", truthStr)

	truthBool, _ := strconv.ParseBool(truthStr)
	fmt.Println("Value of truthBool (string to boolean): ", truthBool)

	fmt.Println("\nArray:")

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array numbers: ", numbers)
	fmt.Println("Length of array: ", len(numbers))
	fmt.Println("First element: ", numbers[0])

	numbers[2] = 10
	fmt.Println("Modified array: ", numbers)

	for i, v := range numbers {
		fmt.Printf("Index %d: Value %d\n", i, v)
	}

	arr1 := [3]string{"A", "B", "C"}
	arr2 := [3]string{"A", "B", "D"}

	fmt.Println("Are arr1 and arr2 equal? ", arr1 == arr2)

	fmt.Println("\nSlice:")

	arr3 := [5]int{10, 20, 30, 40, 50}
	slice1 := arr3[1:4]
	slice2 := arr3[:3]
	slice3 := arr3[2:]
	slice4 := arr3[:]

	fmt.Println("Slice slice1: ", slice1)
	fmt.Println("Slice slice2: ", slice2)
	fmt.Println("Slice slice3: ", slice3)
	fmt.Println("Slice slice4: ", slice4)

	slice5 := make([]int, 3, 5)
	fmt.Println("Slice slice5 (made with make): ", slice5)
	fmt.Println("Length of slice5: ", len(slice5))
	fmt.Println("Capacity of slice5: ", cap(slice5))

	slice5 = append(slice5, 1, 2, 3)
	fmt.Println("Appended slice5: ", slice5)

	slice6 := make([]int, 3)
	slice7 := copy(slice6, slice5)

	fmt.Println("Slice slice6 (made with make): ", slice6)
	fmt.Println("Slice slice7 (copied from slice5): ", slice7)

	slice8 := numbers[1:4]
	fmt.Println("Slice slice8 from numbers: ", slice8)

	fmt.Println("\nMap:")

	student := map[string]string{
		"name":  "Alice",
		"age":   "23",
		"class": "Biology",
	}

	fmt.Println("Map student: ", student)
	fmt.Println("Student Name: ", student["name"])
	fmt.Println("Student Age: ", student["age"])
	fmt.Println("Student Class: ", student["class"])

	student["age"] = "24"
	fmt.Println("Updated Student Age: ", student["age"])

	student["grade"] = "A"
	fmt.Println("Added Student Grade: ", student["grade"])

	fmt.Println("Iterating over map:")
	for key, value := range student {
		fmt.Printf("%s : %s\n", key, value)
	}

	student2 := map[string]int{
		"math":    90,
		"science": 85,
		"english": 88,
	}

	fmt.Println("Map student2: ", student2)

	delete(student2, "science")
	fmt.Println("Deleted science from student2: ", student2)

	student3 := make(map[string]string)
	student3["name"] = "Bob"
	student3["age"] = "22"

	fmt.Println("Map student3 (made with make): ", student3)

	fmt.Println("\nMath Operation:")

	x := 10
	y := 3

	fmt.Printf("%d + %d = %d\n", x, y, x+y)
	fmt.Printf("%d - %d = %d\n", x, y, x-y)
	fmt.Printf("%d * %d = %d\n", x, y, x*y)
	fmt.Printf("%d / %d = %d\n", x, y, x/y)
	fmt.Printf("%d %% %d = %d\n", x, y, x%y)

	fmt.Println("\nMath Operation (Assignment):")

	x += 5
	fmt.Println("x += 5 : ", x)
	x -= 3
	fmt.Println("x -= 3 : ", x)
	x *= 2
	fmt.Println("x *= 2 : ", x)
	x /= 4
	fmt.Println("x /= 4 : ", x)
	x %= 3
	fmt.Println("x %= 3 : ", x)

	fmt.Println("\nComparison Operators:")

	fmt.Printf("%d == %d : %v\n", x, y, x == y)
	fmt.Printf("%d != %d : %v\n", x, y, x != y)
	fmt.Printf("%d > %d : %v\n", x, y, x > y)
	fmt.Printf("%d < %d : %v\n", x, y, x < y)
	fmt.Printf("%d >= %d : %v\n", x, y, x >= y)
	fmt.Printf("%d <= %d : %v\n", x, y, x <= y)

	fmt.Println("\nLogical Operators:")

	aBool := true
	bBool := false

	fmt.Printf("%v && %v : %v\n", aBool, bBool, aBool && bBool)
	fmt.Printf("%v || %v : %v\n", aBool, bBool, aBool || bBool)
	fmt.Printf("!%v : %v\n", aBool, !aBool)

	fmt.Println("\nConditionals / Branching:")

	num := 7
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else if num == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Number is greater than Three")
	}

	if rem := num % 3; rem == 0 {
		fmt.Printf("%d is divisible by 3\n", num)
	}

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Number is greater than Three")
	}

	fmt.Println("\nLoops:")

	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration: ", i)
	}

	for i, v := range []string{"A", "B", "C"} {
		fmt.Printf("Index %d: Value %s\n", i, v)
	}
}
