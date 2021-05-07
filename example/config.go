package example

//go:generate go-assign generate

type Config struct {
	HelloAPIDomain string `json:"helloAPIDomain"`
	HelloAPIToken  []byte `json:"helloAPIToken" go-assign:"fileFieldName=HelloAPITokenFile,fileFieldTag=json:\"helloAPITokenFile\""`
	WorldAPIDomain string `json:"worldAPIDomain"`
	WorldAPIToken  []byte `json:"worldAPIToken" go-assign:"fileFieldName=WorldAPITokenFile,fileFieldTag=json:\"worldAPITokenFile\",base64=true"`
}
