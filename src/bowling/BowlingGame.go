package bowling

// Game structure
type Game struct {
	rolls       [21]int
	currentRoll int
}

// Roll function
func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

// Score function add up the score when required
func (g *Game) Score() int {
	score := 0
	frameIndex := 0
	for frame := 0; frame < 10; frame++ {
		if g.isSpare(frameIndex) {
			score += 10 + g.rolls[frameIndex+2]
			frameIndex += 2
		} else if g.isStrike(frameIndex) {
			score += 10 + g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
			frameIndex++
		} else {
			score += g.rolls[frameIndex] + g.rolls[frameIndex+1]
			frameIndex += 2
		}
	}
	return score
}

func (g *Game) isSpare(rollIndex int) bool {
	return (g.rolls[rollIndex]+g.rolls[rollIndex+1] == 10)
}

func (g *Game) isStrike(rollIndex int) bool {
	return (g.rolls[rollIndex] == 10)
}
