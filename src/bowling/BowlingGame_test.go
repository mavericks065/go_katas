package bowling_test

import (
	. "bowling"
	"strconv"
	"testing"
)

func Test_GutterGame(t *testing.T) {
	game := Game{}

	rollMany(&game, 20, 0)

	score := game.Score()

	if score != 0 {
		t.Errorf("Test_GutterGame does not return 0")
	}
}

func Test_AllOnes(t *testing.T) {
	game := Game{}

	rollMany(&game, 20, 1)

	score := game.Score()

	if score != 20 {
		t.Errorf("Test_AllOnes does not return 20")
	}
}

func Test_OneSpare(t *testing.T) {
	game := Game{}

	rollSpare(&game)
	game.Roll(3)
	rollMany(&game, 17, 0)

	scoreOneSpare := game.Score()

	if scoreOneSpare != 16 {
		t.Log("Obtained Score is " + strconv.Itoa(scoreOneSpare))
		t.Errorf("Test_OneSpare does not return 16")
	}
}

func Test_OneStrike(t *testing.T) {
	game := Game{}

	rollStrike(&game)
	game.Roll(3)
	game.Roll(4)
	rollMany(&game, 16, 0)

	scoreOneStrike := game.Score()

	if scoreOneStrike != 24 {
		t.Log("Obtained Score is " + strconv.Itoa(scoreOneStrike))
		t.Errorf("Test_OneSpare does not return 24")
	}
}

func Test_PerfectGame(t *testing.T) {
	game := Game{}

	rollMany(&game, 12, 10)

	perfectScore := game.Score()

	if perfectScore != 300 {
		t.Log("Obtained Score is " + strconv.Itoa(perfectScore))
		t.Errorf("Test_OneSpare does not return 300")
	}
}

func rollMany(game *Game, n, pins int) {
	for i := 0; i < n; i++ {
		game.Roll(pins)
	}
}

func rollSpare(game *Game) {
	game.Roll(5)
	game.Roll(5)
}

func rollStrike(game *Game) {
	game.Roll(10)
}
