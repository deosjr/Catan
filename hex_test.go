package main

import (
	"reflect"
	"testing"
)

// abusing reflect, only in tests :)
func TestHexCoordinates(t *testing.T) {
	for i, tt := range []struct {
		start hexcoord
		instr []string
		want  hexcoord
	}{
		{
			start: hexcoord{0, 0, 0},
			instr: []string{"Left"},
			want:  hexcoord{-1, +1, 0},
		},
		{
			start: hexcoord{0, +1, -1},
			instr: []string{"DownRight", "Right", "UpLeft"},
			want:  hexcoord{+1, 0, -1},
		},
	} {
		v := tt.start
		for _, str := range tt.instr {
			ts := reflect.TypeOf(&v)
			vs := reflect.ValueOf(&v)
			m, ok := ts.MethodByName(str)
			if !ok {
				t.Fatalf("called method %s but doesnt exist!", str)
			}
			v = m.Func.Call([]reflect.Value{vs})[0].Interface().(hexcoord)
		}
		if v != tt.want {
			t.Errorf("%d): got %v want %v", i, v, tt.want)
		}
	}
}

func TestHexVertices(t *testing.T) {
	for i, tt := range []struct {
		start hexvertex
		instr []string
		want  hexvertex
	}{
		{
			start: hexvertex{c: hexcoord{0, 0, 0}, top: true},
			instr: []string{"Left"},
			want:  hexvertex{c: hexcoord{0, +1, -1}, top: false},
		},
		{
			start: hexvertex{c: hexcoord{0, 0, 0}, top: true},
			instr: []string{"Left", "Down", "Left", "Down"},
			want:  hexvertex{c: hexcoord{-2, 0, +2}, top: true},
		},
	} {
		v := tt.start
		for _, str := range tt.instr {
			ts := reflect.TypeOf(&v)
			vs := reflect.ValueOf(&v)
			m, ok := ts.MethodByName(str)
			if !ok {
				t.Fatalf("called method %s but doesnt exist!", str)
			}
			v = m.Func.Call([]reflect.Value{vs})[0].Interface().(hexvertex)
		}
		if v != tt.want {
			t.Errorf("%d): got %v want %v", i, v, tt.want)
		}
	}
}

func TestHexEdges(t *testing.T) {

}
