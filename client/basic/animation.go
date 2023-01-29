package basic

import (
	"fmt"
	"time"
)

var (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func AnimationStep(text string) {
	fmt.Print(text)
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\n\033[1A\033[K")
}
func loadingAnimation() {
	AnimationStep("(/) Loading .")
	AnimationStep("(|) Loading ..")
	AnimationStep("(\\) Loading ...")
	AnimationStep("(-) Loading ....")
}
func FakeWaiting(iteration int) {
	i := 0
	for i < iteration {
		loadingAnimation()
		i++
	}
}

func ClassAnimation(class ClassImages) {
	for i := 0; i < 5; i++ {
		AnimationStep(class.IMG1)
		AnimationStep(class.IMG2)
		AnimationStep(class.IMG3)
		AnimationStep(class.IMG4)
		AnimationStep(class.IMG5)
	}
}
