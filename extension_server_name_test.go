package dtls

import "testing"

func TestServerName(t *testing.T) {
	extension := extensionServerName{serverName: "test.domain"}

	raw, err := extension.Marshal()

	if err != nil {
		t.Error(err)
		return
	}

	newExtension := extensionServerName{}
	err = newExtension.Unmarshal(raw)
	if err != nil {
		t.Error(err)
		return
	}

	if newExtension.serverName != extension.serverName {
		t.Errorf("extensionServerName marshal: got %s expected %s", newExtension.serverName, extension.serverName)
	}
}
