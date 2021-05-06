package example

//go:generate go-assign

type Config struct {
	APIDomain    string `json:"apiDomain"`
	APIToken     string `json:"apiToken" sourceFileField:"APITokenFile"`
	SSHPublicKey string `json:"sshPublicKey" sourceFileField:"SSHPublicKeyFile"`
}
