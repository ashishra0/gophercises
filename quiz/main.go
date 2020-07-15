package main

import (
		"encoding/csv"
		"flag"
		"fmt"
		"os"
		"time"
)

func main() {
	//	Flags
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timelimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit("Failed to open the csv file")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("There was a problem parsing CSV file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	correct := 0
	for i, p := range problems {
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			answerCh := make(chan string)
			go func() {
					var answer string
					fmt.Scanf("%s\n", &answer)
					answerCh <- answer
			}()
			select {
			case <- timer.C:
					fmt.Println("Oops! You ran out of time")
					fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
					return
					case answer := <- answerCh:
							if answer == p.a {
									correct++
							}
			}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
		ret := make([]problem, len(lines))
		for i, line := range lines {
				ret[i] = problem{
						q: line[0],
						a: line[1],
				}
		}
		return ret
}

type problem struct {
		q string
		a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}