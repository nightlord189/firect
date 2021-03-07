package hooks

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//ElasticHook - hook for logging to ElasticSearch directly via API
//without Logstash or Filebeat
type ElasticHook struct {
	URL      string
	User     string
	Password string
	Client   http.Client
}

//NewElasticHook - constructor
//url - url to log, example: http://elastic.co:9200/your-index/log
//user - user for basic auth (if empty than auth will not be used)
//password - password for basic auth
func NewElasticHook(url, user, password string) *ElasticHook {
	hook := ElasticHook{
		URL:      url,
		User:     user,
		Password: password,
		Client:   http.Client{},
	}
	return &hook
}

func (h *ElasticHook) getBase64Auth() string {
	auth := h.User + ":" + h.Password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//LogStr - log just string
func (h *ElasticHook) LogStr(str string) error {
	fields := map[string]interface{}{"data": str}
	return h.Log(fields)
}

//Log - log map
func (h *ElasticHook) Log(fields map[string]interface{}) error {
	byteData, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", h.URL, bytes.NewBuffer(byteData))
	if err != nil {
		return err
	}
	if h.User != "" {
		fmt.Println("authorizing")
		fmt.Println(req.Method)
		req.Header.Set("Authorization", "Basic "+h.getBase64Auth())
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("bad response, %d, decode resp: %v", resp.StatusCode, err)
		}
		return fmt.Errorf("bad response, %d, %s", resp.StatusCode, string(respBytes))
	}
	return nil
}
