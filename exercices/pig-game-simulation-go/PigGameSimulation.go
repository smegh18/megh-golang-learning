package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const WinningScore = 100

// Player represents a player in the game
// Name: The name of the player
// Strategy: The player's strategy for holding after accumulating a certain score
// TotalScore: The player's total accumulated score
type Player struct {
	Name       string
	Strategy   int
	TotalScore int
}

// RollDice simulates a roll of a 6-sided die and returns the result
func RollDice() int {
	return rand.Intn(6) + 1
}

// PlayTurn simulates a single turn for a player
// The player continues rolling until they roll a 1 or accumulate a turn total that meets their strategy
func PlayTurn(player *Player) int {
	turnTotal := 0
	for {
		dice := RollDice()
		if dice == 1 {
			fmt.Printf("%s rolled a 1! Turn over with no score.\n", player.Name)
			return 0
		}
		turnTotal += dice
		fmt.Printf("%s rolled a %d, turn total: %d\n", player.Name, dice, turnTotal)
		if turnTotal >= player.Strategy {
			fmt.Printf("%s holds with a turn total of %d\n", player.Name, turnTotal)
			return turnTotal
		}
	}
}

// PlayGame simulates a single game between two players
// The game alternates turns between Player 1 and Player 2 until one reaches the winning score
func PlayGame(player1, player2 *Player) string {
	for {
		player1.TotalScore += PlayTurn(player1)
		fmt.Printf("%s's total score: %d\n\n", player1.Name, player1.TotalScore)
		if player1.TotalScore >= WinningScore {
			return player1.Name
		}

		player2.TotalScore += PlayTurn(player2)
		fmt.Printf("%s's total score: %d\n\n", player2.Name, player2.TotalScore)
		if player2.TotalScore >= WinningScore {
			return player2.Name
		}
	}
}

// SimulateGames runs multiple games between two players and tracks the results
// strategy1: The holding strategy for Player 1
// strategy2: The holding strategy for Player 2
// numGames: The number of games to simulate
func SimulateGames(strategy1, strategy2, numGames int) {
	wins1, wins2 := 0, 0
	for i := 0; i < numGames; i++ {
		player1 := &Player{Name: "Player 1", Strategy: strategy1}
		player2 := &Player{Name: "Player 2", Strategy: strategy2}
		fmt.Printf("Game %d:\n", i+1)
		winner := PlayGame(player1, player2)
		if winner == player1.Name {
			wins1++
		} else {
			wins2++
		}
	}
	fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n", strategy1, strategy2, wins1, numGames, float64(wins1)/float64(numGames)*100, wins2, numGames, float64(wins2)/float64(numGames)*100)
}

func main() {
	// Validate that the correct number of command-line arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./pig <strategy1> <strategy2>")
		os.Exit(1)
	}

	// Parse the strategies from command-line arguments
	strategy1, err1 := strconv.Atoi(os.Args[1])
	strategy2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Both strategies must be integers.")
		os.Exit(1)
	}

	// Seed the random number generator for dice rolls with the current time to ensure varied results
	rand.Seed(time.Now().UnixNano())

	// Define the number of games per simulation and simulate the games
	numGames := 10 // Number of games per simulation
	SimulateGames(strategy1, strategy2, numGames)
}
