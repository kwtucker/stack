package src

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Stack struct {
	Reversed     bool   // caller to origin.
	FunctionName bool   // function names.
	Path         bool   // path to file
	Marker       string // example: 1, -, *, =>
	stackList    []string
	lastMarker   interface{}
}

func (s *Stack) StackIt(verbosity int) string {
	// Need to zero out from last call to StackIt
	s.stackList = []string{}
	s.lastMarker = nil

	pc := make([]uintptr, verbosity)
	n := runtime.Callers(2, pc)
	if n < 1 {
		// Don't want 0 caller ever
		return "StackIt Needs a verbosity greater than zero."
	}

	pc = pc[:n]
	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()
		if strings.Contains(frame.File, "runtime") {
			break
		}

		s.line(&frame)

		if !more {
			break
		}
	}

	return s.stackit()
}

func (s *Stack) stackit() string {
	orderedStack := []string{}
	if s.Reversed {
		orderedStack = s.reversedStack()
	}

	for n := 0; n < len(s.stackList); n++ {
		orderedStack = append(orderedStack, "  "+s.nextMarker()+s.stackList[n])
	}

	return " \n" + strings.Join(orderedStack, " \n")
}

func (s *Stack) reversedStack() []string {
	orderedStack := []string{}
	for n := len(s.stackList) - 1; n >= 0; n-- {
		orderedStack = append(orderedStack, "  "+s.nextMarker()+s.stackList[n])
	}

	return orderedStack
}

func (s *Stack) line(frame *runtime.Frame) {
	// Parse out the name of the function name.
	funcn := strings.Split(frame.Function, ".")
	functionName := funcn[len(funcn)-1]
	if strings.Contains(functionName, "func") {
		functionName = funcn[len(funcn)-2]
	}

	line := ""

	if s.FunctionName {
		line += fmt.Sprintf("%s() ", functionName)
	}

	if s.Path {
		line += fmt.Sprintf("%s:%d", frame.File, frame.Line)
	}

	s.stackList = append(s.stackList, line)
}

func (s *Stack) nextMarker() string {
	switch strings.ToLower(s.Marker) {
	case "datetime":
		return "[" + time.Now().Format(time.Stamp) + "] "
	case "int":
		m := 0
		if marker, ok := s.lastMarker.(int); ok {
			m = marker
		}
		m++
		s.lastMarker = m
		return strconv.Itoa(m) + ". "
	default:
		return s.Marker + " "
	}
}
