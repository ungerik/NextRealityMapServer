package mapserver

// import "github.com/ungerik/json-tree"

type TileParams struct {
}

// func (self *TileParams) InitFromJSON(tree json.Tree) {
// 	self.channels.InitFromJSON(tree)
// 	self.horSamples = tree.Select("HorTileSamples").GetInt()
// 	self.verSamples = tree.Select("VerTileSamples").GetInt()
// 	self.Stride = tree.Select("TileStride").GetInt()
// }

// func (self *TileParams) WriteJSON(builder *json.Builder) {
// 	self.channels.WriteJSON(builder)
// 	builder.Name("HorTileSamples").Value(self.horSamples)
// 	builder.Name("VerTileSamples").Value(self.verSamples)
// 	builder.Name("TileStride").Value(self.Stride)
// }

// func (self *TileParams) Channels() Channels {
// 	return self.channels
// }

// func (self *TileParams) NumTileSamples() (hor, ver int) {
// 	return self.horSamples, self.verSamples
// }

// func (self *TileParams) TileStride() int {
// 	return self.Stride
// }
