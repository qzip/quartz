package comp

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"

	//"html/template"
	"bc/cas"
	"io/ioutil"
	"mime"
	"net/http"
	"qz/commands"
	"qz/defs"
	"qz/dsl"
	"qz/seq"
	"qz/util"
	"reflect"
	"strings"
	"text/template"
	"time"
)

// DocZipper json to ziped html generated from a DSL tempate
type DocZipper struct {
}

// ImagesDir used as default images dir
const ImagesDir = "images"

// implements seq.Runner interface methods

// Name  implements seq.Runner interface methods
func (dz *DocZipper) Name() string {
	return "comp.DocZipper"
}

// Help  implements seq.Runner interface methods
func (dz *DocZipper) Help() string {
	return `
	# Converts the JSON Doc to Zip file
	# using DSL (Domain Specific Lang.) Template
	
	{
	 "component": "comp.DocZipper",
	 "param": "<dsl template file>"
	}
	`
}

// ComponentType component type
func (dz *DocZipper) ComponentType() reflect.Type {
	return reflect.TypeOf(dz)
}

// Create doc zipper object
func (dz *DocZipper) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errChan chan error) commands.Pipeline {
	zp := &zipperPipeline{helper: helper, errChan: errChan}

	fl := fmt.Sprintf("%v", param)
	util.DebugInfo(ctx, fmt.Sprintf("DocZipper.Create file [%v] START", fl))
	fbuf, err := ioutil.ReadFile(fl)
	if err != nil {
		util.DebugInfo(ctx, fmt.Sprintf("DocZipper.Create: file: [%v], err: %v ", fl, err.Error()))
		helper.SetExecStatus(seq.ExSerror)
		errChan <- commands.NewFatalError(err.Error())
		return nil
	}
	prs := dsl.NewParser(bytes.NewReader(fbuf))
	zp.blk, err = prs.Parse(ctx)
	if err != nil {
		util.DebugInfo(ctx, fmt.Sprintf("DocZipper.Create: file: [%v], err: %v ", fl, err.Error()))
		helper.SetExecStatus(seq.ExSerror)
		errChan <- commands.NewFatalError(err.Error())
		return nil
	}
	helper.SetExecStatus(seq.ExSinit)
	util.DebugInfo(ctx, fmt.Sprintf("DocZipper.Create file [%v] END OK", fl))
	return zp
}

type zipperPipeline struct {
	helper  seq.CtxHelper
	blk     []dsl.Block
	errChan chan error
}

// Process implements commands.Pipeline method
func (z *zipperPipeline) Process(ctx context.Context) {
	z.helper.SetExecStatus(seq.ExSrunning)
	util.DebugInfo(ctx, "zipperPipeline.Process: START ")
	v := ctx.Value(defs.DocRec)
	if v == nil {
		z.errChan <- commands.NewFatalError("zipperPipeline.Process: NOT FOUND: context value for key reflect.TypeOf(*defs.DocSubCols)")
		z.helper.SetExecStatus(seq.ExSerror)
		return
	}
	dsc, ok := v.(*defs.DocSubCols)
	if !ok {
		z.errChan <- commands.NewFatalError("zipperPipeline.Process: NOT of type. context value for key reflect.TypeOf(*defs.DocSubCols)")
		z.helper.SetExecStatus(seq.ExSerror)
		return
	}
	ast := &doc2zipAST{data: dsc, blk: z.blk, macros: make(map[string][]byte)}
	if err := ast.Block2Ast(ctx); err != nil {
		z.errChan <- commands.NewFatalError(err.Error())
		z.helper.SetExecStatus(seq.ExSerror)
		return
	}
	zf := &GenZip{files: ast.files}
	zf.zipName = ast.zipFlname
	if len(zf.zipName) == 0 {
		zf.zipName = "/tmp/ziped-doc.zip"
	}
	zf.zipName = normalizeFileName(zf.zipName, dsc.Doc)
	util.DebugInfo(ctx, fmt.Sprintf("zipperPipeline.Process: ZIPFILE=%v", zf.zipName))
	if err := zf.Generate(ctx); err != nil {
		util.DebugInfo(ctx, fmt.Sprintf("zipperPipeline.Process: zip gen err %v", err.Error()))
		z.errChan <- commands.NewFatalError(err.Error())
		z.helper.SetExecStatus(seq.ExSerror)
		return
	}
	z.helper.SetExecStatus(seq.ExSok)
	util.DebugInfo(ctx, "zipperPipeline.Process: END OK ")
}

