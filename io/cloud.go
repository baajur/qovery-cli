package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CloudProviders struct {
	Results []CloudProvider `json:"results"`
}

type CloudProvider struct {
	ObjectType string                `json:"object_type"`
	Id         string                `json:"id"`
	Name       string                `json:"name"`
	Regions    []CloudProviderRegion `json:"regions"`
}

type CloudProviderRegion struct {
	ObjectType  string `json:"object_type"`
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
}

func ListCloudProviders() CloudProviders {
	c := CloudProviders{}

	CheckAuthenticationOrQuitWithMessage()

	req, _ := http.NewRequest(http.MethodGet, RootURL+"/cloud", nil)
	req.Header.Set(headerAuthorization, headerValueBearer+GetAuthorizationToken())

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return c
	}

	err = CheckHTTPResponse(resp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &c)

	return c
}
