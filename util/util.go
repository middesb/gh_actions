package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gh_actions/common"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "    ")
	if err != nil {
		return in
	}
	return out.String()
}
func Log(vals ...interface{}) {
	for _, v := range vals {
		log.Println(v)
	}
}

func FetchObjectFromBody(resp *http.Response, v any) string {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return "Error while reading the response bytes:" + err.Error()
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(body, v)
	if err != nil {
		return "Error while Unmarshalling the response bytes:" + err.Error()
	}
	return ""
}

func HealthCheck(w http.ResponseWriter) {
	fmt.Fprintf(w, "Healthy..")
}

func BuildUrl(host string, repoPath string, tailUrl string) string {
	host = strings.TrimSuffix(host, "/")
	repoPath = strings.TrimPrefix(repoPath, "/")
	repoPath = strings.TrimSuffix(repoPath, "/")
	tailUrl = strings.TrimPrefix(tailUrl, "/")
	tailUrl = strings.TrimSuffix(tailUrl, "/")

	return (host + "/" + repoPath + "/" + tailUrl)
}

func LogJson(v any) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(JsonPrettyPrint(string(b)))
}

func WriteResult(v any) {
	path := "/app/ghactions-data/data.txt"
	b, err := json.Marshal(v)
	if err != nil {
		Log("WriteFile: err", err)
		return
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		fmt.Println("WriteFile: err", err)
		return
	}

}
func ReadResult(v any) {
	path := "/app/ghactions-data/data.txt"
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile: err", err)
		v = nil
		return
	}
	err = json.Unmarshal(dat, v)
	if err != nil {
		Log("Error while Unmarshalling the response bytes:" + err.Error())
	}
}
func BuildRequest() common.GHRequest {
	args := os.Args
	var ghRequest common.GHRequest
	ghRequest.Action = args[1]
	ghRequest.Host = args[2]
	ghRequest.Token = args[3]
	ghRequest.RepoPath = args[4]
	ghRequest.ID = args[5]
	return ghRequest
}

func InitStorage() {
	path := "/app/ghactions-data"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		Log(err)
		Log("Creating data folder")
		err = os.MkdirAll(path, 0700) // Create your file
		if err != nil {
			Log(err)
		}
	}
}
