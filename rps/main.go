package main

import (
		"flag"
		"fmt"
		"math/rand"
		"strings"
		"time"
)

func main() {
		rounds := flag.Int("rounds", 1, "Specify no of rounds")
		flag.Parse()
		startGame(*rounds)
}

func startGame(rounds int) {
		var userScore int
		var cpuScore int
		for i := 0; i < rounds; i++ {
				fmt.Println("Enter your move")
				userMove := getMove()
				cpuMove := cpuMove()
				fmt.Println("Cpu chose", cpuMove)
				winner := winnerIs(userMove, cpuMove)
				if winner == "cpu" {
						cpuScore++
				} else if winner == "user" {
						userScore++
				} else {
						cpuScore++
						userScore++
				}
		}
		fmt.Println("Player score is", userScore)
		fmt.Println("CPU score is", cpuScore)
		finalWinner(userScore, cpuScore)
}

func getMove() string {
		var move string
		_, err := fmt.Scanf("%s", &move)
		if err != nil {
				panic(err)
		}
		return move
}
func cpuMove() string {
		moves := []string{"rock", "paper", "scissors"}
		random := randomInt(len(moves))
		return strings.ToLower(moves[random])
}

func randomInt(n int) int {
		rand.Seed(time.Now().UTC().UnixNano())
		return rand.Intn(n)
}

func ruleEngine() map[string]string {
		rule := make(map[string]string)
		rule["rock"] = "scissors"
		rule["paper"] = "rock"
		rule["scissors"] = "paper"
		return rule
}

func winnerIs(userMove string, cpuMove string) string {
		winner := ruleEngine()
		if winner[cpuMove] == userMove {
				fmt.Println("CPU wins this round")
				fmt.Println("-----")
				return "cpu"
		} else if winner[userMove] == cpuMove {
				fmt.Println("Player wins this round")
				fmt.Println("-----")
				return "user"
		} else {
				fmt.Println("This match is tied")
				fmt.Println("-----")
				return "both"
		}
}

func finalWinner(userScore int, cpuScore int) {
		if userScore > cpuScore {
				fmt.Println("Player wins the battle")
		} else if cpuScore > userScore {
				fmt.Println("CPU wins the battle")
		} else {
				fmt.Println("The game is tied")
		}
}