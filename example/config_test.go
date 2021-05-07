package example

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	tmpl, err := template.ParseFS(TestDataDir, "testdata/config-tmpl.json")
	require.NoError(t, err)

	td := os.TempDir()
	defer os.RemoveAll(td)

	dstHelloCredPath := filepath.Join(td, "hello-api-credentials.txt")
	dstWorldCredPath := filepath.Join(td, "world-api-credentials.txt")

	m := map[string]string{
		"HelloAPITokenFile": dstHelloCredPath,
		"WorldAPITokenFile": dstWorldCredPath,
	}

	js := bytes.NewBufferString("")
	err = tmpl.Execute(js, m)
	require.NoError(t, err)

	dstHelloCred, err := os.Create(dstHelloCredPath)
	require.NoError(t, err)
	dstWorldCred, err := os.Create(dstWorldCredPath)
	require.NoError(t, err)

	srcHelloCred, err := TestDataDir.Open("testdata/hello-api-credentials.txt")
	require.NoError(t, err)
	srcWorldCred, err := TestDataDir.Open("testdata/world-api-credentials.txt")
	require.NoError(t, err)

	_, err = io.Copy(dstHelloCred, srcHelloCred)
	require.NoError(t, err)
	_, err = io.Copy(dstWorldCred, srcWorldCred)
	require.NoError(t, err)

	var g ConfigGenerated
	err = json.Unmarshal(js.Bytes(), &g)
	assert.NoError(t, err)
}
