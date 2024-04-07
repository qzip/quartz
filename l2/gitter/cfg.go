package gitter

import (
	"encoding/json"
	"os"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Cfg struct {
	GitRoot         string `json:"gitRoot,omitempty"`    // default "." current dir
	BlindNotaryPath string `json:"notaryPath,omitempty"` // relative to root
	GitUser         Author `json:"author"`
}

func New(jsonFile string) (cfg *Cfg, err error) {
	cfg = &Cfg{GitRoot: ".", BlindNotaryPath: "blind-notary"}
	jsonb, erx := os.ReadFile(jsonFile)
	if erx != nil {
		err = erx
		return
	}
	err = json.Unmarshal(jsonb, &cfg)
	return
}
