package main

type terrain uint8

const (
	desert terrain = iota
	hills
	fields
	forest
	mountains
	pasture
)

type resources []int

type resource uint8

const (
	unknown resource = iota
	brick
	grain
	lumber
	ore
	wool
)

func newResources() resources {
	return make([]int, 5, 5)
}

type tile struct {
	terrain terrain
	number  int
}

type piece struct {
	color     color
	piecetype piecetype
}

type pieces []int

type piecetype uint8

const (
	settlement piecetype = iota
	city
	road
)

type board struct {
	tiles         map[hexcoord]tile
	robber        hexcoord
	intersections map[hexvertex]piece
	paths         map[hexedge]piece
}

func (b board) resourceProduction(diceRoll int) map[color]resources {
	res := map[color]resources{
		red:    newResources(),
		white:  newResources(),
		blue:   newResources(),
		yellow: newResources(),
	}
	for c, t := range b.tiles {
		if t.number != diceRoll {
			continue
		}
		if c == b.robber {
			continue
		}
		for _, v := range c.vertices() {
			piece, ok := b.intersections[v]
			if !ok {
				continue
			}
			switch piece.piecetype {
			case settlement:
				res[piece.color][int(t.terrain)-1] += 1
			case city:
				res[piece.color][int(t.terrain)-1] += 2
			}
		}
	}
	return res
}

func (b board) receiveStartingResources(v hexvertex) resources {
    res := newResources()
    for _, c := range v.hexAdjacent() {
        tile := b.tiles[c]
        res[int(tile.terrain)-1] += 1
    }
    return res
}

func (b board) build(color color, pt piecetype, v hexvertex) {
    b.intersections[v] = piece{color, pt}
}

type devcards struct{}

type color uint8

const (
	black color = iota
	red
	white
	blue
	yellow
)

type player struct {
	color         color
	hand          resources
	devcards      devcards
	pieces        pieces
	piecesReserve pieces
}

type game struct {
	board       board
	players     []player
	longestRoad int
	largestArmy int
	bank        resources
}
