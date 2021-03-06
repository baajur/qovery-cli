package io

import (
	"fmt"
	"net/http"
	"os"
)

func Deploy(projectId string, environmentId string, applicationId string, commitId string) {
	if projectId == "" || environmentId == "" || applicationId == "" || commitId == "" {
		return
	}

	CheckAuthenticationOrQuitWithMessage()

	req, _ := http.NewRequest(http.MethodPost, RootURL+"/project/"+projectId+"/environment/"+environmentId+"/application/"+applicationId+"/commit/"+commitId+"/deploy", nil)
	req.Header.Set(headerAuthorization, headerValueBearer+GetAuthorizationToken())

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	if resp != nil && resp.StatusCode == http.StatusBadRequest {
		fmt.Println("Could not deploy application with commit " + commitId)
		fmt.Println("Are you sure you entered a valid commit sha?")
		os.Exit(1)
	}

	err = CheckHTTPResponse(resp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
