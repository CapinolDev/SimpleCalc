package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Select type:")
	fmt.Println("1 - Basic arithmetic")
	fmt.Println("2 - Geometry")
	scanner.Scan()
	mathType, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	if mathType == 1 {
		fmt.Println("1 - addition")
		fmt.Println("2 - subtraction")
		fmt.Println("3 - division")
		fmt.Println("4 - multiplication")
		scanner.Scan()
		input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if input == 1 {
			fmt.Println("First number:")
			scanner.Scan()
			number1, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			fmt.Println("Second number:")
			scanner.Scan()
			number2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			answer := number1 + number2
			fmt.Printf("The answer: %d", answer)

		}
		if input == 2 {
			fmt.Println("First number:")
			scanner.Scan()
			number1, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			fmt.Println("Second number:")
			scanner.Scan()
			number2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			answer := number1 - number2
			fmt.Printf("The answer: %d", answer)

		}
		if input == 3 {
			fmt.Println("First number:")
			scanner.Scan()
			number1, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			fmt.Println("Second number:")
			scanner.Scan()
			number2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			answer := number1 / number2
			fmt.Printf("The answer: %d", answer)

		}
		if input == 4 {
			fmt.Println("First number:")
			scanner.Scan()
			number1, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			fmt.Println("Second number:")
			scanner.Scan()
			number2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			answer := number1 * number2
			fmt.Printf("The answer: %d", answer)

		}

	}
	if mathType == 2 {
		fmt.Println("Select shape:")
		fmt.Println("1 - Square")
		fmt.Println("2 - Triangle")
		fmt.Println("3 - Rectangle")
		scanner.Scan()
		shape, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if shape == 1 {
			fmt.Println("Which operation?")
			fmt.Println("1 - Calculate circumference")
			fmt.Println("2 - Calculate area")
			scanner.Scan()
			operation, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			if operation == 1 {
				fmt.Println("What's the side lenght?")
				scanner.Scan()
				sideLenght, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				circumference := sideLenght * 4
				fmt.Printf("The circumference is: %d", circumference)
			}
			if operation == 2 {
				fmt.Println("What's the side lenght?")
				scanner.Scan()
				sideLenght, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				area := sideLenght * sideLenght
				fmt.Printf("The area is: %d", area)
			}
		}
		if shape == 2 {
			fmt.Println("Which operation?")
			fmt.Println("1 - Calculate circumference")
			fmt.Println("2 - Calculate area")
			fmt.Println("3 - Calculate height (two sides need to be of equal lenght)")
			scanner.Scan()
			operation, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			if operation == 1 {
				fmt.Println("What's the first side lenght?")
				scanner.Scan()
				firstSideLenght, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				fmt.Println("What's the second side lenght?")
				scanner.Scan()
				secondSideLenght, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				fmt.Println("What's the third side lenght?")
				scanner.Scan()
				thirdSideLenght, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				circumference := firstSideLenght + secondSideLenght + thirdSideLenght
				fmt.Printf("The circumference is: %v", circumference)

			}
			if operation == 2 {
				fmt.Println("How long is the height?")
				scanner.Scan()
				height, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				fmt.Println("How long is the base?")
				scanner.Scan()
				baseLenght, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				area := (height * baseLenght) / 2
				fmt.Printf("The area of the triangle is: %v", area)
			}
			if operation == 3 {
				fmt.Println("Input base lenght: ")
				scanner.Scan()
				baseLenght, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				fmt.Println("Input side lenght: ")
				scanner.Scan()
				sideLenght, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				var height float64 = math.Sqrt(float64(sideLenght*sideLenght - (baseLenght/2)*(baseLenght/2)))
				fmt.Printf("The height is %v", height)

			}
		}
		if shape == 3 {
			fmt.Println("Select operation: ")
			fmt.Println("1 - Circumference")
			fmt.Println("2 - Area")
			scanner.Scan()
			operation, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			if operation == 1 {
				fmt.Println("Input side a: ")
				scanner.Scan()
				sideA, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				fmt.Println("Input side b: ")
				scanner.Scan()
				sideB, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				circumference := sideA*2 + sideB*2
				fmt.Printf("The circumference of the rectangle is %v", circumference)
			}
			if operation == 2 {
				fmt.Println("Input side a: ")
				scanner.Scan()
				sideA, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				fmt.Println("Input side b: ")
				scanner.Scan()
				sideB, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				area := sideA * sideB
				fmt.Printf("The area is %v", area)

			}

		}
	}

}
