package main

import (
    "fmt"
    "strings"
)

type terrain uint8

const (
	desert terrain = iota
	hills
	fields
	forest
	mountains
	pasture
)

func (t terrain) toIndex() int {
    return int(t)-1
}

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

func (r resource) String() string {
    switch r {
    case brick: return "brick"
    case grain: return "grain"
    case lumber: return "lumber"
    case ore: return "ore"
    case wool: return "wool"
    }
    return "wrong resource"
}

func (r resource) toIndex() int {
    return int(r)-1
}

func newResources() resources {
	return make([]int, 5, 5)
}

func (r resources) add(rr resources) resources {
    for i, v := range rr {
        r[i] += v
    }
    return r
}

func (r resources) sub(rr resources) resources {
    for i, v := range rr {
        r[i] -= v
    }
    return r
}

func (r resources) covers(cost resources) bool {
    for i, v := range cost {
        if r[i] < v {
            return false
        }
    }
    return true
}

func (r resources) isEmpty() bool {
    return newResources().covers(r)
}

func (r resources) String() string {
    s := []string{}
    for i, v := range r {
        if v == 0 {
            continue
        }
        s = append(s, fmt.Sprintf("%d %s", v, resource(i+1)))
    }
    return strings.Join(s, ", ")
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

func (p piecetype) cost() resources {
    switch p {
    case settlement:
        return resources([]int{1,1,1,0,1})
    case city:
        return resources([]int{0,0,0,3,2})
    case road:
        return resources([]int{1,0,1,0,0})
    }
    return newResources()
}

type board struct {
	tiles         map[hexcoord]tile
	robber        hexcoord
	intersections map[hexvertex]piece
	paths         map[hexedge]piece
}

func (b board) resourceProduction(diceRoll int) map[color]resources {
	resPerPlayer := map[color]resources{
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
				resPerPlayer[piece.color][t.terrain.toIndex()] += 1
			case city:
				resPerPlayer[piece.color][t.terrain.toIndex()] += 2
			}
		}
	}
	return resPerPlayer
}

func (b board) receiveStartingResources(v hexvertex) resources {
    res := newResources()
    for _, c := range v.hexAdjacent() {
        tile := b.tiles[c]
        res[tile.terrain.toIndex()] += 1
    }
    return res
}

func (b board) buildSettlement(player *player, v hexvertex) error {
    if !player.hand.covers(settlement.cost()) {
        return fmt.Errorf("player cant pay cost")
    }
    if _, ok := b.intersections[v]; ok {
        return fmt.Errorf("intersection already built")
    }
    player.hand = player.hand.sub(settlement.cost())
    // TODO: distance rule
    b.intersections[v] = piece{player.color, settlement}
    return nil
}

func (b board) buildCity(player *player, v hexvertex) error {
    if !player.hand.covers(city.cost()) {
        return fmt.Errorf("player cant pay cost")
    }
    p, ok := b.intersections[v]
    if!ok || p.piecetype != settlement || p.color != player.color {
        return fmt.Errorf("intersection does not contain settlement of player's color")
    }
    player.hand = player.hand.sub(city.cost())
    b.intersections[v] = piece{player.color, city}
    return nil
}

func (b board) buildRoad(player *player, e hexedge) error {
    if !player.hand.covers(road.cost()) {
        return fmt.Errorf("player cant pay cost")
    }
    if _, ok := b.paths[e]; ok {
        return fmt.Errorf("road already built")
    }
    player.hand = player.hand.sub(road.cost())
    b.paths[e] = piece{player.color, road}
    return nil
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

func (c color) String() string {
    switch c {
    case red: return "red"
    case white: return "white"
    case blue: return "blue"
    case yellow: return "yellow"
    }
    return "wrong color"
}

type player struct {
	color         color
	hand          resources
	devcards      devcards
	pieces        pieces
	piecesReserve pieces
}

func newPlayer(color color) *player {
    return &player{
        color: color,
        hand: newResources(),
    }
}

type game struct {
	board       board
    order       []color
	players     map[color]*player
	longestRoad int
	largestArmy int
	bank        resources
}

func (g game) resourceProduction(roll int) map[color]resources {
    raw := g.board.resourceProduction(roll)
    total := newResources()
    for _, res := range raw {
        total.add(res)
    }
    for i, r := range total {
        if g.bank[i] < r {
            // bank does not have enough resource cards
            numPlayers := 0
            for _, res := range raw {
                if res[i] > 0 {
                    numPlayers++
                }
            }
            for _, res := range raw {
                // only one player gets this resource: gets all remaining cards in bank
                if numPlayers == 1 && res[i] > 0 {
                    res[i] = g.bank[i]
                // multiple players would receive resource: no one gets any cards
                } else {
                    res[i] = 0
                }
            }
        }
    }
    return raw
}
