package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Team struct{
	name string
	rank int
	score float64
}

func createTeamArray(filename string) []Team {
	// Open and read file
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to string and split each line into a splice
	teams := string(contents)
	teamsB := strings.Split(teams, "\n")
	
	// Create the 2D splice
	var teamsA [][]string

	// For loop for assigning values to the 2D splice
	for i := 0; i < len(teamsB); i++ {
		teamsBindex := strings.Split(teamsB[i], " ")
		teamsA = append(teamsA, teamsBindex)
	}

	// Create splice for storing the teams
	var teamsArray []Team

	for i := 0; i < len(teamsA)-1; i++ {
		name := teamsA[i][0]

		ranking, err := strconv.Atoi(teamsA[i][1])
		if err != nil {
			log.Fatal(err)
		}	

		score, err := strconv.ParseFloat(teamsA[i][2], 32)
		if err != nil {
			log.Fatal(err)
		}

		team := Team{name, ranking, score}
		teamsArray = append(teamsArray, team)
	}
	return teamsArray
}

func contains(array []Team, str string) Team {
	for i := 0; i < len(array); i++ {
		if array[i].name == str {
			return array[i]
		}
	}
	fmt.Println("error - name not available. Maybe you forgot an underscore? CTRL+C to close the program")
	var a string
	fmt.Scan(&a)
	os.Exit(1)
	return array[0]
}

func containsForCounter(array []Team, str string) int {
	for i := 0; i < len(array); i++ {
		if array[i].name == str {
			return i
		}
		fmt.Println("error - name not available. Maybe you forgot an underscore?")
	}
	return -1
}

func main() {
	teamsArray := createTeamArray("rankings.txt")
	
	var team1 string
	fmt.Println("Enter the name of the winning team")
	fmt.Scan(&team1)

	var team2 string
	fmt.Println("Enter the name of the losing team")
	fmt.Scan(&team2)

	team1Struct := contains(teamsArray, team1)
	arrayValue1 := containsForCounter(teamsArray, team1)
	team2Struct := contains(teamsArray, team2)
	arrayValue2 := containsForCounter(teamsArray, team2)
	var winning float64
	var losing float64

	var (
		bo int
		scoring	string
	)

	fmt.Println("Best of how many games?")
	fmt.Scan(&bo)

	fmt.Println("Overall score?")
	fmt.Scan(&scoring)

	// SCORE DIFFERENCE ASSIGNMENT DISPARITY ON CLOSENESS OF RANKS
	if (team1Struct.rank <= 15 && team2Struct.rank >= 30) {
		winning = 5
		fmt.Println(team1Struct.score)
		losing = 2
		fmt.Println(team2Struct.score)
	} else if (team1Struct.rank - team2Struct.rank < 0) {
		winning = 5
		fmt.Println(team1Struct.score)
		losing = 5
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank <= 5{
		winning = 10
		fmt.Println(team1Struct.score)
		losing = 10
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank == team2Struct.rank{
		winning = 7
		fmt.Println(team1Struct.score)
		losing = 7
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank <= 10 {
		winning = 15
		fmt.Println(team1Struct.score)
		losing = 15
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank <= 20 {
		winning = 20
		fmt.Println(team1Struct.score)
		losing = 20
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank <= 30 {
		winning = 25
		fmt.Println(team1Struct.score)
		losing = 25
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank > 30 {
		winning = 30
		fmt.Println(team1Struct.score)
		losing = 30
		fmt.Println(team2Struct.score)
	} 

	// SCORE DIFFERENCE ADJUSTMENT DISPARITY ON BEST OF (X) AND SCORING
	switch bo {
	case 1:
		winning *= 0.9
		losing *= 0.9
	case 3:
		if scoring == "2-0" {
			winning *= 1.2
			losing *= 1.2
		}
	case 5:
		if scoring == "3-0" {
			winning *= 1.3
			losing *= 1.3
		} else if scoring == "3-1" {
			winning *= 1.15
			losing *= 1.15
		}
	}
	
	// ADJUST NEW SCORES
	team1Struct.score += winning
	team2Struct.score -= losing
	
	// DELETE OLD FILE AND CREATE NEW ONE
	err := os.Remove("rankings.txt")
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create("rankings.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(newFile)
	}
	
	file, err := os.OpenFile("rankings.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}

	// SET THE NEW STRUCTS WITH ADJUSTED SCORES
	teamsArray[arrayValue1] = team1Struct
	teamsArray[arrayValue2] = team2Struct

	for i := 0; i < len(teamsArray); i++ {
		total := teamsArray[i].name + " " + strconv.Itoa(teamsArray[i].rank) + " " + fmt.Sprintf("%f", teamsArray[i].score) + "\n"
		l, err := file.WriteString(total)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully updated %v bytes", l)
	}
	fmt.Println("Enter to close program")
	var i string
	fmt.Scanln(&i)
	fmt.Println(i)
}