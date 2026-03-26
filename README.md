# Sleeping Dragons 🐉

A dungeon-crawling adventure game where you must collect dragon eggs while avoiding awakening dragons. Built with Go and raylib-go.

## Description

Navigate through a dangerous dungeon with three dragon lairs, each guarding precious eggs. Time your movements carefully - dragons cycle between sleeping and awake states. Collect eggs when dragons sleep, but watch out! If a dragon wakes up while you're nearby, you'll lose a life and respawn at the start.

### Game Features

- **Three Difficulty Levels**: Easy, Medium, and Hard lairs with different sleep cycles
- **Dynamic Dragon Behavior**: Dragons alternate between sleeping and awake states
- **Strategic Gameplay**: Collect 20 eggs to win while managing 3 lives
- **Real-time Collision Detection**: Precise hitboxes for dragons and eggs

## How to Play

### Controls
- **Arrow Keys**: Move the hero up, down, left, and right
- **ESC**: Exit the game

### Objective
Collect 20 eggs total across all three lairs to win the game.

### Game Mechanics

**Dragon Lairs:**
- **Easy Lair (Top)**: 1 egg, dragon sleeps for 10 seconds
- **Medium Lair (Middle)**: 2 eggs, dragon sleeps for 7 seconds  
- **Hard Lair (Bottom)**: 3 eggs, dragon sleeps for 4 seconds

**Rules:**
- Dragons wake up after their sleep timer expires and stay awake for 2 seconds
- Collected eggs respawn after 2 seconds
- Getting caught by an awake dragon costs 1 life and resets your position
- Game over when you run out of lives
- Win by collecting 20 eggs total

## Installation

### Prerequisites
- Go 1.21 or higher
- raylib-go library

### Setup

1. Clone the repository:
```bash
git clone <your-repo-url>
cd sleeping-dragons
```

2. Install dependencies:
```bash
go get github.com/gen2brain/raylib-go/raylib
```

3. Ensure you have the `images/` directory with the following assets:
   - `dungeon.png` - Background
   - `hero.png` - Player character
   - `dragon-asleep.png` - Sleeping dragon sprite
   - `dragon-awake.png` - Awake dragon sprite
   - `one-egg.png` - Single egg sprite
   - `two-eggs.png` - Double egg sprite
   - `three-eggs.png` - Triple egg sprite
   - `egg-count.png` - UI egg counter icon
   - `life-count.png` - UI heart icon

4. Run the game:
```bash
go run main.go
```

## Project Structure
```
sleeping-dragons/
├── main.go           # Main game logic
├── images/           # Game assets
│   ├── dungeon.png
│   ├── hero.png
│   ├── dragon-asleep.png
│   ├── dragon-awake.png
│   ├── one-egg.png
│   ├── two-eggs.png
│   ├── three-eggs.png
│   ├── egg-count.png
│   └── life-count.png
└── README.md
```

## Code Structure

The game follows a clean entity-component architecture:

- **GameState**: Manages lives, egg collection, win/loss conditions
- **Actor**: Player character with movement and collision box
- **Enemy**: Dragon entities with awake/asleep states
- **Object**: Collectible eggs with collision detection
- **Lair**: Container for dragon, eggs, and timing logic
