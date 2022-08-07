package gitMessage

import (
	"github.com/davidalpert/go-git-mob/internal/env"
	"github.com/davidalpert/go-git-mob/internal/gitCommands"
	"path"
)

const (
	EnvKeyGitMobMessagePath = "GITMOB_MESSAGE_PATH"
)

// Path is ported from git-mob/src/git-message/index/gitMessagePath
func Path() string {
	return env.GetValueOrDefaultString(EnvKeyGitMobMessagePath, gitCommands.GetTemplatePath())
}

// CommitTemplatePath is ported from git-mob/src/git-message/index/commitTemplatePath
func CommitTemplatePath() string {
	return env.GetValueOrDefaultString(EnvKeyGitMobMessagePath, path.Join(env.HomeDir, ".gitmessage"))
}
