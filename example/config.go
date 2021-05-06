package example

//go:generate go-assign

type Config struct {
	HelloAPIDomain string `json:"helloAPIDomain"`
	HelloAPIToken  []byte `json:"helloAPIToken" go-assign:"sourceFileField=HelloAPITokenFile"`
	WorldAPIDomain string `json:"worldAPIDomain"`
	WorldAPIToken  []byte `json:"worldAPIToken" go-assign:"sourceFileField=WorldAPITokenFile,base64=true"`
}
