package mapserver

type Tile struct {
	Map  *Map
	Pos  IntPoint
	Data []byte
}