func normalizeFileName(flname string, dat map[string]interface{}) string {
	st := strings.Index(flname, "{{")
	en := strings.Index(flname, "}}")
	if st < 0 || st > en {

		return flname
	}
	prefix := ""
	postfix := ""
	if st > 0 {
		prefix = flname[:st]
	}
	if en < (len(flname) - 3) {
		postfix = flname[en+2:]
	}
	infixVar := flname[st+2 : en]
	if infix, ok := dat[infixVar]; ok {
		return fmt.Sprintf("%v%v%v", prefix, infix, postfix)
	}

	return flname

}

// doc2zipAST AST (Abstract Syntax Tree) for processing JSON to HTML zip
// assumes every thing fits into mem
type doc2zipAST struct {
	data *defs.DocSubCols
	blk  []dsl.Block
	//keyvals map[string]string
	zipFlname string
	subCol    string
	macros    map[string][]byte
	files     []Zfile
}

func (za *doc2zipAST) Block2Ast(ctx context.Context) (err error) {
	util.DebugInfo(ctx, "doc2zipAST.Block2Ast: START. getting images (commented off)")
	if err = za.getImages(); err != nil {
		return
	}
	util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.Block2Ast: got images. Processing Node 0, total nodes %v", len(za.blk)))
	// Process root Block[0]
	for _, n := range za.blk[0].Nodes {
		key := strings.ToLower(strings.TrimSpace(n.Key))
		switch {
		case key == "macro":
			var b bytes.Buffer
			for i := 1; i < len(n.Lines); i++ {
				b.Write([]byte(n.Lines[i]))
			}
			za.macros[strings.TrimSpace(n.Lines[0])] = b.Bytes()
		case key == "zipfile":
			za.zipFlname = strings.TrimSpace(n.Lines[0])
		case key == "copyzip":
			zcpy := strings.TrimSpace(n.Lines[0])
			if err = za.copyzip(ctx, zcpy); err != nil {
				util.DebugInfo(ctx, err.Error())
				return
			}
		default:
			util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.Block2Ast: root keyword %v ignored", key))

		}
	}
	util.DebugInfo(ctx, "doc2zipAST.Block2Ast: Processed Node 0")
	// get children
	children := (dsl.Blocks(za.blk)).ChilderenOf(0)
	util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.Block2Ast: Processing child nodes %v", len(children)))
	for _, ndx := range children {
		blk := za.blk[ndx]
		if strings.EqualFold("template", strings.TrimSpace(blk.Key)) {
			if err = za.processTemplate(ctx, &blk); err != nil {
				util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.Block2Ast: template [%v] processing error %v", blk.Name(), err.Error()))
				return
			}
		} else {
			util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.Block2Ast: ignoring block %v", blk.Key))
		}

	}
	// File Blocks
	util.DebugInfo(ctx, "doc2zipAST.Block2Ast: END OK")
	return
}

func (za *doc2zipAST) copyzip(ctx context.Context, zipName string) error {
	r, err := zip.OpenReader(zipName)
	if err != nil {
		return err
	}
	defer r.Close()
	util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.copyzip: opened %v", zipName))
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		//TODO: use go routine
		zf := Zfile{path: f.Name}
		if zf.content, err = ioutil.ReadAll(rc); err != nil {
			return err
		}
		za.files = append(za.files, zf)
		rc.Close()

	}
	return nil
}

