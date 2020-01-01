# Lakra
Simple game engine built with Golang

## Preview

![game preview]('./preview.png')

## Game Engine Info

### File by Info

* `game.go`:
	* Gets the current level
	* Gets the input broadcasted
* `game_object.go`:
	* Gets the current sprite of the games' state
	* Gets the current sprite frame based on ticker
	* Gets dimensions of the current game object
	* Handles position of game object
* `level.go`:
	* Handles transitioning to levels after completions
* `sprite.go`:
	* Handles creation of a single sprite and adding it to an image canvas
* `sprite_group.go`:
	* Handles creation of sprite group and adding them to image canvas
* `window.go`:
	* Creates the game window and renders image being shown

## Build

The `main.go` file builds to a game written in the file

```bash
go build examples/main.go
```

## Run

After bulding, you should be able to run the game:
* Windows

Click on the exe file built

* Unix (Linux/Mac OS)

```bash
./main
```

## Roadmap

* [x] Create basic game engine
* [ ] Create GUI for sprite creation
* [ ] Create GUI for defining sprite control
