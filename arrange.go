package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Team struct {
	name    string
	ranking int
	score   int
}

func main() {
	contents, err := ioutil.ReadFile("rankings.txt")
	if err != nil {
		fmt.Println("have you run teamAdd.exe first? try that, as rankings.txt may not have been created yet.\n and so could lead to more errors here")
	}
	teams := string(contents)
	teamsB := strings.Split(teams, "\n")
	var teamsA [][]string
	for i := 0; i < len(teamsB); i++{
		teamsBindex := strings.Split(teamsB[i], " ")
		teamsA = append(teamsA, teamsBindex)
		fmt.Println(teamsA[i])
	}
	
	fmt.Println(teamsA)
	var teamsArray []Team

	for i := 0; i < len(teamsA)-1; i++ {
		ranking, err1 := strconv.Atoi(teamsA[i][1])
		fmt.Println(ranking)
		if err1 != nil {
			log.Fatal(err1)
		}
		score, err2 := strconv.Atoi(teamsA[i][2])
		fmt.Println(score)
		if err2 != nil {
			log.Fatal(err2)
		}
		team := Team{teamsA[i][0], ranking, score}
		teamsArray = append(teamsArray, team)
	}

	sort.Slice(teamsArray, func(i, j int) bool {
		return teamsArray[i].score > teamsArray[j].score
	})

	fmt.Println(teamsArray[0].name)

	for i := 0; i < len(teamsArray); i++ {
		if (i > 0) {
			fmt.Println("hi")
			if (teamsArray[i-1].score == teamsArray[i].score) {
				fmt.Println("heya")
				teamsArray[i].ranking = teamsArray[i-1].ranking
			}
			
		} else {
			teamsArray[i].ranking = teamsArray[i-1].ranking + 1
		}
	}

	fmt.Println("helloyeetus")
	err = os.Remove("rankings.txt")
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
	
	for i := 0; i < len(teamsArray); i++ {
		total := teamsArray[i].name + " " + strconv.Itoa(teamsArray[i].ranking) + " " + strconv.Itoa(teamsArray[i].score) + "\n"
		l, err := file.WriteString(total)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfuly printed %v bytes\n", l)
	}

}
