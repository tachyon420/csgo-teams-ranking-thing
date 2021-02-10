package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	file *os.File
	err  error
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	i := false
	if i == true {
		fmt.Println(info)
	}
	return true
}

func main() {
	if fileExists("rankings.txt") == false {
		file, err := os.Create("rankings.txt")
		if err != nil {
			log.Fatal(err)
		}

		i := false
		if i == true {
			fmt.Println(file)
		}
	}
	file.Close()

	file, err := os.OpenFile("rankings.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("error opening file,", err)
	}

	var active string
	fmt.Println("yes to keep adding more teams, no to cancel")
	fmt.Scanln(&active)
	for active != "no" {
		var teamName string
		fmt.Println("Enter the team name")
		fmt.Scanln(&teamName)

		var ranking int = 0

		var score int
		fmt.Println("Enter current score of the team")
		fmt.Scanln(&score)

		total := (teamName + " " + strconv.Itoa(ranking) + " " + strconv.Itoa(score) + "\n")

		l, err := file.WriteString(total)
		if err != nil {
			fmt.Println("error printing:", err)
		}
		fmt.Printf("Successfully printed %v bytes\n", l)

		fmt.Println("yes to keep adding more teams, no to cancel")
		fmt.Scanln(&active)
	}
	file.Close()
}
