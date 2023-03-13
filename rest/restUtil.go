package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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