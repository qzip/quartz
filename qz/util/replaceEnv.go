package util

import (
	"fmt"
	"os"
)

func ReplaceEnv(estr string) (repl string, err error) {
	if len(estr) > 2 && estr[0] == '$' {
		estr = os.Getenv(estr[1:])
		if estr == "" {
			return estr, fmt.Errorf("env var %v is not set", estr)
		}

	}
	return estr, nil
}
