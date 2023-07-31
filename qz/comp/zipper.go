package comp

import (
	"archive/zip"
	"context"
	"os"
)

// Zfile path & content as string (for small content)
type Zfile struct {
	path    string
	content []byte
}

//GenZip generates zip file
type GenZip struct {
	files   []Zfile
	zipName string
}

//Generate generates the zip
func (zf *GenZip) Generate(ctx context.Context) error {
	fl, err := os.OpenFile(zf.zipName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer fl.Close()
	w := zip.NewWriter(fl)
	for _, file := range zf.files {
		//util.DebugInfo(ctx, fmt.Sprintf("GenZip.Generate: %v file : [%v]\n", i, file.path))
		f, err := w.Create(file.path)
		if err != nil {
			return err
		}
		if _, err = f.Write(file.content); err != nil {
			return err
		}

	}
	return w.Close()
}
