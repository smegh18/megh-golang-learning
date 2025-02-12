package main

import (
	"math/rand"
	"testing"
	"time"
)

// TestRollDice ensures that RollDice produces values between 1 and 6
func TestRollDice(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		diceRoll := RollDice()
		if diceRoll < 1 || diceRoll > 6 {
			t.Errorf("RollDice() returned %d, expected between 1 and 6", diceRoll)
		}
	}
}

// TestPlayTurn ensures that PlayTurn correctly follows game rules
func TestPlayTurn(t *testing.T) {
	player := &Player{Name: "TestPlayer", Strategy: 10}
	rand.Seed(1) // Set seed for deterministic results
	score := PlayTurn(player)
	if score < 0 || score > player.Strategy {
		t.Errorf("PlayTurn() returned %d, expected between 0 and %d", score, player.Strategy)
	}
}

// TestPlayGame ensures that a game completes and a winner is determined
func TestPlayGame(t *testing.T) {
	player1 := &Player{Name: "Player1", Strategy: 10}
	player2 := &Player{Name: "Player2", Strategy: 15}
	rand.Seed(1)

	winner := PlayGame(player1, player2)
	if winner != "Player1" && winner != "Player2" {
		t.Errorf("PlayGame() returned unexpected winner: %s", winner)
	}
}

// TestSimulateGames ensures that multiple games are simulated without errors
func TestSimulateGames(t *testing.T) {
	rand.Seed(1)
	SimulateGames(10, 15, 5) // Simulating 5 games with strategies 10 and 15
}
