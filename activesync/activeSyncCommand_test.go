package activesync

import "testing"

func TestGetSslStringAppendix(t *testing.T) {
	if GetSslStringAppendix(true) != "s" || GetSslStringAppendix(false) != "" {
		t.Error("GetSslStringAppendix did not return expected result")
	}
}

func TestBuildUriSsl(t *testing.T) {
	uriWithSsl := BuildUri("TestServer", "RequestLine", true)
	if uriWithSsl != "https://TestServer/Microsoft-Server-ActiveSync?RequestLine" {
		t.Error("BuildUri with SSL failed")
	}

	uriWithoutSsl := BuildUri("TestServer", "RequestLine", false)
	if uriWithoutSsl != "http://TestServer/Microsoft-Server-ActiveSync?RequestLine" {
		t.Error("BuildUri without SSL failed")
	}
}
