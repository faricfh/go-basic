package main

import "fmt"

type validationError struct {
	message string
}

func (v *validationError) Error() string {
	return v.message
}

type notFoundError struct {
	message string
}

func (n *notFoundError) Error() string {
	return n.message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID is required"}
	}

	if id != "123" {
		return &notFoundError{"Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)

	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error:", validationErr.message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error:", notFoundErr.message)
		} else {
			fmt.Println("Unknown error:", err)
		}

		// switch finalError := err.(type) {
		// case *validationError:
		// 	fmt.Println("Validation error:", finalError.message)
		// case *notFoundError:
		// 	fmt.Println("Not found error:", finalError.message)
		// default:
		// 	fmt.Println("Unknown error:", err)
		// }

	} else {
		fmt.Println("Data saved successfully")
	}
}
