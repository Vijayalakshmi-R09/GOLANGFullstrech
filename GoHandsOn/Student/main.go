package main

import "fmt"

type Student struct {
	id      int
	name    string
	mark1   int
	mark2   int
	mark3   int
	mark4   int
	mark5   int
	total   int
	average float32
	grade   string
}

func insertData(s [1][2]Student, n int) [1][2]Student {
	for i := 0; i < n; i++ {
		for j := 0; j >= i; j++ {
			if j == n {
				break
			}
			fmt.Println("Enter Student ID:")
			fmt.Scanln(&s[i][j].id)
			fmt.Println("Enter Student Name:")
			fmt.Scanln(&s[i][j].name)
			fmt.Println("Enter Mark1:")
			fmt.Scanln(&s[i][j].mark1)
			fmt.Println("Enter Mark2:")
			fmt.Scanln(&s[i][j].mark2)
			fmt.Println("Enter Mark3:")
			fmt.Scanln(&s[i][j].mark3)
			fmt.Println("Enter Mark4:")
			fmt.Scanln(&s[i][j].mark4)
			fmt.Println("Enter Mark5:")
			fmt.Scanln(&s[i][j].mark5)
			s[i][j].total = s[i][j].mark1 + s[i][j].mark2 + s[i][j].mark3 + s[i][j].mark4 + s[i][j].mark5
			s[i][j].average = float32(s[i][j].total) / 5
		}
	}
	fmt.Println("Record Added Successfully")
	return s
}

func searchData(s [1][2]Student, sid int) {
	for i := 0; i < 2; i++ {
		for j := 0; j >= i; j++ {
			if j == 2 {
				break
			}
			if s[i][j].id == sid {
				fmt.Println("Record you asked for: \n", s[i][j])
			}
		}
	}
}

func main() {

	var s [1][2]Student

	flag := true
	for flag {
		fmt.Println("Select One option \n1.Add student 2.View all 3.Search 4.Exit")
		var option int
		fmt.Scanln(&option)

		switch option {

		case 1:
			fmt.Println("Enter the no of records to be added: ")
			var n int
			fmt.Scanln(&n)
			if n <= 2 {
				s = insertData(s, n)
			} else {
				fmt.Println("Maximum Size exceeded")
			}

		case 2:
			fmt.Println("Student Records:")
			fmt.Println(s)

		case 3:
			fmt.Println("Enter the student ID to be searched: ")
			var sid int
			fmt.Scanln(&sid)
			searchData(s, sid)

		case 4:
			flag = false
			break

		default:
			fmt.Println("Incorrect Selection Pls Select from given option")
		}
	}
}
