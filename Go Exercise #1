/*1. Define a “student” struct that contains: name, major, age
2. Read user input from the command line to construct an array of three students
3. Output the array that is sorted in an increasing order of ages*/

/*Sample command line input:
	Bob Psychology 55
	Mary Computer_Science 23
	Sally Sociology 24*/

// What is the difference between an array and a slice in Go?
/*An array is a static partition of memory. A slice is a reference to a specific section of an array.
Thus, a slice can dynamically change its shape by redefining the cells of the array it is referring to. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Default length of slice is 3, per exercise instructions.
	students := make([]Student, 3)
	i := 0

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input student information with name, major, and age. Type 'exit' when finished.")

	for scanner.Scan() {
		// Text() reads lines of input at a time.
		line := scanner.Text()

		if line == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// Convert line of user input into a student type by splitting it into words.
		studentInfo := strings.Fields(line)

		fmt.Println(len(studentInfo))

		if len(studentInfo) != 3 {
			fmt.Println("Please provide 3 arguments. Add an underscore to names or subjects with multiple words.")
		}

		students[i] = Student{studentInfo[0], studentInfo[1], studentInfo[2]}
		i++
	}

	// Sort array of students by age by passing in a custom less(i, j int) bool function for the student array.
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].age < students[j].age
	})

	fmt.Println(students)
}

type Student struct {
	name string
	major string
	age string
}
