package entity

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// NewEndCrystal creates a new EndCrystal entity at the given position with the specified owner.
func NewEndCrystal(pos mgl64.Vec3) *Ent {
	return Config{
		Behaviour: crystalConfig.New(),
	}.New(EndCrystal{}, pos)
}

var crystalConfig = PassiveBehaviourConfig{}

// EndCrystal represents an End Crystal entity in the Minecraft world.
type EndCrystal struct {
}

// EncodeEntity returns the string identifier for the EndCrystal entity.
func (EndCrystal) EncodeEntity() string {
	return "minecraft:end_crystal"
}

// BBox returns the bounding box for the EndCrystal entity relative to its position.
func (EndCrystal) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.25, 0, -0.25, 0.25, 0.875, 0.25)
}
