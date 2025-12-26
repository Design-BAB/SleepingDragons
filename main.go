//Author: Design-BAB
//Date: 12/24/2025
//Description: Time for dungeon crawling! The goal is to reach 268 lines of code

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Width            = 800
	Height           = 600
	CenterX          = Width / 2
	CenterY          = Height / 2
	EggTarget        = 20
	AttackDistance   = 200
	DragonAndEggTime = 2
	MoveDistance     = 5
	MaxFrames        = 432000
)

type GameState struct {
	Lives         int
	EggsCollected int
	IsOver        bool
	IsComplete    bool
	ResetRequired bool
}

func newGame() *GameState {
	return &GameState{Lives: 3}
}

type Actor struct {
	Texture rl.Texture2D
	//this is the collision box``
	rl.Rectangle // This gives Actor all the fields of rl.Rectangle (X, Y, Width, Height)
	Speed        float32
}

func newActor(texture rl.Texture2D, x, y float32) *Actor {
	return &Actor{Texture: texture, Rectangle: rl.Rectangle{X: x, Y: y, Width: float32(texture.Width), Height: float32(texture.Height)}, Speed: MoveDistance}
}

// as this game becomes more complex, I probably add more variables to enemy
type Enemy struct {
	Texture rl.Texture2D
	//this is the collision box``
	rl.Rectangle // This gives Actor all the fields of rl.Rectangle (X, Y, Width, Height)
}

func newEnemy(texture rl.Texture2D, x, y float32) *Enemy {
	return &Enemy{Texture: texture, Rectangle: rl.Rectangle{X: x, Y: y, Width: float32(texture.Width), Height: float32(texture.Height)}}
}

type Object struct {
	Texture rl.Texture2D
	//this is the collision box``
	rl.Rectangle // This gives Actor all the fields of rl.Rectangle (X, Y, Width, Height)
}

func newObject(texture rl.Texture2D, x, y float32) *Object {
	return &Object{Texture: texture, Rectangle: rl.Rectangle{X: x, Y: y, Width: float32(texture.Width), Height: float32(texture.Height)}}
}

type Lair struct {
	Dragon       *Enemy
	Egg          *Object
	EggCount     int
	HideEgg      bool
	HideEggCount int
	SleepLength  int
	SleepCount   int
	WakeCount    int
	Difficulty   int
}

func newLair(dragon *Enemy, egg *Object, eggCount, sleepLength, difficulty int) *Lair {
	return &Lair{Dragon: dragon, Egg: egg, EggCount: eggCount, SleepLength: sleepLength, Difficulty: difficulty}
}

// fromPg186
func drawLair(lairs [3]*Lair) {
	for _, lair := range lairs {
		rl.DrawTexture(lair.Dragon.Texture, int32(lair.Dragon.X), int32(lair.Dragon.Y), rl.White)
		if lair.HideEgg == false {
			rl.DrawTexture(lair.Egg.Texture, int32(lair.Egg.X), int32(lair.Egg.Y), rl.White)
		}
	}
}

func drawCounter() {

}

func draw(lairs [3]*Lair, background rl.Texture2D, hero *Actor, dragon *Enemy, yourGame *GameState) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(background, 0, 0, rl.White)
	if yourGame.IsOver {
		msg1 := "Game over!"
		w1 := rl.MeasureText(msg1, 40)
		rl.DrawText(msg1, CenterX-(w1/2), CenterY, 40, rl.Red)
		msg3 := "Press ESC to exit"
		w3 := rl.MeasureText(msg3, 20)
		rl.DrawText(msg3, CenterX-(w3/2), CenterY+100, 20, rl.DarkGray)
	} else if yourGame.IsComplete {
		msg1 := "You won!"
		w1 := rl.MeasureText(msg1, 40)
		rl.DrawText(msg1, CenterX-(w1/2), CenterY, 40, rl.Green)
		msg3 := "Press ESC to exit"
		w3 := rl.MeasureText(msg3, 20)
		rl.DrawText(msg3, CenterX-(w3/2), CenterY+100, 20, rl.DarkGray)

	} else {
		//this is what is drawn during gameplay
		rl.DrawTexture(hero.Texture, int32(hero.X), int32(hero.Y), rl.White)
		drawLair(lairs)
	}
	rl.EndDrawing()
}

func main() {
	rl.InitWindow(Width, Height, "Sleeping Dragons")
	defer rl.CloseWindow()
	yourGame := newGame()
	rl.SetTargetFPS(60)
	//import Textures
	background := rl.LoadTexture("images/dungeon.png")
	defer rl.UnloadTexture(background)
	heroTexture := rl.LoadTexture("images/hero.png")
	defer rl.UnloadTexture(heroTexture)
	DragonTexture := map[string]rl.Texture2D{
		"sleeping": rl.LoadTexture("images/dragon-asleep.png"),
		"awake":    rl.LoadTexture("images/dragon-awake.png"),
	}
	for _, texture := range DragonTexture {
		defer rl.UnloadTexture(texture)
	}
	//Why the long name? To remind myself that 0 is 1 egg
	var eggPlusOneTexture [3]rl.Texture2D
	eggPlusOneTexture[0] = rl.LoadTexture("images/one-egg.png")
	eggPlusOneTexture[1] = rl.LoadTexture("images/two-eggs.png")
	eggPlusOneTexture[2] = rl.LoadTexture("images/three-eggs.png")
	for _, texture := range eggPlusOneTexture {
		defer rl.UnloadTexture(texture)
	}
	//Our pieces in the game!
	hero := newActor(heroTexture, 200, 300)
	dragon := newEnemy(DragonTexture["sleeping"], 600, 100)
	//from pg184- making lairs. I had to put it down here because I can't make lairs without importing textures first
	var lairs [3]*Lair
	lairs[0] = newLair(newEnemy(DragonTexture["sleeping"], 600, 100), newObject(eggPlusOneTexture[0], 400, 100), 1, 10, 0)
	lairs[1] = newLair(newEnemy(DragonTexture["sleeping"], 600, 300), newObject(eggPlusOneTexture[1], 400, 300), 2, 7, 1)
	lairs[2] = newLair(newEnemy(DragonTexture["sleeping"], 600, 500), newObject(eggPlusOneTexture[2], 400, 500), 3, 4, 2)
	frames := 0
	for !rl.WindowShouldClose() && frames < MaxFrames {
		draw(lairs, background, hero, dragon, yourGame)
	}
}
