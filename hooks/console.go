package hooks

import "fmt"

type ConsoleHook struct {
}

func (h *ConsoleHook) LogStr(str string) error {
	fmt.Println(str)
	return nil
}

func (h *ConsoleHook) Log(fields map[string]interface{}) error {
	fmt.Println(fields)
	return nil
}
