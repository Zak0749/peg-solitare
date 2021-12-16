package types

type Tile struct {
	Value  bool
	Exists bool
}

type Board [9][9]Tile
