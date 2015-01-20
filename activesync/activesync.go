// activesync project activesync.go
package activesync

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

type Device struct {
	XMLName                 xml.Name `xml:"Settings settings:Set"`
	DeviceID                string   `xml:"settings:DeviceID"`
	DeviceType              string   `xml:"settings:DeviceType"`
	Model                   string   `xml:"settings:Model,omitempty"`
	IMEI                    string   `xml:"settings:IMEI,omitempty"`
	FriendlyName            string   `xml:"settings:FriendlyName,omitempty"`
	OperatingSystem         string   `xml:"settings:OS,omitempty"`
	OperatingSystemLanguage string   `xml:"settings:OSLanguage,omitempty"`
	PhoneNumber             string   `xml:"settings:PhoneNumber,omitempty"`
	MobileOperator          string   `xml:"settings:MobileOperator,omitempty"`
	UserAgent               string   `xml:"settings:UserAgent,omitempty"`
}

type NetworkCredentials struct {
	Domain   string
	Username string
	Password string
}

func NetworkCredentialsInit(cred NetworkCredentials) bool {
	if cred.Username == "" && cred.Password == "" {
		return false
	}
	return true
}

func GetBasicUsername(cred NetworkCredentials) string {
	if cred.Domain != "" {
		return fmt.Sprintf("%s/%s", cred.Domain, cred.Username)
	}
	return cred.Username
}

func GetDeviceInformationNode(dev Device) string {
	data, err := xml.MarshalIndent(dev, "", "  ")
	if err != nil {
		//handle error
	}
	buf := bytes.NewBuffer(data)

	// the Marshal method creates the xml header
	//		<settings:Set xmlns="Settings">
	// but we need
	//		<settings:Set xmlns:settings="Settings">
	// I have no clue on how to get this done with Go, so
	// we'll just replace the first line of the output XML
	ret := strings.Replace(buf.String(), "xmlns", "xmlns:settings", 1)

	return (ret)
}
