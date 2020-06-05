package main

import (
	"math/rand"
	"time"
)

const (
	nCards int = 52
)

// Game is a struct to manage the information of a holdem game
type Game struct {
	nPlayers int
	players  []player
	table    []int
	deck     []int
}

type player struct {
	pocketCards []int
	bank        int
	bet         int
}

// NewGame initializes a new game struct and returns it
func NewGame(nPlayers, startingBank int) (g Game) {
	g.nPlayers = nPlayers
	for i := 0; i < nPlayers; i++ {
		g.players = append(g.players, newPlayer(startingBank))
	}

	g.resetCards()
	return
}

func newPlayer(startingBank int) player {
	var p player
	p.bank = startingBank
	return p
}

func (g *Game) resetCards() {
	// Reset table
	g.table = make([]int, 5, 5)

	// Reset pocket cards
	for i := 0; i < g.nPlayers; i++ {
		g.players[i].pocketCards = make([]int, 2, 2)
	}

	// Reset deck
	g.deck = make([]int, nCards, nCards)
	for i := 0; i < nCards; i++ {
		g.deck[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(nCards, func(i, j int) {
		g.deck[i], g.deck[j] = g.deck[j], g.deck[i]
	})
}
