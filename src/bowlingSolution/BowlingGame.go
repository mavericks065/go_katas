package bowlingSolution

// Game structure
type Game struct {
	rolls       [maxThrowsPerGame]int
	currentRoll int
}

const (
	// allPins is the number of pins allocated per fresh throw.
	allPins = 10

	// framesPerGame is the number of frames per bowling game.
	framesPerGame = 10

	// maxThrowsPerGame is the maximum number of throws possible in a single game.
	maxThrowsPerGame = 21
)

// Roll function
func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

// Score function add up the score when required
func (g *Game) Score() int {
	score := 0
	frameIndex := 0
	for frame := 0; frame < framesPerGame; frame++ {
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
	return (g.rolls[rollIndex]+g.rolls[rollIndex+1] == allPins)
}

func (g *Game) isStrike(rollIndex int) bool {
	return (g.rolls[rollIndex] == allPins)
}

func main() {
	panic("Run test instead!")
}
