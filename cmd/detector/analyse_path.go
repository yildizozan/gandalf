package detector

import (
	"github.com/spf13/viper"
	"regexp"
	"strings"
)

func analysePath(path string, ch chan bool) {
	prefix := viper.GetString("app.rules.path.prefix")
	exact := viper.GetString("app.rules.path.exact")
	match := viper.GetString("app.rules.path.match")

	if len(prefix) != 0 {
		result := strings.HasPrefix(path, prefix)
		if result {
			ch <- true
			return
		}

	}

	if len(exact) != 0 {
		if regex(path, regexp.QuoteMeta(match)) {
			ch <- true
			return
		}
	}

	if len(match) != 0 {
		if strings.TrimRight(path, "\n") == exact {
			ch <- true
			return
		}
	}

	ch <- false
}
