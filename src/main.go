package main

import (
	"fmt"
	"os"
)

func NextTower(from TowerID, to TowerID) TowerID {
	ids := []TowerID{TowerID(A), TowerID(B), TowerID(C)}
	
	for _, s := range ids {
		if s == from || s == to {
			continue
		}

		return s
	}

	os.Exit(1)
	return TowerID(A)
}

func HanoiTower(towers *Towers, fromID TowerID, toID TowerID) {
	from := towers.Tower(fromID)
	to := towers.Tower(toID)

	if from.Size() == 1 {
		towers.Move(from, to)
		return
	}

	towers.Decrease()

	HanoiTower(
		towers,
		fromID,
		NextTower(fromID, toID),
	)

	towers.Increase()
	towers.Move(from, to)
	towers.Decrease()

	HanoiTower(
		towers,
		NextTower(fromID, toID),
		toID,
	)

	towers.Increase()
}

func main() {
	size := 8
	towers := NewTowers(size)

	HanoiTower(towers, TowerID(A), TowerID(C))
	fmt.Printf("Tower of Hanoi with %d disks takes %d moves to be completed", size, towers.Moves())
}