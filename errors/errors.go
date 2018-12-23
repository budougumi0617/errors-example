package errors

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// NewByPkg returns error by pkg/errors.
func NewByPkg(s string) error {
	return errors.New(s)
}

// New returns MyError
func New(s string) error {
	return MyError{
		cause:  s,
		frames: NewFrame(callers()),
	}
}

// MyError is cutom error.
type MyError struct {
	cause  string
	frames []Frame
}

// Error はerrorインターフェースの実装
func (me MyError) Error() string {
	return fmt.Sprintf("Cause: %s\nTrace:\n%s", me.cause, me.traces())
}

// おおざっぱなトレース出力
func (me MyError) traces() string {
	var buf bytes.Buffer
	for _, fr := range me.frames {
		fmt.Fprintf(&buf, "%s\n", fr)
	}
	return buf.String()
}

// Frame はスタックフレームを表現する独自構造体
type Frame struct {
	// File はそのスタックが発生した開始されたファイル
	File string
	// LineNumber はそのスタックが開始されたファイル行数
	LineNumber int
	// Name はそのスタックが開始された関数・メソッド名
	Name string
	// ProgramCounter は元データ
	ProgramCounter uintptr
}

// https://github.com/pkg/errors/blob/816c9085562cd7ee03e7f8188a1cfd942858cded/stack.go#L133
func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[0 : n-2]
}

// NewFrame はスタックトレースの各スタックフレームをFrameに格納する
func NewFrame(pcs []uintptr) []Frame {
	frames := []Frame{}

	for _, pc := range pcs {
		frame := Frame{ProgramCounter: pc}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			return frames
		}
		frame.Name = trimPkgName(fn)

		frame.File, frame.LineNumber = fn.FileLine(pc - 1)
		frames = append(frames, frame)
	}
	return frames
}

// 大雑把なフレーム出力
func (frame Frame) String() string {
	return fmt.Sprintf("%s:%d %s", frame.File, frame.LineNumber, frame.Name)
}

// package名を取り除く
func trimPkgName(fn *runtime.Func) string {
	name := fn.Name()
	if ld := strings.LastIndex(name, "."); ld >= 0 {
		name = name[ld+1:]
	}
	return name
}
