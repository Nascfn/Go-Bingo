package logic

import (
	"math/rand"
	"time"
)

type Player struct {
	name       string
	B_nums     []int
	I_nums     []int
	N_nums     []int
	G_nums     []int
	O_nums     []int
	joinedGame *Game
}

func (p *Player) GenerateCard() {
	B_holder := sliceOnRange(1, 15)
	shuffleInts(B_holder)
	p.B_nums = B_holder[:5]
	I_holder := sliceOnRange(16, 30)
	shuffleInts(I_holder)
	p.I_nums = I_holder[:5]
	N_holder := sliceOnRange(31, 45)
	shuffleInts(N_holder)
	p.N_nums = N_holder[:4]
	G_holder := sliceOnRange(46, 60)
	shuffleInts(G_holder)
	p.G_nums = G_holder[:5]
	O_holder := sliceOnRange(61, 75)
	shuffleInts(O_holder)
	p.O_nums = O_holder[:5]
}

func NewPlayer(name string) *Player {
	p := &Player{name: name}
	p.GenerateCard()
	return p
}

func (p *Player) JoinGame(g *Game) {
	p.joinedGame = g
	g.AddPlayer(p)
}

func (p *Player) GetGame() *Game {
	return p.joinedGame
}

func shuffleInts(nums []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

func sliceOnRange(start, end int) []int {
	if end < start {
		return []int{}
	}

	size := end - start + 1
	nums := make([]int, size)

	for i := 0; i < size; i++ {
		nums[i] = start + i
	}

	return nums
}
