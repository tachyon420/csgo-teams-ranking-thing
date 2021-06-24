package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Team struct{
	name string
	rank int
	score float64
}

func ending(exStatus int) {
	fmt.Println("right so you just put a character instead of a number or vice versa... try again bb - \npress x and enter to close program, or anything other character and enter to restart")
	var x string
	fmt.Scan(&x)
	if x == "x" {
		os.Exit(exStatus)
	} else {
		main()
	}
}

func createTeamArray(filename string) []Team {
	// Open and read file
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		ending(1)
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
			ending(1)
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
	fmt.Println(array)
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
	}
	return -1
}

func scoreComparison(bo int, scoring string, winning float64, losing float64) (float64, float64, float64, float64, []string, []string) {
	var (
		winningCh float64
		losingCh float64
		scoreArray []string
		mapMultArray []string
	)
	switch bo {
	case 1:
		winning *= 0.9
		losing *= 0.9

		var game1 string
		fmt.Println("Enter the game 1")
		fmt.Scan(&game1)

		// split game1
		game1Split := strings.Split(game1, "-")
		game1Win, err := strconv.Atoi(game1Split[0])
		if err != nil {
			log.Fatal(err)
			ending(1)
		}
		game1Lost, err1 := strconv.Atoi(game1Split[1])
		if err1 != nil {
			log.Fatal(err1)
			ending(1)
		}

		// for the change log
		winningCh, losingCh = 0.9, 0.9
		scoreArray = append(scoreArray, game1)

		var (
			winningMultArray float64
			losingMultArray float64
		)

		diff := game1Win - game1Lost
		if diff <= 4 {
			winning *= 0.9
			losing *= 0.7
			winningMultArray = 0.9
			losingMultArray = 0.8
		} else if diff <= 9 {
			winning = winning
			losing = losing
			winningMultArray = 0.9
			losingMultArray = 0.8
		} else if diff <= 13 {
			winning *= 1.2
			losing *= 1.2
			winningMultArray = 0.9
			losingMultArray = 0.8
		} else if diff <= 16 {
			winning *= 1.5
			losing *= 0.7
			winningMultArray = 0.9
			losingMultArray = 0.8
		}

		mapMultArray = append(mapMultArray, fmt.Sprintf("%.2f / %.2f", winningMultArray, losingMultArray) )
	case 3:
		var game1 string
		var game2 string
		fmt.Println("Enter the game 1")
		fmt.Scan(&game1)
		fmt.Println("Enter the game 2")
		fmt.Scan(&game2)
		// split game1
		game1Split := strings.Split(game1, "-")
		game1Win, err := strconv.Atoi(game1Split[0])
		if err != nil {
			ending(1)
		}
		game1Lost, err1 := strconv.Atoi(game1Split[1])
		if err1 != nil {
			log.Fatal(err1)
			ending(1)
		}
		
		// split game 2
		game2Split := strings.Split(game2, "-")
		game2Win, err := strconv.Atoi(game2Split[0])
		if err != nil {
			log.Fatal(err)
			ending(1)
		}
		game2Lost, err1 := strconv.Atoi(game2Split[1])
		if err1 != nil {
			log.Fatal(err1)
			ending(1)
		}
		// set the multipliers deppending on the difference in games

		var winningA float64
		var losingA float64
		diffArray :=  []int{game1Win, game1Lost, game2Win, game2Lost}

		switch scoring {
		case "2-0":
			winning *= 1.2
			losing *= 1.2

			// for the change log
			winningCh, losingCh = 1.2, 1.2
			scoreArray = append(scoreArray, game1, game2)

			var (
				winningMultArray float64
				losingMultArray float64
			)

			for i := 0; i < 4; i += 2 {
				diff := diffArray[i] - diffArray[i+1]
				if diff <= 4 {
					winningA += 0.9
					losingA += 0.7
					winningMultArray = 0.9
					losingMultArray = 0.7
				} else if diff <= 9 {
					winningMultArray = 1.0
					losingMultArray = 1.0
				} else if diff <= 13 {
					winningA += 1.2
					losingA += 1.2
					winningMultArray = 1.2
					losingMultArray = 1.2
				} else if diff <= 16 {
					winningA += 1.5
					losingA += 1.3
					winningMultArray = 1.5
					losingMultArray = 1.3
				}
				mapMultArray = append(mapMultArray, fmt.Sprintf("%.2f / %.2f", winningMultArray, losingMultArray) )
				
			}
			winning *= (winningA/2)
			losing *= (losingA/2)
		case "2-1":
			// declare the game3 and assign it to user input
			var game3 string
			fmt.Println("Enter the game that they lost")
			fmt.Scan(&game3)

			// split the game and assign it to game3split
			game3Split := strings.Split(game3, "-")
			game3Win, err := strconv.Atoi(game3Split[0])
			if err != nil {
				ending(1)
			}

			game3Lost, err := strconv.Atoi(game3Split[1])
			if err != nil {
				ending(1)
			}

			diffArray = append(diffArray, game3Win, game3Lost)

			// for the change log
			winningCh, losingCh = 1.00, 1.00
			scoreArray = append(scoreArray, game3)
			var winningMultArray float64
			var losingMultArray float64

			for i := 0; i < 4; i += 2 {
				diff := diffArray[i] - diffArray[i+1]
				if diff <= 4 {
					winningA += 0.9
					losingA += 0.8
					winningMultArray = 0.9
					losingMultArray = 0.8
				} else if diff <= 9 {
					winningMultArray = 1.0
					losingMultArray = 1.0
				} else if diff <= 13 {
					winningA += 1.15
					losingA += 1.15
					winningMultArray = 1.15
					losingMultArray = 1.15
				} else if diff <= 16 {
					winningA += 1.3
					losingA += 1.3
					winningMultArray = 1.3
					losingMultArray = 1.3
				}
				mapMultArray = append(mapMultArray, fmt.Sprintf("%.2f / %.2f", winningMultArray, losingMultArray) )
			}
			diff := game3Lost - game3Win

			var (
				winningUpd float64
				losingUpd float64
			)
			if diff <= 4 {
				winningA += 1.05
				winningUpd = 1.05
				losingA += 1.10
				losingUpd = 1.10
			} else if diff <= 9 {
				fmt.Println("cool")
				winningUpd = 1.0
				losingUpd = 1.0
			} else if diff <= 13 {
				winningA += 0.95
				winningUpd = 0.95
				losingA += 1.3
				losingUpd = 1.3
			} else if diff <= 16 {
				winningA += 1.15
				winningUpd = 1.15
				losingA += 1.5
				losingUpd = 1.5
			}
			mapMultArray = append(mapMultArray, fmt.Sprintf("%.2f / %.2f", winningUpd, losingUpd) )
			winning *= (winningA/3)
			losing *= (losingA/3)			
		}
	case 5:
		//declare initial won games
		var game1 string
		fmt.Println("Enter game score 1")
		fmt.Scan(&game1)
		
		var game2 string
		fmt.Println("Enter game score 2")
		fmt.Scan(&game2)

		var game3 string
		fmt.Println("Enter game score 3")
		fmt.Scan(&game3)

		//splitting the game 1
		game1Split := strings.Split(game1, "-")
		game1Win, err := strconv.Atoi(game1Split[0])
		if err != nil {
			ending(1)
		}
		game1Lose, err := strconv.Atoi(game1Split[1])
		if err != nil {
			ending(1)
		}

		game2Split := strings.Split(game2, "-")
		game2Win, err := strconv.Atoi(game2Split[0])
		if err != nil {
			ending(1)
		}
		game2Lose, err := strconv.Atoi(game2Split[1])
		if err != nil {
			ending(1)
		}

		game3Split := strings.Split(game3, "-")
		game3Win, err := strconv.Atoi(game3Split[0])
		if err != nil {
			ending(1)
		}
		game3Lose, err := strconv.Atoi(game3Split[1])
		if err != nil {
			ending(1)
		}

		//declare the multipliers
		var winningA float64
		var losingA float64

		switch scoring {
		case "3-0":
			winningA += 1.3
			losingA += 0.7
		case "3-1":
			winningA += 1.15
			losingA += 0.85
		}
		winningCh, losingCh = winningA, losingA
		scoreArray = append(scoreArray, game1, game2, game3)

		diffArray := []int{game1Win, game1Lose, game2Win, game2Lose, game3Win, game3Lose}

		//for loop to calculate the multipliers in accordance to the diffArray
		for i := 0; i < 6; i += 2 {
			diff := diffArray[0] - diffArray[i+1]
			if diff <= 4 {
				winningA += 0.8
				losingA += 0.85
			} else if diff <= 9 {
				winningA += 1.10
				losingA += 1
			} else if diff <= 13 {
				winningA += 1.35
				losingA += 1.20
			} else if diff <= 16 {
				winningA += 1.50
				losingA += 1.40
			}
			mapMultArray = append(mapMultArray, fmt.Sprintf("%.2f / %.2f", winningA, losingA) )
		}

		winning *= (winningA/3)
		losing *= (losingA/3)
		
		//if loops to check the other scoring standard, for games lost
		if scoring == "3-1" || scoring == "3-2" {
			var game4 string
			fmt.Println("Enter the game they lost")
			fmt.Scan(&game4)

			//splitting game 4
			game4Split := strings.Split(game4, "-")
			game4Lose, err := strconv.Atoi(game4Split[0])
			if err != nil {
				ending(1)
			}
			game4Win, err := strconv.Atoi(game4Split[1])
			if err != nil {
				ending(1)
			}

			diffArrayLost := append(diffArray, game4Win, game4Lose)
			// to make sure it divides correctly
			mult := 1.0

			// for the change log
			scoreArray = append(scoreArray, game4)

			if scoring == "3-2" {
				mult = 2.0

				var game5 string
				fmt.Println("enter the last game they lost")
				fmt.Scan(&game5)

				//splitting game5
				game5Split := strings.Split(game5, "-")
				game5Lose, err := strconv.Atoi(game5Split[0])
				if err != nil {
					ending(1)
				}	
				game5Win, err := strconv.Atoi(game5Split[1])
				if err != nil {
					ending(1)
				}

				//add the values to the array
				diffArrayLost = append(diffArrayLost, game5Win, game5Lose) 
				winningCh, losingCh = 1, 1
				scoreArray = append(scoreArray, game5)
			}
			var winningB float64
			var losingB float64

			for i := 0; i < len(diffArrayLost); i+= 2 {
				diff := diffArrayLost[i] - diffArrayLost[i+1]
				if diff <= 4 {
					winningB += 1.10
					losingB += 1.10
				} else if diff <= 9 {
					winningB += 1
					losingB += 0.9
				} else if diff <= 13 {
					winningB += 0.7
					losingB += 0.8
				} else if diff <= 16 {
					winningB += 0.50
					losingB += 0.65
				}
				mapMultArray = append(mapMultArray, fmt.Sprintf("%.2f / %.2f", winningB, losingB) )
			}

			
			winning *= (winningB/mult)
			losing *= (losingB/mult)
		}
	default:
		ending(3)	
	}
	return winning, losing, winningCh, losingCh, scoreArray, mapMultArray
}

