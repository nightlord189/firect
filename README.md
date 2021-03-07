# firect
Simple and extendable logger for Go

Firect is simple logger for Golang

Example:
```
package main

import "github.com/nightlord189/firect"

func main() {
	logger := firect.Default()
	logger.Debug("some-server", "main", "start", nil, nil)
}
```

Example with hooks to file and ElasticSearch:
```
package main

import (
	"github.com/nightlord189/firect"
	"github.com/nightlord189/firect/hooks"
)

func main() {
	logger := firect.Default()
	fileHook, _ := hooks.NewFileHook("log.txt")
	logger.AddHook(fileHook)
	elasticHook := hooks.NewElasticHook("http://elastic.url:9200/index/log", "elastic", "password")
	logger.AddHook(elasticHook)
	logger.OutHooks = append(logger.OutHooks, elasticHook)
	logger.Debug("some-server", "main", "start", nil, nil)
}
```

You also could implement your custom hook. Methods LogStr and Log should be implemented:
```
//Hook - interface for hooks
type Hook interface {
	LogStr(string) error
	Log(map[string]interface{}) error
}
```
