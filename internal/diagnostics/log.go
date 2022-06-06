package diagnostics

import (
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/davidalpert/go-git-mob/internal/diagnostics/plaintext"
	"github.com/davidalpert/go-git-mob/internal/env"
	"github.com/davidalpert/go-git-mob/internal/version"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var Log *log.Entry

const (
	ENVKEY_LOG_LEVEL  = "GITMOB_LOG_LEVEL"
	ENVKEY_LOG_FORMAT = "GITMOB_LOG_FORMAT"
	ENVKEY_LOG_FILE   = "GITMOB_LOG_FILE"
	ENVKEY_DEBUG      = "GITMOB_DEBUG"
)

func init() {
	Log = log.WithFields(log.Fields{
		"app":         version.Detail.AppName,
		"app_version": version.Detail.Version,
	})

}

func ConfigureLogger(streams utils.IOStreams) (cleanupFn func()) {
	// default cleanup: nothing to do
	cleanupFn = func() {}

	// configure logging
	log.SetLevel(env.GetValueOrDefaultLogLevel(ENVKEY_LOG_LEVEL, log.FatalLevel))
	// log sink
	var sink io.Writer
	var logFile = env.GetValueOrDefaultString(ENVKEY_LOG_FILE, "")
	if logFile == "" {
		sink = streams.Out
	} else {
		fullPath, err := filepath.Abs(logFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		logFile, err := os.Create(fullPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		cleanupFn = func() { logFile.Close() }
		if env.GetValueOrDefaultBool(ENVKEY_DEBUG, false) {
			fmt.Println("logging to:", fullPath)
		}
		sink = logFile
	}

	if strings.EqualFold(env.GetValueOrDefaultString(ENVKEY_LOG_FORMAT, "text"), "json") {
		log.SetHandler(json.New(sink))
	} else {
		log.SetHandler(plaintext.New(sink))
	}

	Log.Debug("logging initialized")

	return
}
