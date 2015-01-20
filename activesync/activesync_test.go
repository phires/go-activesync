package activesync

import "testing"

func TestGetDeviceInformationNode(t *testing.T) {
	dev := Device{}
	res := GetDeviceInformationNode(dev)
	if res != `<settings:Set xmlns:settings="Settings">
  <settings:DeviceID></settings:DeviceID>
  <settings:DeviceType></settings:DeviceType>
</settings:Set>` {
		t.Error("Unexpected GetDeviceInformationNode response")
	}
}

func TestNetworkCredentialInit(t *testing.T) {
	cred := NetworkCredentials{"Domain", "Username", "Password"}
	if NetworkCredentialsInit(cred) != true {
		t.Error("NetworkCredentialInit wrong response")
	}
}

func TestGetBasicUsername(t *testing.T) {
	cred := NetworkCredentials{"Domain", "Username", "Password"}
	if GetBasicUsername(cred) != "Domain/Username" {
		t.Error("GetBasicUsername wrong response (with Domain)")
	}

	cred = NetworkCredentials{"", "Username", "Password"}
	if GetBasicUsername(cred) != "Username" {
		t.Error("GetBasicUsername wrong response (without Domain)")
	}
}
