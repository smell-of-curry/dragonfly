package entity

import (
	"github.com/df-mc/dragonfly/server/entity/physics"
	"github.com/go-gl/mathgl/mgl64"
)

// Text is an entity that only displays floating text. The entity is otherwise invisible and cannot be moved.
type Text struct {
	transform
	text string
}

// NewText creates and returns a new Text entity with the text and position provided.
func NewText(text string, pos mgl64.Vec3) *Text {
	t := &Text{text: text}
	t.transform = newTransform(t, pos)
	return t
}

// Name returns the name of the text entity, including the text written on it.
func (t *Text) Name() string {
	return "Text(" + t.text + ")"
}

// EncodeEntity returns the ID for falling blocks.
func (t *Text) EncodeEntity() string {
	return "minecraft:falling_block"
}

// AABB returns an empty physics.AABB so that players cannot interact with the entity.
func (t *Text) AABB() physics.AABB {
	return physics.AABB{}
}

// Immobile always returns true.
func (t *Text) Immobile() bool {
	return true
}

// NameTag returns the text passed to NewText.
func (t *Text) NameTag() string {
	return t.text
}
