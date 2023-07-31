package comp

import (
	"bytes"
	"context"
	"fmt"

	//"html/template"
	"qz/util"
	"text/template"
)

//Templater executes the html template and converts to zfile
type Templater interface {
	Execute(context.Context) ([]byte, error)
}

type templater struct {
	data    interface{}
	tmplArr []SubTemplate
	funcMap template.FuncMap
}

//NewTemplater construct a new object
func NewTemplater(data interface{}, funcMap template.FuncMap, tmplArr []SubTemplate) Templater {
	return &templater{data: data, funcMap: funcMap, tmplArr: tmplArr}
}

func (t *templater) Execute(ctx context.Context) ([]byte, error) {
	var b bytes.Buffer
	util.DebugInfo(ctx, fmt.Sprintf("templater.Execute: start template count %v", len(t.tmplArr)))
	for i, tm := range t.tmplArr {
		tx, err := template.New(tm.Name).Funcs(t.funcMap).Parse(tm.Content)
		util.DebugInfo(ctx, fmt.Sprintf("templater.Execute: ndx:%v, name:%v err:%v", i, tm.Name, err))
		if err != nil {

			return nil, err
		}
		if err = tx.Execute(&b, t.data); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

//SubTemplate sub template components,
// like header, body , footer
type SubTemplate struct {
	Name    string
	Content string
}
