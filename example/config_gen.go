// Code generated by go-assign; DO NOT EDIT.

package example

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type ConfigGenerated struct {
	Config
	HelloAPITokenFile string `json:"helloAPITokenFile"`
	WorldAPITokenFile string `json:"worldAPITokenFile"`
}

func (g *ConfigGenerated) UnmarshalJSON(data []byte) error {
	type Alias ConfigGenerated
	s := struct {
		*Alias
		HelloAPITokenFile string `json:"helloAPITokenFile"`
		WorldAPITokenFile string `json:"worldAPITokenFile"`
	}{
		Alias: (*Alias)(g),
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("go-assign: %w", err)
	}
	// Assign "HelloAPITokenFile" and "HelloAPIToken"
	g.HelloAPITokenFile = s.HelloAPITokenFile
	helloApitokenFileData, err := os.ReadFile(s.HelloAPITokenFile)
	if err != nil {
		return fmt.Errorf("go-assign(HelloAPITokenFile): %w", err)
	}
	g.HelloAPIToken = helloApitokenFileData

	// Assign "WorldAPITokenFile" and "WorldAPIToken"
	g.WorldAPITokenFile = s.WorldAPITokenFile
	worldApitokenFileData, err := os.ReadFile(s.WorldAPITokenFile)
	if err != nil {
		return fmt.Errorf("go-assign(WorldAPITokenFile): %w", err)
	}
	g.WorldAPIToken = make([]byte, base64.StdEncoding.EncodedLen(len(worldApitokenFileData)))
	if _, err := base64.StdEncoding.Decode(g.WorldAPIToken, worldApitokenFileData); err != nil {
		return fmt.Errorf("go-assign(WorldAPITokenFile): %w", err)
	}

	return nil
}
