package version

import "fmt"

var (
	osArch        = "unspecified"
	gitCommit     = "unspecified"
	gitCommitFull = "unspecified"
	buildDate     = "unspecified"
	version       = "unspecified"
)

type Info struct {
	OSArch        string
	GitCommit     string
	GitCommitFull string
	BuildDate     string
	Version       string
}

func Get() Info {
	return Info{
		OSArch:        osArch,
		GitCommit:     gitCommit,
		GitCommitFull: gitCommitFull,
		BuildDate:     buildDate,
		Version:       version,
	}
}

func (i Info) String() string {
	return fmt.Sprintf(
		"go-assign %s Version: %s, GitCommit: %s, GitCommitFull: %s, BuildDate: %s",
		i.OSArch,
		i.Version,
		i.GitCommit,
		i.GitCommitFull,
		i.BuildDate,
	)
}
