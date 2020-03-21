package logging

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Infof("infof test %s", "do what you want to do")
	Warnf("warnf test %s", "do what you want to do")
	Debugf("debugf test %s", "do what you want to do")
	Errorf("errorf test %s", "do what you want to do")
	// Fatalf("fatalf test %s", "do what you want to do") //  fatal exit it
	// Panicf("panicf test %s", "do what you want to do") //  fatal exit it

	Info("info test ", "do what you want to do")
	Warn("warn test ", "do what you want to do")
	Debug("debug test ", "do what you want to do")
	Error("error test ", "do what you want to do")
	// Fatal("fatal test ", "do what you want to do") // fatal exit it
	// Panic("fatalf test ", "do what you want to do") //  fatal exit it

}
