package io

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strconv"
)

type Logs struct {
	Results []Log `json:"results"`
}

type Log struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Message   string `json:"message"`
}

func ListApplicationLogs(lastLines int, follow bool, projectId string, environmentId string, applicationId string) {
	if projectId == "" || environmentId == "" || applicationId == "" {
		return
	}

	CheckAuthenticationOrQuitWithMessage()

	url := RootURL + "/project/" + projectId + "/environment/" + environmentId + "/application/" + applicationId + "/log"

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	q := req.URL.Query()
	q.Add("tail", strconv.Itoa(lastLines))
	q.Add("follow", strconv.FormatBool(follow))
	req.URL.RawQuery = q.Encode()

	req.Header.Set(headerAuthorization, headerValueBearer+GetAuthorizationToken())
	req.Header.Set("accept", "application/stream+json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return
	}

	reader := bufio.NewReader(resp.Body)

	for {
		bytes, _ := reader.ReadBytes('\n')
		if len(bytes) > 0 {
			var log Log
			_ = json.Unmarshal(bytes, &log)
			print(log.Message)
		} else if !follow {
			return
		}
	}
}
