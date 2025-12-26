//Author: Design-BAB
//Date: 12/24/2025
//Description: Time for dungeon crawling! The goal is to reach 268 lines of code
//Notes: Continue onward from suggestions on pg191

package main

import (
	"math"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Width                 = 800
	Height                = 600
	CenterX               = Width / 2
	CenterY               = Height / 2
	EggTarget             = 20
	AttackDistance        = 125
	DragonAwakeAndEggTime = 2
	MoveDistance          = 5
	MaxFrames             = 432000
	HeroOriginalX         = 200
	HeroOriginalY         = 300
)

type GameState struct {
	Lives         int
	EggsCollected int
	IsOver        bool
	IsComplete    bool
	ScheduleLair  time.Time
}

func newGame() *GameState {
	now := time.Now()
	return &GameState{Lives: 3, ScheduleLair: now}
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
	IsAwake      bool
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

func update(hero *Actor, lairs [3]*Lair, DragonTexture map[string]rl.Texture2D, yourGame *GameState) {
	if rl.IsKeyDown(rl.KeyRight) {
		hero.X = hero.X + hero.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		hero.X = hero.X - hero.Speed
	}
	if rl.IsKeyDown(rl.KeyUp) {
		hero.Y = hero.Y - hero.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		hero.Y = hero.Y + hero.Speed
	}
	//collision with the window
	hero.X = rl.Clamp(hero.X, 0.0, Width-hero.Width)
	hero.Y = rl.Clamp(hero.Y, 0.0, Height-hero.Height)
	checkForCollisions(lairs, hero, yourGame)
	//the scheduler
	if time.Since(yourGame.ScheduleLair) >= 1*time.Second {
		updateLairs(lairs, DragonTexture, yourGame)
		now := time.Now()
		yourGame.ScheduleLair = now
	}
	if yourGame.Lives <= 0 {
		yourGame.IsOver = true
	}
}

func updateLairs(lairs [3]*Lair, DragonTexture map[string]rl.Texture2D, yourGame *GameState) {
	for _, lair := range lairs {
		if lair.Dragon.IsAwake == false {
			updateSleepingDragon(lair, DragonTexture)
		} else if lair.Dragon.IsAwake == true {
			updateWakingDragon(lair, DragonTexture)
		}
		updateEgg(lair)
	}
}

func updateSleepingDragon(lair *Lair, DragonTexture map[string]rl.Texture2D) {
	if lair.SleepCount >= lair.SleepLength {
		lair.Dragon.Texture = DragonTexture["awake"]
		lair.Dragon.IsAwake = true
		lair.SleepCount = 0
	} else {
		lair.SleepCount++
	}
}

func updateWakingDragon(lair *Lair, DragonTexture map[string]rl.Texture2D) {
	if lair.WakeCount >= DragonAwakeAndEggTime {
		lair.Dragon.Texture = DragonTexture["sleeping"]
		lair.Dragon.IsAwake = false
		lair.WakeCount = 0
	} else {
		lair.WakeCount++
	}
}

func updateEgg(lair *Lair) {
	if lair.HideEgg == true {
		if lair.HideEggCount >= DragonAwakeAndEggTime {
			lair.HideEgg = false
			lair.HideEggCount = 0
		} else {
			lair.HideEggCount++
		}
	}
}

func checkForCollisions(lairs [3]*Lair, hero *Actor, yourGame *GameState) {
	for _, lair := range lairs {
		if lair.HideEgg == false {
			checkForEggCollision(lair, hero, yourGame)
		}
		if lair.Dragon.IsAwake {
			checkForDragonCollision(lair, hero, yourGame)
		}
	}
}

func checkForEggCollision(lair *Lair, hero *Actor, yourGame *GameState) {
	//collision between egg and hero
	if rl.CheckCollisionRecs(hero.Rectangle, lair.Egg.Rectangle) {
		lair.HideEgg = true
		yourGame.EggsCollected += lair.EggCount
		if yourGame.EggsCollected >= EggTarget {
			yourGame.IsComplete = true
		}
	}
}

func checkForDragonCollision(lair *Lair, hero *Actor, yourGame *GameState) {
	xDistance := hero.X - lair.Dragon.X
	yDistance := hero.Y - lair.Dragon.Y
	distance := math.Hypot(float64(xDistance), float64(yDistance))
	if distance < AttackDistance {
		handleDragonCollision(hero, yourGame)
	}
}

func handleDragonCollision(hero *Actor, yourGame *GameState) {
	// the original code uses animate(hero, pos_HERO_START, on_finished=subtract_life)... I deal with that later
	hero.X = HeroOriginalX
	hero.Y = HeroOriginalY
	yourGame.Lives--
}

func draw(lairs [3]*Lair, background, heart, egg rl.Texture2D, hero *Actor, yourGame *GameState) {
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
		drawGUI(heart, egg, yourGame)
	}
	rl.EndDrawing()
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

// In the book, it calls it draw_counters(eggs_collected, live/s)
func drawGUI(heart, egg rl.Texture2D, yourGame *GameState) {
	rl.DrawTexture(egg, 0, Height-30, rl.White)
	rl.DrawText(strconv.Itoa(yourGame.EggsCollected), 30, Height-30, 40, rl.RayWhite)
	rl.DrawTexture(heart, 60, Height-30, rl.White)
	rl.DrawText(strconv.Itoa(yourGame.Lives), 90, Height-30, 40, rl.RayWhite)
}

func main() {
	rl.InitWindow(Width, Height, "Sleeping Dragons")
	defer rl.CloseWindow()
	yourGame := newGame()
	rl.SetTargetFPS(60)
	//import Textures
	background := rl.LoadTexture("images/dungeon.png")
	defer rl.UnloadTexture(background)
	heart := rl.LoadTexture("images/life-count.png")
	defer rl.UnloadTexture(heart)
	heroTexture := rl.LoadTexture("images/hero.png")
	defer rl.UnloadTexture(heroTexture)
	DragonTexture := map[string]rl.Texture2D{
		"sleeping": rl.LoadTexture("images/dragon-asleep.png"),
		"awake":    rl.LoadTexture("images/dragon-awake.png"),
	}
	for _, texture := range DragonTexture {
		defer rl.UnloadTexture(texture)
	}
	//Remember, 0 is reserved for the GUI, and the rest of the numbers represent the number of eggs
	var eggTexture [4]rl.Texture2D
	eggTexture[0] = rl.LoadTexture("images/egg-count.png")
	eggTexture[1] = rl.LoadTexture("images/one-egg.png")
	eggTexture[2] = rl.LoadTexture("images/two-eggs.png")
	eggTexture[3] = rl.LoadTexture("images/three-eggs.png")
	for _, texture := range eggTexture {
		defer rl.UnloadTexture(texture)
	}
	//Our pieces in the game!
	hero := newActor(heroTexture, 200, 300)
	//from pg184- making lairs. I had to put it down here because I can't make lairs without importing textures first
	var lairs [3]*Lair
	lairs[0] = newLair(newEnemy(DragonTexture["sleeping"], 500, 70), newObject(eggTexture[1], 420, 140), 1, 10, 0)
	lairs[1] = newLair(newEnemy(DragonTexture["sleeping"], 500, 230), newObject(eggTexture[2], 420, 290), 2, 7, 1)
	lairs[2] = newLair(newEnemy(DragonTexture["sleeping"], 500, 400), newObject(eggTexture[3], 420, 470), 3, 4, 2)
	frames := 0
	for !rl.WindowShouldClose() && frames < MaxFrames {
		update(hero, lairs, DragonTexture, yourGame)
		draw(lairs, background, heart, eggTexture[0], hero, yourGame)
		frames++
	}
}
