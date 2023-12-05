package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Timer struct {
	startTime float64 // (seconds)
	lifeTime  float64 // (seconds)
	start     bool
	ifDone    bool
}

func NewTimer(life float64) *Timer {
	return &Timer{
		startTime: 0,
		lifeTime:  life,
		start:     false,
		ifDone:    false,
	}
}

func (t *Timer) Start() {
	t.start = true
	t.startTime = rl.GetTime()
}
func (t *Timer) IsDone() bool {
	if rl.GetTime()-t.startTime >= t.lifeTime && !t.ifDone {
		t.ifDone = true
		return true
	} else {
		return false
	}
}
func (t *Timer) Elapsed() float64 {
	return rl.GetTime() - t.startTime
}

func (t *Timer) Reinitiate() {
	t.start = true
	t.startTime = rl.GetTime()
	t.ifDone = false
}
