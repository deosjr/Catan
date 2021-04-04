package main

// as always, https://www.redblobgames.com/grids/hexagons/
// pointy-topped
type hexcoord struct {
	x, y, z int
}

func (c hexcoord) Left() hexcoord {
	return hexcoord{x: c.x - 1, y: c.y + 1, z: c.z}
}

func (c hexcoord) Right() hexcoord {
	return hexcoord{x: c.x + 1, y: c.y - 1, z: c.z}
}

func (c hexcoord) UpLeft() hexcoord {
	return hexcoord{x: c.x, y: c.y + 1, z: c.z - 1}
}

func (c hexcoord) UpRight() hexcoord {
	return hexcoord{x: c.x + 1, y: c.y, z: c.z - 1}
}

func (c hexcoord) DownLeft() hexcoord {
	return hexcoord{x: c.x - 1, y: c.y, z: c.z + 1}
}

func (c hexcoord) DownRight() hexcoord {
	return hexcoord{x: c.x, y: c.y - 1, z: c.z + 1}
}

// first two are used in printTile()
func (c hexcoord) vertices() []hexvertex {
	return []hexvertex{
		{c: c, top: true},
		{c: c.UpLeft(), top: false},
		{c: c, top: false},
		{c: c.UpRight(), top: false},
		{c: c.DownLeft(), top: true},
		{c: c.DownRight(), top: true},
	}
}

// each vertex is defined by for which hex it is a top/bottom corner
type hexvertex struct {
	c   hexcoord
	top bool //false = bottom
}

func (v hexvertex) Left() hexvertex {
	if v.top {
		return hexvertex{c: v.c.UpLeft(), top: false}
	}
	// bottom
	return hexvertex{c: v.c.DownLeft(), top: true}
}

func (v hexvertex) Right() hexvertex {
	if v.top {
		return hexvertex{c: v.c.UpRight(), top: false}
	}
	// bottom
	return hexvertex{c: v.c.DownRight(), top: true}
}

func (v hexvertex) Up() hexvertex {
	if !v.top {
		panic("invalid hexvertex movement")
	}
	return hexvertex{c: v.c.UpLeft().UpRight(), top: false}
}

func (v hexvertex) Down() hexvertex {
	if v.top {
		panic("invalid hexvertex movement")
	}
	return hexvertex{c: v.c.DownLeft().DownRight(), top: true}
}

func (v hexvertex) LeftEdge() hexedge {
	if v.top {
		return hexedge{c: v.c, upleft: true}
	}
	// bottom
	return hexedge{c: v.c.DownLeft(), upright: true}
}

func (v hexvertex) RightEdge() hexedge {
	if v.top {
		return hexedge{c: v.c, upright: true}
	}
	// bottom
	return hexedge{c: v.c.DownRight(), upleft: true}
}

func (v hexvertex) UpEdge() hexedge {
	if !v.top {
		panic("invalid hexvertex movement")
	}
	return hexedge{c: v.c.UpRight(), left: true}
}

func (v hexvertex) DownEdge() hexedge {
	if v.top {
		panic("invalid hexvertex movement")
	}
	return hexedge{c: v.c.DownRight(), left: true}
}

// similarly for edges, we use the upper left half of each hex
type hexedge struct {
	c hexcoord
	// only one can be true
	left    bool
	upleft  bool
	upright bool
}