/*
optional:
subCol <sub doc name>

required keywords inside BEGIN template ...END
file <file path within the zip>
header, body, footer
*/
func (za *doc2zipAST) processTemplate(ctx context.Context, blk *dsl.Block) (err error) {
	zf := Zfile{path: "undefined-" + time.Now().String() + ".html"}
	var data interface{}
	tmplArr := make([]SubTemplate, 3)
	data = za.data.Doc
	for _, keyLine := range blk.Nodes {
		key := strings.ToLower(strings.TrimSpace(keyLine.Key))
		switch {
		case key == "file":
			if len(keyLine.Lines) > 0 {
				zf.path = strings.TrimSpace(keyLine.Lines[0])
			} else {
				util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.processTemplate: keyword file near line %v, name missing", keyLine.Number))
			}

		case key == "subcol":
			if len(keyLine.Lines) > 0 {
				p := strings.TrimSpace(keyLine.Lines[0])
				ok := true
				data, ok = za.data.SubCols[p]
				if !ok {
					util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.processTemplate: keyword subCol near line %v, does not have any member %v", keyLine.Number, p))
				}
			} else {
				util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.processTemplate: keyword file near line %v, name missing", keyLine.Number))
			}
		case key == "header":
			//tmplArr = append(tmplArr, SubTemplate{Name: key, Content: keyLine.Lines2String()})
			tmplArr[0] = SubTemplate{Name: key, Content: keyLine.Lines2String()}
		case key == "body":
			//tmplArr = append(tmplArr, SubTemplate{Name: key, Content: keyLine.Lines2String()})
			tmplArr[1] = SubTemplate{Name: key, Content: keyLine.Lines2String()}
		case key == "footer":
			//tmplArr = append(tmplArr, SubTemplate{Name: key, Content: keyLine.Lines2String()})
			tmplArr[2] = SubTemplate{Name: key, Content: keyLine.Lines2String()}
		default:
			util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.processTemplate: ignoring keyword %v at line %v", keyLine.Key, keyLine.Number))
		}
	}
	var funcMap = make(template.FuncMap)
	/*
		funcMap["macro"] = func(m string) (string, error) {
			if mac, ok := za.macros[m]; ok {
				return string(mac), nil
			}
			err = fmt.Errorf("doc2zipAST.processTemplate: macro %v not found", m)
			util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.processTemplate:macro err %v", err.Error()))
			return "", err
		}*/

	for k, m := range za.macros {
		funcMap[k] = func() string {
			mx := string(m)
			return mx
		}
	}
	// func: article
	funcMap["article"] = func(id, ele string) string {
		arts, ok := za.data.SubCols["articles"]
		if !ok {
			return "articles subcol not found"
		}
		rec, ok := arts[id]
		if !ok {
			return "articles" + id + " not found"
		}
		m, ok := rec.(map[string]interface{})
		if !ok {
			return "articles" + id + " not of map[string]interface{} type"
		}
		return fmt.Sprintf("%v", m[ele])
	}
	funcMap["str"] = func(v interface{}) string {
		return fmt.Sprintf("%v", v)
	}
	funcMap["subcol"] = func(scol string) (map[string]map[string]interface{}, error) {
		util.DebugInfo(ctx, "subcol")
		ele, ok := za.data.SubCols[scol]
		if !ok {
			return nil, fmt.Errorf("subcol: %v not found", scol)
		}
		ret := make(map[string]map[string]interface{})
		for k, v := range ele {
			ret[k], ok = v.(map[string]interface{})
			if !ok {
				kv := make(map[string]interface{})
				kv[k] = v
				ret[k] = kv
			}
		}
		return ret, nil
	}
	util.DebugInfo(ctx, "doc2zipAST.processTemplate: Processing templates")

	tmplr := NewTemplater(data, funcMap, tmplArr)
	if zf.content, err = tmplr.Execute(ctx); err == nil {
		util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.processTemplate:processed template file [%v]", zf.path))
		za.files = append(za.files, zf)
	}
	util.DebugInfo(ctx, fmt.Sprintf("doc2zipAST.processTemplate: file: %v, len: %v  err %v", zf.path, len(zf.content), err))

	return
}

func (za *doc2zipAST) getImages() error {
	if profImg, ok := za.data.Doc["profileImage"]; ok {
		if err := za.getImage(fmt.Sprintf("%v", profImg)); err != nil {
			return err
		}
		za.data.Doc["profileImage"] = za.files[len(za.files)-1].path
	}
	if actions, ok := za.data.SubCols["actions"]; ok {
		for _, v := range actions {
			if m, ok := v.(map[string]interface{}); ok {
				if img, ok := m["image"]; ok {
					if err := za.getImage(fmt.Sprintf("%v", img)); err != nil {
						return err
					}
				}
				m["image"] = za.files[len(za.files)-1].path
			}
		}
	}
	if gallery, ok := za.data.SubCols["gallery"]; ok {
		for _, v := range gallery {
			if m, ok := v.(map[string]interface{}); ok {
				if img, ok := m["image"]; ok {
					if err := za.getImage(fmt.Sprintf("%v", img)); err != nil {
						return err
					}
				}
				m["image"] = za.files[len(za.files)-1].path
			}
		}
	}
	if journals, ok := za.data.SubCols["journals"]; ok {
		for _, v := range journals {
			if m, ok := v.(map[string]interface{}); ok {
				if img, ok := m["image"]; ok {
					if err := za.getImage(fmt.Sprintf("%v", img)); err != nil {
						return err
					}
				}
				m["image"] = za.files[len(za.files)-1].path
			}
		}
	}

	return nil
}

func (za *doc2zipAST) handleMap(m map[string]interface{}) error {

	return nil
}

func (za *doc2zipAST) getImage(url string) error {
	buf, err := GetContent(url)
	if err != nil {
		return err
	}
	ch := cas.NewHashData(buf)
	fname := ch.String()
	types, err := mime.ExtensionsByType(http.DetectContentType(buf))
	if err != nil {
		return err
	}
	zf := Zfile{path: ImagesDir + "/" + fname + types[0],
		content: buf}
	za.files = append(za.files, zf)
	return nil
}

type processor func(context.Context, *Zfile, *dsl.Block) error

var processors map[string]processor

func init() {

}

func setMacro(ctx context.Context, zf *Zfile, blk *dsl.Block) error {

	return nil
}
func setZipFlname(ctx context.Context, zf *Zfile, blk *dsl.Block) error {

	return nil
}
