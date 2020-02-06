package sdk

import "fmt"

type BaseRequest struct {
	Method     string      `json:"method"`
	Id         int32       `json:"id"`
	Parameters interface{} `json:"params"`
}

const SfapiError = "sfapi"
const NetworkError = "http"

type SdkError struct {
	Code   string
	Detail string
}

func (this *SdkError) Error() string {
	return fmt.Sprintf("%s : %s", string(this.Code), this.Detail)

}

//error returned from the SF API
type SFAPIError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Name    string `json:"name"`
}
type BaseResponse struct {
	Id     int32       `json:"id"`
	Result interface{} `json:"result"`
	Error  SFAPIError  `json:"error"`
}
type SFClient struct {
	userId   string
	password string
	baseUrl  string
}

//a client that has nothing but stubs that return an error
//Can be used as a base for overriding stubs for testing
type SFStubClient struct {
}
