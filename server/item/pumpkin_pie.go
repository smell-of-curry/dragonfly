package item

import "github.com/df-mc/dragonfly/server/world"

// PumpkinPie is a food item that can be eaten by the player.
type PumpkinPie struct {
	defaultFood
}

// Consume ...
func (PumpkinPie) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(8, 4.8)
	return Stack{}
}

// CompostChance ...
func (PumpkinPie) CompostChance() float64 {
	return 1
}

// EncodeItem ...
func (PumpkinPie) EncodeItem() (name string, meta int16) {
	return "minecraft:pumpkin_pie", 0
}
