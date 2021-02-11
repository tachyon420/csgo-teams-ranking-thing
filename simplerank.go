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
	score int
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

		score, err := strconv.Atoi(teamsA[i][2])
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
		fmt.Println("error - name not available. Maybe you forgot an underscore?")
	}
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

	if (team1Struct.rank <= 15 && team2Struct.rank >= 30) {
		team1Struct.score += 5
		fmt.Println(team1Struct.score)
		team2Struct.score -= 2
		fmt.Println(team2Struct.score)
	} else if (team1Struct.rank - team2Struct.rank < 0) {
		team1Struct.score += 5
		fmt.Println(team1Struct.score)
		team2Struct.score -= 5
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank <= 5{
		team1Struct.score += 10
		fmt.Println(team1Struct.score)
		team2Struct.score -= 10
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank == team2Struct.rank{
		team1Struct.score += 7
		fmt.Println(team1Struct.score)
		team2Struct.score -= 7
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank <= 10 {
		team1Struct.score += 15
		fmt.Println(team1Struct.score)
		team2Struct.score -= 15
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank <= 20 {
		team1Struct.score += 20
		fmt.Println(team1Struct.score)
		team2Struct.score -= 20
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank <= 30 {
		team1Struct.score += 25
		fmt.Println(team1Struct.score)
		team2Struct.score -= 25
		fmt.Println(team2Struct.score)
	} else if team1Struct.rank - team2Struct.rank > 30 {
		team1Struct.score += 30
		fmt.Println(team1Struct.score)
		team2Struct.score -= 30
		fmt.Println(team2Struct.score)
	} 
	
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
		total := teamsArray[i].name + " " + strconv.Itoa(teamsArray[i].rank) + " " + strconv.Itoa(teamsArray[i].score) + "\n"
		l, err := file.WriteString(total)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully updated %v bytes", l)
	}
	fmt.Scan()
}