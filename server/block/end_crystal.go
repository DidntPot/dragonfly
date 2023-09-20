package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"hash/fnv"
)

// EndCrystal represents the End Crystal block in Minecraft.
type EndCrystal struct {
	transparent
}

// Hash ...
func (e EndCrystal) Hash() uint64 {
	h := fnv.New64a()
	h.Write([]byte("minecraft:end_crystal"))
	return h.Sum64()
}

// Model ...
func (e EndCrystal) Model() world.BlockModel {
	return &endCrystalModel{}
}

// EncodeItem ...
func (e EndCrystal) EncodeItem() (name string, meta int16) {
	return "minecraft:end_crystal", 0
}

// EncodeBlock ...
func (e EndCrystal) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:end_crystal", nil
}

// BreakInfo ...
func (e EndCrystal) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness: 0.0,
	}
}

// UseOnBlock ...
func (e EndCrystal) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) bool {
	if face == cube.FaceDown || face == cube.FaceNorth || face == cube.FaceEast || face == cube.FaceSouth || face == cube.FaceWest {
		return false
	}

	if _, ok := w.Block(pos).(Obsidian); !ok {
		return false
	}

	ent := w.EntityRegistry().Config().EndCrystal
	w.AddEntity(ent(pos.Vec3()))

	user.(*player.Player).Message("test...")

	return true
}

type endCrystalModel struct{}

// BBox ...
func (m *endCrystalModel) BBox(pos cube.Pos, w *world.World) []cube.BBox {
	return nil
}

// FaceSolid ...
func (m *endCrystalModel) FaceSolid(pos cube.Pos, face cube.Face, w *world.World) bool {
	return false
}
