package firect

//Hook - interface for hooks
type Hook interface {
	LogStr(string) error
	Log(map[string]interface{}) error
}
