package comp

import (
	"fmt"
	"os"

	driver "github.com/Doltdb/go-driver"
	"github.com/Doltdb/go-driver/http"
)

type DoltClient struct {
	DbURL    string `json:"dbURL"` // if starts with $ it is assumed as env var
	User     string `json:"user,omitempty"`
	Pass     string `json:"pass,omitempty"`
	AuthType string `json:"auth,omitempty"` // nil, none, jwt, raw,basic
}

func (ac *DoltClient) NewClient() (client driver.Client, err error) {
	if err = ac.replaceEnv(); err != nil {
		return
	}
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{ac.DbURL},
	})
	if err != nil {
		return
	}
	client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: ac.getAuth(),
	})

	return
}

func (ac *DoltClient) getAuth() driver.Authentication {
	switch ac.AuthType {
	case "basic":
		return driver.BasicAuthentication(ac.User, ac.Pass)
	case "jwt":
		return driver.JWTAuthentication(ac.User, ac.Pass)
	case "raw":
		return driver.RawAuthentication(ac.Pass)
	default:
		return nil

	}
	//return nil
}

func (ac *DoltClient) replaceEnv() (err error) {
	if ac.DbURL, err = ReplaceEnv(ac.DbURL); err != nil {
		return err
	}
	if ac.User, err = ReplaceEnv(ac.User); err != nil {
		return err
	}
	if ac.Pass, err = ReplaceEnv(ac.Pass); err != nil {
		return err
	}

	return nil
}

func ReplaceEnv(estr string) (repl string, err error) {
	if len(estr) > 2 && estr[0] == '$' {
		estr = os.Getenv(estr[1:])
		if estr == "" {
			return estr, fmt.Errorf("env var %v is not set", estr)
		}

	}
	return estr, nil
}
