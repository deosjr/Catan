package main

import (
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

func TestBuild(t *testing.T) {
	for i, tt := range []struct{}{} {
		t.Logf("%d) TODO %v", i, tt)
	}
}
