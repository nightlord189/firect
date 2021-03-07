package hooks

import "fmt"

//ConsoleHook - default hook to just print logs to console
type ConsoleHook struct {
}

//LogStr - log just string
func (h *ConsoleHook) LogStr(str string) error {
	fmt.Println(str)
	return nil
}

//Log - log map
func (h *ConsoleHook) Log(fields map[string]interface{}) error {
	fmt.Println(fields)
	return nil
}