func main() {
	teamsArray := createTeamArray("rankings.txt")
	
	var team1 string
	fmt.Println("Enter the name of the winning team")
	fmt.Scan(&team1)

	var team2 string
	fmt.Println("Enter the name of the losing team")
	fmt.Scan(&team2)

	// check if values exist, and get values for the array
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

	// insure that the correct values are entered
	bestOfCorrect := false
	for bestOfCorrect == false {
		fmt.Println("Best of how many games?")
		fmt.Scan(&bo)
		if bo == 1 || bo == 3 || bo == 5 {
			bestOfCorrect = true
			break
		}
	}
	
	// get scoring for bo
	fmt.Println("Overall score?")
	fmt.Scan(&scoring)

	// SCORE DIFFERENCE ASSIGNMENT DISPARITY ON CLOSENESS OF RANKS
	if (team1Struct.rank <= 15 && team2Struct.rank >= 30) {
		winning = 5
		losing = 2
	} else if (team1Struct.rank - team2Struct.rank < 0) {
		winning = 5
		losing = 5
	} else if team1Struct.rank - team2Struct.rank <= 5{
		winning = 10
		losing = 10
	} else if team1Struct.rank == team2Struct.rank{
		winning = 7
		losing = 7
	} else if team1Struct.rank - team2Struct.rank <= 10 {
		winning = 15
		losing = 15
	} else if team1Struct.rank - team2Struct.rank <= 20 {
		winning = 20
		losing = 20
	} else if team1Struct.rank - team2Struct.rank <= 30 {
		winning = 25
		losing = 25
	} else if team1Struct.rank - team2Struct.rank > 30 {
		winning = 30
		losing = 30
	} 
	
	// for the change log
	winningUpdate := winning
	// losingUpdate := losing
	var (
		winningCh float64
		losingCh float64
		scoreArray []string
		mapMultArray []string
	)

	// SCORE DIFFERENCE ADJUSTMENT DISPARITY ON BEST OF (X) AND SCORING
	winning, losing, winningCh, losingCh, scoreArray, mapMultArray = scoreComparison(bo, scoring, winning, losing)
	
	// ADJUST NEW SCORES
	team1Struct.score += winning
	team2Struct.score -= losing
	
	// DELETE OLD RANKING FILE AND CREATE NEW ONE
	err := os.Remove("rankings.txt")
	if err != nil {
		log.Fatal(err)
		ending(1)
	}

	newFile, err := os.Create("rankings.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(newFile)
		ending(1)
	}
	
	file, err := os.OpenFile("rankings.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
		ending(1)
	}

	// SET THE NEW STRUCTS WITH ADJUSTED SCORES
	teamsArray[arrayValue1] = team1Struct
	teamsArray[arrayValue2] = team2Struct

	for i := 0; i < len(teamsArray); i++ {
		total := teamsArray[i].name + " " + strconv.Itoa(teamsArray[i].rank) + " " + fmt.Sprintf("%f", teamsArray[i].score) + "\n"
		l, err := file.WriteString(total)
		
		if err != nil {
			log.Fatal(err, l)
			ending(1)
		}
		// fmt.Printf("Successfully updated %v bytes", l)
	}

	// SET UPDATES INTO A FILE
	// create values for the updates

	var rankDiff int = team1Struct.rank - team2Struct.rank

	// print/add the new values to the file
	changelog := fmt.Sprintf("\n[%v] - [%v] beat [%v] in a bo[%v] with a ranking difference of %v resulting in [%v] points in a [%v] giving a map multiplier of [%v]/[%v]. The scores were %v giving an overall score multiplier per map of %v. Overall point change of %.2f.",
	time.Now().Format("02-01-2006"), team1Struct.name, team2Struct.name, bo, rankDiff, winningUpdate, scoring, winningCh, losingCh, scoreArray, mapMultArray, winning)
	// winningUpdate, scoring, winningCh, losingCh, scoreArray, mapMultArray, winning)

	newFile, errors := os.OpenFile("record.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(errors)
	}

	l, err := newFile.WriteString(changelog)
	if err != nil {
		fmt.Printf("error printing: %v", err)
	}

	for i := 0; i < l; i++ {
		i++
	}



	fmt.Println("Enter to close program")
	var i string
	fmt.Scanln(&i)
	fmt.Println(i)
}