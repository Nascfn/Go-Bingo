package logic

type Game struct {
	players     []*Player
	gameNumbers []int
}

func (g *Game) GenerateNumbers() {
	tempNumbers := sliceOnRange(1, 75)
	shuffleInts(tempNumbers)
	g.gameNumbers = tempNumbers[:75]
}

func (g *Game) AddPlayer(p *Player) {
	g.players = append(g.players, p)
}

func (g *Game) GetPlayers() []*Player {
	return g.players
}

func NewGame() *Game {
	g := &Game{}
	return g
}
