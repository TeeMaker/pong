package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var WIDTH int32 = 600
var HEIGHT int32 = 600
var labda ball = ball{}
var scoreP int32 = 0
var scoreE int32 = 0
var speed int32

type ball struct {
	X   int32
	Y   int32
	vX  int32
	vY  int32
	box *sdl.Rect
}

type paddle struct {
	X   int32
	Y   int32
	vY  int32
	box *sdl.Rect
}

func labdaUpdate(this *ball) {
	this.X += this.vX
	this.Y += this.vY
	this.box.X = this.X
	this.box.Y = this.Y
}

func paddleUpdate(this *paddle) {
	if this.vY > 20 {
		this.vY = 10
	}
	this.Y += this.vY * speed
	if this.Y > HEIGHT-this.box.H {
		this.Y = HEIGHT - this.box.H
	}
	if this.Y < 0 {
		this.Y = 0
	}
	this.box.X = this.X
	this.box.Y = this.Y
}

func collision(b ball, p paddle) bool {
	if (b.X < p.X+p.box.W && b.X+b.box.W > p.X) && b.Y < p.Y+p.box.H && b.Y+b.box.H > p.Y {
		return true
	}

	return false
}

func run() (err error) {
	var window *sdl.Window

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, 0, sdl.GL_ACCELERATED_VISUAL)
	if err != nil {
		return err
	}
	defer renderer.Destroy()

	zero, _ := sdl.LoadBMP("./numbers/0.bmp")
	zeroTex, _ := renderer.CreateTextureFromSurface(zero)
	one, _ := sdl.LoadBMP("./numbers/1.bmp")
	oneTex, _ := renderer.CreateTextureFromSurface(one)
	two, _ := sdl.LoadBMP("./numbers/2.bmp")
	twoTex, _ := renderer.CreateTextureFromSurface(two)
	three, _ := sdl.LoadBMP("./numbers/3.bmp")
	threeTex, _ := renderer.CreateTextureFromSurface(three)
	four, _ := sdl.LoadBMP("./numbers/4.bmp")
	fourTex, _ := renderer.CreateTextureFromSurface(four)
	five, _ := sdl.LoadBMP("./numbers/5.bmp")
	fiveTex, _ := renderer.CreateTextureFromSurface(five)
	six, _ := sdl.LoadBMP("./numbers/6.bmp")
	sixTex, _ := renderer.CreateTextureFromSurface(six)
	seven, _ := sdl.LoadBMP("./numbers/7.bmp")
	sevenTex, _ := renderer.CreateTextureFromSurface(seven)
	eight, _ := sdl.LoadBMP("./numbers/8.bmp")
	eightTex, _ := renderer.CreateTextureFromSurface(eight)
	nine, _ := sdl.LoadBMP("./numbers/9.bmp")
	nineTex, _ := renderer.CreateTextureFromSurface(nine)
	var numbers []*sdl.Texture = []*sdl.Texture{zeroTex, oneTex, twoTex, threeTex, fourTex, fiveTex, sixTex, sevenTex, eightTex, nineTex}

	var labdabox sdl.Rect = sdl.Rect{X: 0, Y: 0, W: WIDTH / 20, H: WIDTH / 20}
	labda = ball{X: WIDTH / 2, Y: HEIGHT / 2, vX: 7 * speed, vY: 0, box: &labdabox}
	var player paddle = paddle{X: 0 + WIDTH/20, Y: HEIGHT / 2, box: &sdl.Rect{X: 0, Y: 0, W: WIDTH / 20, H: HEIGHT / 5}}
	var enemy paddle = paddle{X: WIDTH - 2*WIDTH/20, Y: HEIGHT / 2, box: &sdl.Rect{X: 0, Y: 0, W: WIDTH / 20, H: HEIGHT / 5}}
	var pI int = 0
	var eI int = 0
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED {
					switch t.Keysym.Scancode {
					case sdl.SCANCODE_DOWN:
						player.vY = 5
						break
					case sdl.SCANCODE_UP:
						player.vY = -5
						break
					}
				} else if t.State == sdl.RELEASED {
					switch t.Keysym.Scancode {
					case sdl.SCANCODE_DOWN:
						player.vY = 0
						break
					case sdl.SCANCODE_UP:
						player.vY = 0
						break
					}
				}
			}
		}
		if pI > 0 {
			pI--
		}
		if eI > 0 {
			eI--
		}
		labdaUpdate(&labda)
		paddleUpdate(&player)
		paddleUpdate(&enemy)
		if labda.X > WIDTH/4 {
			if enemy.Y+enemy.box.H/2 < labda.Y+rand.Int31n(50)-rand.Int31n(50) {
				enemy.vY += 1
			}
			if enemy.Y+enemy.box.H/2 > labda.Y+rand.Int31n(50)-rand.Int31n(50) {
				enemy.vY -= 1
			}
		} else if enemy.vY > 0 {
			enemy.vY -= 1
		} else if enemy.vY < 0 {
			enemy.vY += 1
		}

		if collision(labda, player) && pI == 0 {
			labda.vY = -1 * (labda.vY + player.vY/2)
			labda.vX = -1 * labda.vX
			pI = int(100 / speed)
			// labda.X = player.X + player.box.W + labda.box.W
		}
		if collision(labda, enemy) && eI == 0 {
			labda.vY = -1 * (labda.vY + enemy.vY/2)
			labda.vX = -1 * labda.vX
			eI = int(100 / speed)
			// labda.X = enemy.X - labda.box.W
		}
		if labda.Y >= HEIGHT || labda.Y <= 0 {
			labda.vY = -1 * labda.vY
		}
		if labda.X > WIDTH {
			labda = ball{X: WIDTH / 2, Y: HEIGHT / 2, vX: -7 * speed, vY: 0, box: &labdabox}
			enemy.Y = HEIGHT / 2
			enemy.vY = 0
			player.Y = HEIGHT / 2
			scoreP += 1
		}
		if labda.X < 0 {
			labda = ball{X: WIDTH / 2, Y: HEIGHT / 2, vX: 7 * speed, vY: 0, box: &labdabox}
			enemy.Y = HEIGHT / 2
			player.Y = HEIGHT / 2
			enemy.vY = 0
			scoreE += 1
		}
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.FillRect(player.box)
		renderer.FillRect(labda.box)
		renderer.FillRect(enemy.box)
		var H int32 = HEIGHT / 30
		var Y int32 = H
		var W int32 = WIDTH / 35
		for Y <= HEIGHT+H {
			renderer.FillRect(&sdl.Rect{X: WIDTH/2 - W/2, Y: Y, W: W, H: H / 2})
			Y += H
		}
		scorePstring := strconv.Itoa(int(scoreP))
		for i := len(scorePstring) - 1; i >= 0; i-- {
			currentNumber, _ := strconv.Atoi(string(scorePstring[i]))
			texture := numbers[currentNumber]
			renderer.Copy(texture, &sdl.Rect{X: 0, Y: 0, W: 16, H: 16}, &sdl.Rect{X: WIDTH/2 - int32(len(scorePstring)-i+1)*WIDTH/10 - WIDTH/15, Y: 70, W: WIDTH / 15, H: WIDTH / 15})
		}
		scoreEstring := strconv.Itoa(int(scoreE))
		for i := 0; i <= len(scoreEstring)-1; i++ {
			currentNumber, _ := strconv.Atoi(string(scoreEstring[i]))
			texture := numbers[currentNumber]
			renderer.Copy(texture, &sdl.Rect{X: 0, Y: 0, W: 16, H: 16}, &sdl.Rect{X: WIDTH/2 + int32(len(scorePstring)+i+1)*WIDTH/10, Y: 70, W: WIDTH / 15, H: WIDTH / 15})
		}

		renderer.Present()
	}

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 0 {
		speedI, _ := strconv.Atoi(args[0])
		speed = int32(speedI)
	} else {
		speed = 1
	}
	fmt.Println(speed)
	if err := run(); err != nil {
		os.Exit(1)
	}
}
