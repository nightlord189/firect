package firect

import (
	"fmt"
	"time"

	"github.com/nightlord189/firect/hooks"
)

const timeFormat = "2006-01-02T15:04:05.000-0700"

//LogLevel - enum for logging level
type LogLevel int

//LogLevel - Debug/Info/Error
const (
	Debug LogLevel = iota
	Info
	Error
)

//Logger - logger struct
type Logger struct {
	OutHooks    []Hook
	Level       LogLevel
	PrintErrors bool
}

//NewLogger - constructor of Logger type
//hooks - array of hooks
//for example, hook to file, hook to console, hook to ELK
func NewLogger(hooks []Hook, level LogLevel, printErrors bool) *Logger {
	logger := Logger{
		OutHooks:    hooks,
		Level:       level,
		PrintErrors: printErrors,
	}
	return &logger
}

//Default - constructor of default logger with logging to console
func Default() *Logger {
	consoleHook := hooks.ConsoleHook{}
	logger := Logger{
		OutHooks:    []Hook{&consoleHook},
		Level:       Debug,
		PrintErrors: true,
	}
	return &logger
}

//AddHook - adds new hook to logger
func (l *Logger) AddHook(hook Hook) {
	l.OutHooks = append(l.OutHooks, hook)
}

//Debug - log debug message with preordered fields
func (l *Logger) Debug(component, action, message string, data interface{}, fields map[string]interface{}) {
	if l.Level > Debug {
		return
	}
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["component"] = component
	fields["action"] = action
	fields["message"] = message
	fields["data"] = data
	fields["time"] = time.Now().Format(timeFormat)
	fields["type"] = "Debug"
	l.LogFields(fields)
}

//DebugFields - log debug custom fields
func (l *Logger) DebugFields(fields map[string]interface{}) {
	fields["type"] = "Debug"
	l.LogFields(fields)
}

//Info - log info message with preordered fields
func (l *Logger) Info(component, action, message string, data interface{}, fields map[string]interface{}) {
	if l.Level > Info {
		return
	}
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["component"] = component
	fields["action"] = action
	fields["message"] = message
	fields["data"] = data
	fields["time"] = time.Now().Format(timeFormat)
	fields["type"] = "Info"
	l.LogFields(fields)
}

//InfoFields - log info custom fields
func (l *Logger) InfoFields(fields map[string]interface{}) {
	fields["type"] = "Info"
	l.LogFields(fields)
}

//Info - log error message with preordered fields
func (l *Logger) Error(component, action, message string, data interface{}, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["component"] = component
	fields["action"] = action
	fields["message"] = message
	fields["data"] = data
	fields["time"] = time.Now().Format(timeFormat)
	fields["type"] = "Error"
	l.LogFields(fields)
}

//ErrorFields - log error custom fields
func (l *Logger) ErrorFields(fields map[string]interface{}) {
	fields["type"] = "Error"
	l.LogFields(fields)
}

//LogFields - log fully custom fields
func (l *Logger) LogFields(fields map[string]interface{}) {
	for _, hook := range l.OutHooks {
		err := hook.Log(fields)
		if err != nil && l.PrintErrors {
			fmt.Printf("log error: %v\n", err)
		}
	}
}
