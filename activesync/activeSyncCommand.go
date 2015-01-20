package activesync

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

type CommandParameter struct {
	Parameter string
	Value     string
}

type ASCommandRequest struct {
	UseSSL                bool
	WBXMLBytes            []byte
	XmlString             string
	Server                string
	RequestLine           string
	UseEncodedRequestLine bool
	ProtocolVersion       string
	Command               string
	PolicyKey             uint32
	CommandParameter      []CommandParameter
	Device                Device
	Credentials           NetworkCredentials
}

type ASCommandResponse struct {
	WBXMLBytes     []byte
	XmlString      string
	HttpStatusCode int
}

func BuildUri(Server string, RequestLine string, UseSsl bool) string {
	uri := fmt.Sprintf("http%s://%s/Microsoft-Server-ActiveSync?%s", GetSslStringAppendix(UseSsl), Server, RequestLine)

	return uri
}

func GetSslStringAppendix(UseSsl bool) string {
	if UseSsl == true {
		return "s"
	}
	return ""
}

func GetResponse(asCmdReq ASCommandRequest) (ASCommandResponse, error) {
	var response ASCommandResponse
	var errorString string

	if NetworkCredentialsInit(asCmdReq.Credentials) == false {
		errorString = "No credentials"
	}
	if asCmdReq.Server == "" {
		errorString = "Server not set"
	}
	if asCmdReq.ProtocolVersion == "" {
		errorString = "ProtocolVersion not set"
	}
	if asCmdReq.WBXMLBytes == nil {
		errorString = "WBMXMLBytes not set"
	}

	if errorString != "" {
		return response, errors.New(errorString)
	}
	uri := BuildUri(asCmdReq.Server, asCmdReq.RequestLine, asCmdReq.UseSSL)

	client := &http.Client{}

	// Build Request object
	buf := bytes.NewReader(asCmdReq.WBXMLBytes)
	req, _ := http.NewRequest("POST", uri, buf)

	// Authentication
	req.SetBasicAuth(GetBasicUsername(asCmdReq.Credentials), asCmdReq.Credentials.Password)

	// Set headers
	req.Header.Add("Content-Type", "application/vnd.ms-sync.wbxml")
	if asCmdReq.UseEncodedRequestLine == false {
		req.Header.Add("MS-ASProtocolVersion", asCmdReq.ProtocolVersion)
		req.Header.Add("X-MS-PolicyKey", string(asCmdReq.PolicyKey))
	}

	res, _ := client.Do(req)
	defer res.Body.Close()

	return response, nil
}
