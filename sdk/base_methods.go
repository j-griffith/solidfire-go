package sdk

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

//Connect just really sets based info and then does an API call to see if it works
//its all stateless after the initial call
func (sfClient *SFClient) Connect(ctx context.Context, host string, version string, uid string, password string) *SdkError {
	sfClient.baseUrl = fmt.Sprintf("https://%s/json-rpc/%s", host, version)
	var entry BaseRequest

	entry.Id = 1
	entry.Method = "GetAPI"
	entry.Parameters = nil
	var res GetAPIResult
	sfClient.userId = uid
	sfClient.password = password
	_, sdkError := sfClient.MakeSFCall(ctx, "GetAPI", 1, nil, &res)
	return sdkError
}

func (sfClient *SFClient) MakeSFCall(ctx context.Context, method string, id int32, params interface{}, res interface{}) (BaseResponse, *SdkError) {
	log.WithContext(ctx).Debugf("Starting call %s", method)
	var entry BaseRequest
	var returnError *SdkError = nil

	client := mkHttpClient()
	entry.Id = id
	entry.Method = method
	entry.Parameters = params

	bits, _ := json.Marshal(entry)
	//inputStr := string(bits)
	req, reqerr := http.NewRequest("POST", sfClient.baseUrl, bytes.NewReader(bits))
	if reqerr != nil {
		fmt.Println(reqerr.Error())
		os.Exit(1)
	}

	req.Header.Add("Authorization", "Basic "+makeBasicAuthHeader(sfClient.userId, sfClient.password))
	req.Header.Add("Content", "Application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error " + err.Error())
		var errorData SdkError
		errorData.Code = fmt.Sprintf("%s.unkown", SfapiError)
		errorData.Detail = err.Error()
		returnError = &errorData
	} else if resp.StatusCode != 200 {
		log.Printf("Status code %d\n", resp.StatusCode)
		var errorData SdkError
		errorData.Code = fmt.Sprintf("%s.%d", NetworkError, resp.StatusCode)
		errorData.Detail = resp.Status
		returnError = &errorData

	}
	var result BaseResponse
	if returnError == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(body, &result)
		if result.Error.Code != 0 {
			var errorData SdkError
			errorData.Code = fmt.Sprintf("%s.%d", SfapiError, result.Error.Code)
			errorData.Detail = fmt.Sprintf("%d:%s", result.Error.Code, result.Error.Message)
			returnError = &errorData
		} else {
			tmpResults, reqerr := json.Marshal(result.Result)
			if reqerr == nil {
				json.Unmarshal(tmpResults, &res)
			}
		}
	}
	log.WithContext(ctx).Debugf("Ending call %s", method)
	return result, returnError
}

func makeBasicAuthHeader(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func mkHttpClient() *http.Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	return client
}
