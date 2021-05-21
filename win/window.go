package win

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	width  int
	height int
	glfw   *glfw.Window

	inputManager  *InputManager
	firstFrame    bool
	dTime         float64
	lastFrameTime float64
}

func (w *Window) InputManager() *InputManager {
	return w.inputManager
}

func NewWindow(width, height int, title string) *Window {

	gWindow, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	gWindow.MakeContextCurrent()
	gWindow.SetInputMode(glfw.CursorMode, glfw.CursorNormal)

	im := NewInputManager()

	gWindow.SetKeyCallback(im.keyCallback)
	gWindow.SetCursorPosCallback(im.mouseCallback)

	return &Window{
		width:        width,
		height:       height,
		glfw:         gWindow,
		inputManager: im,
		firstFrame:   true,
	}
}

func (w *Window) Width() int {
	return w.width
}

func (w *Window) Height() int {
	return w.height
}

func (w *Window) ShouldClose() bool {
	return w.glfw.ShouldClose()
}

// StartFrame sets everything up to start rendering a new frame.
// This includes swapping in last rendered buffer, polling for window events,
// checkpointing cursor tracking, and updating the time since last frame.
func (w *Window) StartFrame() {
	// swap in the previous rendered buffer
	w.glfw.SwapBuffers()

	// poll for UI window events
	glfw.PollEvents()

	if w.inputManager.IsActive(PROGRAM_QUIT) {
		w.glfw.SetShouldClose(true)
	}

	// base calculations of time since last frame (basic program loop idea)
	// For better advanced impl, read: http://gafferongames.com/game-physics/fix-your-timestep/
	curFrameTime := glfw.GetTime()

	if w.firstFrame {
		w.lastFrameTime = curFrameTime
		w.firstFrame = false
	}

	w.dTime = curFrameTime - w.lastFrameTime
	w.lastFrameTime = curFrameTime

	w.inputManager.CheckpointCursorChange()
}

func (w *Window) SinceLastFrame() float64 {
	return w.dTime
}

func InitGlfw(versionMajor, versionMinor int) {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, versionMajor)
	glfw.WindowHint(glfw.ContextVersionMinor, versionMinor)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}
