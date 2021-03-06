# go-assign

go-assign is a tool that supports unmarshal structs from data and provides a generator to generate extended structs.

# Motivation

Nowadays, the workload of managing application configuration with GitOps is popular. This style is efficient and easy to operate, but secret management requires more care.

In regard to security, the secret data should not be committed to the repository, even if you're using a private repository. Therefore, the secret should be injected into the application from outside the repository.

For example, consider a workload that manages application configuration files with json. You must exclude the secret from json to include this file in the repository. To achieve this, you need to set the secret file path in the config file and load it from your application.

```json
{
  "helloAPIDomain": "hello-api.example.com",
  "helloAPITokenFile": "/path-to-path/secrets/hello-api-credential.txt"
}
```

This is a good idea, but as the number of secrets increases, so does the cost of implementation in your application. go-assign solves this problem with its own tags and generator. 

# Installation

```shell
$ go install github.com/stormcat24/go-assign/cmd/go-assign@v0.0.2
```

# Usage

## Original file

```go
package example

//go:generate go-assign generate

type Config struct {
	HelloAPIDomain string `json:"helloAPIDomain"`
	HelloAPIToken  []byte `json:"helloAPIToken" go-assign:"fileFieldName=HelloAPITokenFile,fileFieldTag=json:\"helloAPITokenFile\""`
	WorldAPIDomain string `json:"worldAPIDomain"`
	WorldAPIToken  []byte `json:"worldAPIToken" go-assign:"fileFieldName=WorldAPITokenFile,fileFieldTag=json:\"worldAPITokenFile\",base64=true"`
}
```

When you run `go generate`, the following extended structure will be generated.

```go
// Code generated by go-assign; DO NOT EDIT.

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
	s := struct {
		*Alias
		HelloAPITokenFile string `json:"helloAPITokenFile"`
		WorldAPITokenFile string `json:"worldAPITokenFile"`
	}{
		Alias: (*Alias)(g),
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	// Assign "HelloAPITokenFile" and "HelloAPIToken"
	g.HelloAPITokenFile = s.HelloAPITokenFile
	helloApitokenFileData, err := os.ReadFile(s.HelloAPITokenFile)
	if err != nil {
		return err
	}
	g.HelloAPIToken = helloApitokenFileData

	// Assign "WorldAPITokenFile" and "WorldAPIToken"
	g.WorldAPITokenFile = s.WorldAPITokenFile
	worldApitokenFileData, err := os.ReadFile(s.WorldAPITokenFile)
	if err != nil {
		return err
	}
	g.WorldAPIToken = make([]byte, base64.StdEncoding.EncodedLen(len(worldApitokenFileData)))
	if _, err := base64.StdEncoding.Decode(g.WorldAPIToken, worldApitokenFileData); err != nil {
		return err
	}

	return nil
}
```


License
===
See [LICENSE](LICENSE).

Copyright ?? stromcat24. All Rights Reserved.
