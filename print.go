package main

import (
	"fmt"
	"strings"

	fatih "github.com/fatih/color"
)

func (p piece) getColor() *fatih.Color {
	switch p.color {
	case red:
		return fatih.New(fatih.FgHiRed)
	case white:
		return fatih.New(fatih.FgHiWhite)
	case blue:
		return fatih.New(fatih.FgBlue)
	case yellow:
		return fatih.New(fatih.FgYellow)
	}
	panic("wrong color")
}

func (p piece) print() {
	c := p.getColor()
	switch p.piecetype {
	case settlement:
		c.Print("n")
	case city:
		c.Print("M")
	}
}

func (b board) print() {
	b.printTileRow(hexcoord{0, 2, -2}, 3, false)
	b.printTileRow(hexcoord{-1, 2, -1}, 4, false)
	b.printTileRow(hexcoord{-2, 2, 0}, 5, false)
	b.printTileRow(hexcoord{-2, 1, 1}, 4, true)
	b.printTileRow(hexcoord{-2, 0, 2}, 3, true)
	b.printTileRow(hexcoord{-2, -1, 3}, 2, true)
}

func (b board) printTileRow(start hexcoord, n int, lowerhalf bool) {
	printFuncs := []func(hexcoord){b.printTileTopVertex, b.printTileTopEdges, b.printTileUpLeftVertex, b.printTileLeftEdge}
	prefix := "   "
	for j, f := range printFuncs {
		leftmost := hexcoord{x: start.x - 1, y: start.y + 1, z: start.z}
		if j == 0 {
			// printing top vertex
			if lowerhalf {
				fmt.Print(strings.Repeat(prefix, 5-n-1))
				f(leftmost)
			} else {
				fmt.Print(strings.Repeat(prefix, 5-n+1))
			}
		} else if j == 1 && lowerhalf {
			// printing topedges for lower half
			fmt.Print(strings.Repeat(prefix, 5-n-1))
			topleftmost := leftmost.vertices()[0]
			if road, ok := b.paths[topleftmost.RightEdge()]; ok {
				road.getColor().Print(" \\ ")
			} else {
				fmt.Print(prefix)
			}
		} else {
			fmt.Print(strings.Repeat(prefix, 5-n))
		}
		for i := 0; i < n; i++ {
			c := hexcoord{x: start.x + i, y: start.y - i, z: start.z}
			f(c)
		}
		rightmost := hexcoord{x: start.x + n, y: start.y - n, z: start.z}
		if j == 0 && lowerhalf {
			f(rightmost)
		}
		if j == 1 && lowerhalf {
			f(rightmost)
		}
		if j == 2 {
			// printing upleftvertex
			f(rightmost)
		}
		if j == 3 {
			// printing leftedge
			f(rightmost)
		}
		fmt.Println()
	}
}

func (b board) printTileTopVertex(c hexcoord) {
	vertices := c.vertices()
	top := vertices[0]
	if piece, ok := b.intersections[top]; ok {
		piece.print()
	} else {
		fmt.Print(".")
	}
	fmt.Print("     ")
}

func (b board) printTileTopEdges(c hexcoord) {
	vertices := c.vertices()
	top := vertices[0]
	if road, ok := b.paths[top.LeftEdge()]; ok {
		road.getColor().Print(" / ")
	} else {
		fmt.Print("   ")
	}
	if road, ok := b.paths[top.RightEdge()]; ok {
		road.getColor().Print("  \\")
	} else {
		fmt.Print("   ")
	}
}

func (b board) printTileUpLeftVertex(c hexcoord) {
	vertices := c.vertices()
	upleft := vertices[1]
	if piece, ok := b.intersections[upleft]; ok {
		piece.print()
	} else {
		fmt.Print(".")
	}
	fmt.Print("     ")
}

func (b board) printTileLeftEdge(c hexcoord) {
	vertices := c.vertices()
	upleft := vertices[1]
	if road, ok := b.paths[upleft.DownEdge()]; ok {
		road.getColor().Print("| ")
	} else {
		fmt.Print("  ")
	}
	tile, ok := b.tiles[c]
	if !ok {
		return
	}
	switch tile.terrain {
	case desert:
		fatih.New(fatih.FgBlack, fatih.BgYellow).Print("~~~")
	case hills:
		fatih.New(fatih.FgBlack, fatih.BgRed).Printf("H%2d", tile.number)
	case fields:
		fatih.New(fatih.FgBlack, fatih.BgHiYellow).Printf("G%2d", tile.number)
	case forest:
		fatih.New(fatih.FgBlack, fatih.BgGreen).Printf("F%2d", tile.number)
	case mountains:
		fatih.New(fatih.FgBlack, fatih.BgHiBlack).Printf("M%2d", tile.number)
	case pasture:
		fatih.New(fatih.FgBlack, fatih.BgHiGreen).Printf("P%2d", tile.number)
	}
	fmt.Print(" ")
}
