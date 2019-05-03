package main

import (
	"fmt"
	"io/ioutil"
)

type marble struct {
	score      int
	prev, next *marble
}

func main() {
	var maxPlayers, lastMarbleWorth int

	input, _ := ioutil.ReadFile("input.txt")
	fmt.Sscanf(
		string(input),
		"%d players; last marble is worth %d points",
		&maxPlayers, &lastMarbleWorth,
	)

	currentPlayer := 0
	scores := make(map[int]int)

	firstMarble := &marble{score: 0, prev: nil, next: nil}
	firstMarble.prev = firstMarble
	firstMarble.next = firstMarble

	lastPlayedMarble := firstMarble

	for i := 1; i <= lastMarbleWorth; i++ {
		newMarble := &marble{score: i, next: nil, prev: nil}

		if i%23 == 0 {
			scores[currentPlayer] += newMarble.score

			toRemove := lastPlayedMarble.prev.prev.prev.prev.prev.prev.prev
			scores[currentPlayer] += toRemove.score

			toRemove.prev.next = toRemove.next
			toRemove.next.prev = toRemove.prev

			lastPlayedMarble = toRemove.next
		} else {
			lastPlayedMarble.next.next.prev = newMarble
			newMarble.next = lastPlayedMarble.next.next

			lastPlayedMarble.next.next = newMarble
			newMarble.prev = lastPlayedMarble.next

			lastPlayedMarble = newMarble
		}

		currentPlayer = (currentPlayer + 1) % maxPlayers
	}

	fmt.Println("Part 1: ", highestScore(scores))
}

func highestScore(scores map[int]int) int {
	highest := -1
	for _, score := range scores {
		if score > highest {
			highest = score
		}
	}
	return highest
}
