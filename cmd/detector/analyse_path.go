package detector

import (
	"github.com/yildizozan/gandalf/cmd/config"
	"regexp"
	"strings"
)

func analysePath(rules *config.Path, path *string, ch chan bool) {

	if len(rules.Prefix) != 0 {
		result := strings.HasPrefix(*path, rules.Prefix)
		if result {
			ch <- true
			return
		}

	}

	if len(rules.Match) != 0 {
		if regex(*path, regexp.QuoteMeta(rules.Match)) {
			ch <- true
			return
		}
	}

	if len(rules.Exact) != 0 {
		if strings.TrimRight(*path, "\n") == rules.Exact {
			ch <- true
			return
		}
	}

	ch <- false
}
