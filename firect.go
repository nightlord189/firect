package firect

import "github.com/nightlord189/firect/hooks"

//Logger - logger struct
type Logger struct {
	OutHooks []Hook
}

//NewLogger - constructor of Logger type
//hooks - array of hooks
//for example, hook to file, hook to console, hook to ELK
func NewLogger(hooks []Hook) *Logger {
	logger := Logger{
		OutHooks: hooks,
	}
	return &logger
}

//Default - constructor of default logger with logging to console
func Default() *Logger {
	consoleHook := hooks.ConsoleHook{}
	logger := Logger{
		OutHooks: []Hook{&consoleHook},
	}
	return &logger
}

func (l *Logger) Debug(component, action, message string, data interface{}, fields map[string]interface{}) {
	fields["component"] = component
	fields["action"] = action
	fields["data"] = data
	fields["type"] = "Debug"
	l.LogFields(fields)
}

func (l *Logger) DebugFields(fields map[string]interface{}) {
	fields["type"] = "Debug"
	l.LogFields(fields)
}

func (l *Logger) Info(component, action, message string, data interface{}, fields map[string]interface{}) {
	fields["component"] = component
	fields["action"] = action
	fields["data"] = data
	fields["type"] = "Info"
	l.LogFields(fields)
}

func (l *Logger) InfoFields(fields map[string]interface{}) {
	fields["type"] = "Info"
	l.LogFields(fields)
}

func (l *Logger) Error(component, action, message string, data interface{}, fields map[string]interface{}) {
	fields["component"] = component
	fields["action"] = action
	fields["data"] = data
	fields["type"] = "Error"
	l.LogFields(fields)
}

func (l *Logger) ErrorFields(fields map[string]interface{}) {
	fields["type"] = "Error"
	l.LogFields(fields)
}

func (l *Logger) LogFields(fields map[string]interface{}) {
	for _, hook := range l.OutHooks {
		hook.Log(fields)
	}
}
