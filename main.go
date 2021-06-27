package main

//import "fmt"

func main() {
	board := illustrationA()
    game := game{
        board:board,
        players:map[color]*player{
            red: newPlayer(red),
            blue: newPlayer(blue),
            white: newPlayer(white),
            yellow: newPlayer(yellow),
        },
        bank: resources([]int{19, 19, 19, 19, 19}),
    }
    game.players[red].hand = game.players[red].hand.add([]int{0,1,2,0,0})
    game.players[blue].hand = game.players[blue].hand.add([]int{1,0,1,1,0})
    game.players[white].hand = game.players[white].hand.add([]int{1,1,1,0,0})
    game.players[yellow].hand = game.players[yellow].hand.add([]int{0,2,0,1,0})
	board.print()
    loop(game)
}

// returns a full board as per illustration A of the base rules
func illustrationA() board {
	tiles := map[hexcoord]tile{}
	for i, t := range []tile{{mountains, 10}, {pasture, 2}, {forest, 9}} {
		c := hexcoord{x: i, y: 2 - i, z: -2}
		tiles[c] = t
	}
	for i, t := range []tile{{fields, 12}, {hills, 6}, {pasture, 4}, {hills, 10}} {
		c := hexcoord{x: -1 + i, y: 2 - i, z: -1}
		tiles[c] = t
	}
	for i, t := range []tile{{fields, 9}, {forest, 11}, {desert, 0}, {forest, 3}, {mountains, 8}} {
		c := hexcoord{x: -2 + i, y: 2 - i, z: 0}
		tiles[c] = t
	}
	for i, t := range []tile{{forest, 8}, {mountains, 3}, {fields, 4}, {pasture, 5}} {
		c := hexcoord{x: -2 + i, y: 1 - i, z: 1}
		tiles[c] = t
	}
	for i, t := range []tile{{hills, 5}, {fields, 6}, {pasture, 11}} {
		c := hexcoord{x: -2 + i, y: -i, z: 2}
		tiles[c] = t
	}
	// piece coordinates given as relative from robber tile
	robbertop := hexvertex{c: hexcoord{0, 0, 0}, top: true}
	robberdown := hexvertex{c: hexcoord{0, 0, 0}, top: false}
	intersections := map[hexvertex]piece{}
	paths := map[hexedge]piece{}

	redhouse1 := robbertop.Up().Left()
	intersections[redhouse1] = piece{red, settlement}
	paths[redhouse1.RightEdge()] = piece{red, road}
	redhouse2 := robberdown.Left().Left().Left()
	intersections[redhouse2] = piece{red, settlement}
	paths[redhouse2.RightEdge()] = piece{red, road}

	yellowhouse1 := robbertop.Right().Right().Up()
	intersections[yellowhouse1] = piece{yellow, settlement}
	paths[yellowhouse1.LeftEdge()] = piece{yellow, road}
	yellowhouse2 := robberdown.Down()
	intersections[yellowhouse2] = piece{yellow, settlement}
	paths[yellowhouse2.RightEdge()] = piece{yellow, road}

	whitehouse1 := robbertop.Left().Left()
	intersections[whitehouse1] = piece{white, settlement}
	paths[whitehouse1.LeftEdge()] = piece{white, road}
	whitehouse2 := robberdown.Right().Right().Right()
	intersections[whitehouse2] = piece{white, settlement}
	paths[whitehouse2.UpEdge()] = piece{white, road}

	bluehouse1 := robberdown.Left().Left().Down()
	intersections[bluehouse1] = piece{blue, settlement}
	paths[bluehouse1.RightEdge()] = piece{blue, road}
	bluehouse2 := robberdown.Right().Right().Down()
	intersections[bluehouse2] = piece{blue, settlement}
	paths[bluehouse2.UpEdge()] = piece{blue, road}

	// todo: harbours

	return board{
		tiles:         tiles,
		robber:        hexcoord{0, 0, 0},
		intersections: intersections,
		paths:         paths,
	}
}
