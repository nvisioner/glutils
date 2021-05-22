package gfx

import "github.com/go-gl/glfw/v3.3/glfw"

type AnimationManager struct {
	previousTime                 float64
	totalElapsed                 float64
	globalAnimationCount         int
	angle                        float64
	elapsed                      float64
	animationTimes               []float64
	animationFunctions           []func(t float32)
	animationContinuousFunctions []func()
}

func NewAnimationManager() *AnimationManager {
	return &AnimationManager{}
}

func (am *AnimationManager) AddAnimation(animation func(t float32), time float64) {
	am.animationFunctions, am.animationTimes = append(am.animationFunctions, animation), append(am.animationTimes, time)
}

func (am *AnimationManager) AddContunuousAnimation(animation func()) {
	am.animationContinuousFunctions = append(am.animationContinuousFunctions, animation)
}

func (am *AnimationManager) GetElapsed() float64 {
	return am.elapsed
}

func (am *AnimationManager) Update() {
	time := glfw.GetTime()
	am.elapsed = time - am.previousTime
	am.totalElapsed += am.elapsed
	am.previousTime = time
	am.angle += am.elapsed

	if am.globalAnimationCount < len(am.animationFunctions) {
		if animationTime := am.animationTimes[am.globalAnimationCount]; am.totalElapsed < animationTime {
			t := float32(am.totalElapsed / animationTime)
			am.animationFunctions[am.globalAnimationCount](t)
		} else {
			am.animationFunctions[am.globalAnimationCount](1)
			am.totalElapsed = 0
			am.globalAnimationCount++
		}
	}
	for _, animation := range am.animationContinuousFunctions {
		animation()
	}
}

func (am *AnimationManager) Init() {
	am.previousTime = glfw.GetTime()
}

func (am *AnimationManager) GetAngle() float64 {
	return am.angle
}
