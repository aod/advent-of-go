package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/aoktayd/adventofgode/internal/error"
	"github.com/aoktayd/adventofgode/internal/input"
)

const guardShiftTimeLayout = "2006-01-02 15:04"

type sleepRecord struct {
	guard     int
	timestamp time.Time
	asleep    bool
}

type byTimestamp []sleepRecord

func (a byTimestamp) Len() int           { return len(a) }
func (a byTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byTimestamp) Less(i, j int) bool { return a[i].timestamp.Before(a[j].timestamp) }

func parseInput() []sleepRecord {
	var sleepRecords []sleepRecord

	s, f := input.Scanner(input.Puzzle{Year: 2018, Day: 4})
	defer f.Close()

	for s.Scan() {
		line := s.Text()

		split := strings.SplitAfter(line, "]")
		timestamp, err := time.Parse(guardShiftTimeLayout, strings.Trim(split[0], "[]"))
		error.Check(err)

		var guardID int
		asleep := false
		rest := split[1]
		if strings.Contains(rest, "Guard") {
			fmt.Sscanf(rest, " Guard #%d", &guardID)
		} else if strings.Contains(rest, "asleep") {
			asleep = true
		}

		sleepRecords = append(sleepRecords, sleepRecord{guardID, timestamp, asleep})
	}

	return sleepRecords
}

func main() {
	sleepRecords := parseInput()

	// Prerequisite: Sort the input by date
	// This is needed to analyse the sleeping pattern of the guards later on
	sort.Sort(byTimestamp(sleepRecords))

	// Part 1 - Find the guard that has the most minutes asleep.
	// 					What minute does that guard spend asleep the most?

	// 1.1 Analyse the sleeping patterns
	// To keep track of how often a guard goes to sleep on a specfic
	// minute i've come up with the following map data structure
	// GuardID: []Minute: Amount
	sleepPatterns := make(map[int]map[int]int)

	currentGuard := 0
	sleepMinute := 0
	for _, sleepRecord := range sleepRecords {
		if sleepRecord.guard != 0 {
			currentGuard = sleepRecord.guard
			continue
		}

		if sleepRecord.asleep {
			sleepMinute = sleepRecord.timestamp.Minute()
			continue
		}

		sleepPattern, sleepPatternExists := sleepPatterns[currentGuard]
		if !sleepPatternExists {
			sleepPattern = make(map[int]int)
			sleepPatterns[currentGuard] = sleepPattern
		}
		for i := sleepMinute; i < sleepRecord.timestamp.Minute(); i++ {
			sleepPattern[i]++
		}
	}

	// 1.2 Which guard slept the most?
	var sleepiestGuard int
	biggestSleepTime := 0

	for guardID, sleepPattern := range sleepPatterns {
		currentSleepTime := 0
		for _, minutesSlept := range sleepPattern {
			currentSleepTime += minutesSlept
		}

		if currentSleepTime > biggestSleepTime {
			biggestSleepTime = currentSleepTime
			sleepiestGuard = guardID
		}
	}

	// 1.3 At which minute did the sleepiest guard sleep the longest
	var sleepiestMinute int
	biggestSleep := 0
	for minute, currentSleep := range sleepPatterns[sleepiestGuard] {
		if currentSleep > biggestSleep {
			biggestSleep = currentSleep
			sleepiestMinute = minute
		}
	}

	fmt.Println("Part 1", sleepiestGuard*sleepiestMinute)
}
