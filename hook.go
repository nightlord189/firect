package firect

type Hook interface {
	LogStr(string) error
	Log(map[string]interface{}) error
}
