package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// IronBars are blocks that serve a similar purpose to glass panes, but made of iron instead of glass.
type IronBars struct {
	transparent
	thin
}

// BreakInfo ...
func (i IronBars) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    5,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(i, 1)),
	}
}

// CanDisplace ...
func (i IronBars) CanDisplace(b world.Liquid) bool {
	_, water := b.(Water)
	return water
}

// SideClosed ...
func (i IronBars) SideClosed(cube.Pos, cube.Pos, *world.World) bool {
	return false
}

// EncodeItem ...
func (IronBars) EncodeItem() (name string, meta int16) {
	return "minecraft:iron_bars", 0
}

// EncodeBlock ...
func (IronBars) EncodeBlock() (string, map[string]interface{}) {
	return "minecraft:iron_bars", nil
}
