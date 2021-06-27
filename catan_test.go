package main

import (
    "fmt"
	"reflect"
	"testing"
)

func TestResourceProduction(t *testing.T) {
	for i, tt := range []struct {
		roll  int
		board board
		bank  resources
		want  map[color]resources
	}{
		{
			roll:  5,
			board: illustrationA(),
			bank:  []int{10, 10, 10, 10, 10},
			want: map[color]resources{
				red:    []int{0, 0, 0, 0, 0},
				blue:   []int{1, 0, 0, 0, 1},
				white:  []int{0, 0, 0, 0, 1},
				yellow: []int{0, 0, 0, 0, 0},
			},
		},
		{
			roll:  8,
			board: illustrationA(),
			bank:  []int{10, 10, 10, 10, 10},
			want: map[color]resources{
				red:    []int{0, 0, 1, 0, 0},
				blue:   []int{0, 0, 1, 0, 0},
				white:  []int{0, 0, 0, 1, 0},
				yellow: []int{0, 0, 0, 0, 0},
			},
		},
        // bank does not have enough wool
		{
			roll:  5,
			board: illustrationA(),
			bank:  []int{10, 10, 10, 10, 1},
			want: map[color]resources{
				red:    []int{0, 0, 0, 0, 0},
				blue:   []int{1, 0, 0, 0, 0},
				white:  []int{0, 0, 0, 0, 0},
				yellow: []int{0, 0, 0, 0, 0},
			},
		},
        // bank does not have enough brick 
		{
			roll:  5,
			board: func(b board) board {
                // replace all settlements with cities
                for c, p := range b.intersections {
                    b.intersections[c] = piece{color:p.color, piecetype:city}
                }
                return b
            }(illustrationA()),
			bank:  []int{1, 10, 10, 10, 10},
			want: map[color]resources{
				red:    []int{0, 0, 0, 0, 0},
				blue:   []int{1, 0, 0, 0, 2},
				white:  []int{0, 0, 0, 0, 2},
				yellow: []int{0, 0, 0, 0, 0},
			},
		},
	} {
        players := []*player{newPlayer(red), newPlayer(white), newPlayer(blue), newPlayer(yellow)}
        game := game{board:tt.board, bank:tt.bank, players:players}
		game.resourceProduction(tt.roll)
        for _, p := range game.players {
		    if !reflect.DeepEqual(p.hand, tt.want[p.color]) {
			    t.Errorf("%d) got %v want %v", i, p.hand, tt.want[p.color])
		    }
        }
	}
}

func TestRawResourceProduction(t *testing.T) {
	for i, tt := range []struct {
		roll  int
		board board
		want  map[color]resources
	}{
		{
			roll:  5,
			board: illustrationA(),
			want: map[color]resources{
				red:    []int{0, 0, 0, 0, 0},
				blue:   []int{1, 0, 0, 0, 1},
				white:  []int{0, 0, 0, 0, 1},
				yellow: []int{0, 0, 0, 0, 0},
			},
		},
		{
			roll:  8,
			board: illustrationA(),
			want: map[color]resources{
				red:    []int{0, 0, 1, 0, 0},
				blue:   []int{0, 0, 1, 0, 0},
				white:  []int{0, 0, 0, 1, 0},
				yellow: []int{0, 0, 0, 0, 0},
			},
		},
		{
			roll:  5,
			board: func(b board) board {
                // replace all settlements with cities
                for c, p := range b.intersections {
                    b.intersections[c] = piece{color:p.color, piecetype:city}
                }
                return b
            }(illustrationA()),
			want: map[color]resources{
				red:    []int{0, 0, 0, 0, 0},
				blue:   []int{2, 0, 0, 0, 2},
				white:  []int{0, 0, 0, 0, 2},
				yellow: []int{0, 0, 0, 0, 0},
			},
		},
		{
			roll:  5,
			board: func(b board) board {
                // put the robber on the pasture 5
                b.robber = hexcoord{1, -2, 1}
                return b
            }(illustrationA()),
			want: map[color]resources{
				red:    []int{0, 0, 0, 0, 0},
				blue:   []int{1, 0, 0, 0, 0},
				white:  []int{0, 0, 0, 0, 0},
				yellow: []int{0, 0, 0, 0, 0},
			},
		},
	} {
		got := tt.board.resourceProduction(tt.roll)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d) got %v want %v", i, got, tt.want)
		}
	}
}

func TestReceiveStartingResources(t *testing.T) {
	for i, tt := range []struct {
        board board
        vertex hexvertex
		want  resources
	}{
		{
            board: illustrationA(),
            vertex: hexvertex{c:hexcoord{-2,0,2}, top:true},
            want: []int{1,0,1,1,0},
        },
    }{
		got := tt.board.receiveStartingResources(tt.vertex)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d) got %v want %v", i, got, tt.want)
		}
    }
}

func TestTrade(t *testing.T) {
	for i, tt := range []struct{}{} {
		t.Logf("%d) TODO %v", i, tt)
	}
}

func TestBuildSettlement(t *testing.T) {
	for i, tt := range []struct {
        board board
        player *player
        vertex hexvertex
		want   error
	}{
		{
            board: illustrationA(),
            player: &player{color: red, hand: settlement.cost()},
            vertex: hexvertex{c:hexcoord{-2,0,2}, top:false},
        },
		{
            board: illustrationA(),
            player: &player{color: red, hand: settlement.cost()},
            vertex: hexvertex{c:hexcoord{-2,0,2}, top:true},
            want: fmt.Errorf("intersection already built"),
        },
		{
            board: illustrationA(),
            player: &player{color: red, hand: newResources()},
            vertex: hexvertex{c:hexcoord{-2,0,2}, top:true},
            want: fmt.Errorf("player cant pay cost"),
        },
    }{
		got := tt.board.buildSettlement(tt.player, tt.vertex)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d) got %v want %v", i, got, tt.want)
		}
        if tt.want == nil {
            // verify settlement was built if no error
            p, ok := tt.board.intersections[tt.vertex]
            wantPiece := piece{red, settlement}
            if !ok || p != wantPiece {
                t.Errorf("%d) got %v want red settlement", i, p)
            }
        }
    }
}
