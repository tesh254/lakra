package engine

/*
* Defines particular constants in the game
 */

// Direction constants
const (
	// The values are determined via a catersian plane
	DirLeft       = -1 // x-axis
	DirRight      = 1  // x-axis
	DirStationary = 0  // x-axis
)

// Event constants
const (
	EventFloorCollision = 0
	EventDropOffLevel   = 1
	EventFreeFall       = 2
)

// Collision Edges
const (
	EdgeTop         = "top"
	EdgeTopLeft     = "top_left"
	EdgeTopRight    = "top_right"
	EdgeBottom      = "bottom"
	EdgeBottomLeft  = "bottom_left"
	EdgeBottomRight = "bottom_right"
	EdgeLeft        = "left"
	EdgeRight       = "right"
	EdgeNone        = "none"
)
