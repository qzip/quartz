package arango

import "qz/splayd/fs"

/*
implements ../fs/SplayedStore

   arangoStore
   ArangoFsConfig  provides name of dburl, database &collections

*/

func NewArangoFs(cfg ArangoFsConfig) (store fs.SplayedStore, err error) {
	client := &ArangoClient{
		DbURL:    cfg.DbURL,
		User:     cfg.User,
		Pass:     cfg.Pass,
		AuthType: cfg.AuthType,
	}
	store = &arangoStore{
		client: client,
		cfg:    cfg,
	}
	return
}

type ArangoFsConfig struct {
	DbURL    string `json:"dbURL"` // if starts with $ it is assumed as env var
	User     string `json:"user,omitempty"`
	Pass     string `json:"pass,omitempty"`
	AuthType string `json:"auth,omitempty"` // nil, none, jwt, raw,basic

	DbName    string `json:"dbname"`
	ClassName string `json:"classname"`
	CASname   string `json:"casname"`
}

type arangoStore struct {
	client *ArangoClient
	cfg    ArangoFsConfig
}

func (as *arangoStore) LoadCAS(hash fs.HashData) (content []byte, err error) {

	return
}

// SaveCAS does not break content in chuncks
// inserts the content and returns the content hash
func (as *arangoStore) SaveCAS(content []byte) (hash fs.HashData, err error) {
	hash = fs.HashContent(content)

	return
}

func (as *arangoStore) SaveSplay(sfile *fs.SplayFile) (hash fs.HashData, err error) {
	return
}

func (as *arangoStore) LoadSplayByUUID(uuid string) (file *fs.SplayFile, err error) {
	return
}

func (as *arangoStore) LoadUUID(path, className string) (uuid []string, err error) {
	return
}
