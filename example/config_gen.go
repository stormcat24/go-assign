package example

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

type ConfigGenerated struct {
	Config
	HelloAPITokenFile string `json:"helloAPITokenFile"`
	WorldAPITokenFile string `json:"worldAPITokenFile"`
}

func (g *ConfigGenerated) UnmarshalJSON(data []byte) error {
	type Alias ConfigGenerated
	y := struct {
		*Alias
		HelloAPITokenFile string `json:"helloAPITokenFile"`
		WorldAPITokenFile string `json:"worldAPITokenFile"`
	}{
		Alias: (*Alias)(g),
	}

	if err := json.Unmarshal(data, &y); err != nil {
		return err
	}

	helloAPIToken, err := os.ReadFile(y.HelloAPITokenFile)
	if err != nil {
		return err
	}
	worldAPIToken, err := os.ReadFile(y.WorldAPITokenFile)
	if err != nil {
		return err
	}
	g.HelloAPIToken = helloAPIToken
	g.HelloAPITokenFile = y.HelloAPITokenFile
	g.WorldAPIToken = make([]byte, base64.StdEncoding.EncodedLen(len(worldAPIToken)))
	g.WorldAPITokenFile = y.WorldAPITokenFile

	if _, err := base64.StdEncoding.Decode(g.WorldAPIToken, worldAPIToken); err != nil {
		return err
	}
	return nil
}
