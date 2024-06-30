/*
Copyright (c) 2017, BON BIZ IT Services Pvt Ltd

http://www.apache.org/licenses/LICENSE-2.0

Author: Ashish Banerjee, <ashish@bonbiz.in>


*/

package bdl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"sync"
	//	"plugin"
)

/*
*TODO:* 30-apr-2017

;Keyword: PLUGIN
  ; UUID Gen example
  PLUGIN uuid As "uuid.so"
  OBECT  GenUUID as uuid
  WITH GenUUID
     set count = 2
  END

  GO GenUUID

Go Plugin cannot be implemented till  https://github.com/golang/go/issues/19023 is fixed.
 go build -buildmode=plugin -o uuid.so uuid

.Alternative Design using plugin, but without the keyword PLUGIN
----

  OBJECT <instance name> as <symbol name> From <plugin name>
  WITH <instance name>
     SET <key> = <value>
  END
  GO <instance name>


----

*/

/*

  REM (also `;`)      : are comments
  OBJECT              : Instantiate an plugin.Symbol from a Class
  WITH and END WITH   : Delimits the SET Blocks
  SET                 : Calls the `setter` method of a class.
  GO  (also RUN)      : Synonym to RUN, execute the go call.



*/

type Object interface{}

type ObjectFactory interface {
	CreateInstance(objectName string) *Object
}

type Runnable interface {
	Run()
}

type ObjectMap map[string]*Object
type state uint

type BDL struct {
	objMap    ObjectMap
	err       error
	st        state
	curObjCtx string
}

// By default the StructTag id used for setting,
//  https://golang.org/pkg/reflect/#StructTag
// but optionally :
// the structure implementing the Runnable interface cal also implement Bean interface,
// if Bean interface is implemented then StructTag setting will not be used.

type Bean interface {
	Set(key string, val *Object) error
}

var registry = make(map[string]*ObjectFactory)

var WaitForAll sync.WaitGroup

func RegisterFactory(cpath string, fact *ObjectFactory) {
	registry[cpath] = fact
}

func New(omap ObjectMap) *BDL {
	b := &BDL{objMap: omap, err: nil, st: ST_START}
	return b
}

const (
	_ state = iota
	ST_START
	ST_OBJ
	ST_WITH
)

type token uint

const (
	_ token = iota
	REM
	OBJECT
	WITH
	END
	SET
	GO
	RUN = GO
)

var state2tokens = map[state][]token{
	ST_START: {REM, GO, OBJECT},
	ST_OBJ:   {SET, WITH, REM},
	ST_WITH:  {SET, END, REM},
}

func isValidToken(st state, tk token) bool {
	tks := state2tokens[st]
	if tks == nil {
		return false
	}
	ret := false

	for i := 0; i < len(tks); i++ {
		if tk == tks[i] {
			ret = true
			break
		}
	}
	return ret
}

var tokstrings = [...]string{
	REM:    "REM",
	OBJECT: "OBJECT",
	WITH:   "WITH",
	END:    "END",
	SET:    "SET",
	GO:     "GO",
}

var strTok = map[string]token{
	"REM":    REM,
	";":      REM,
	"OBJECT": OBJECT,
	"WITH":   WITH,
	"END":    END,
	"SET":    SET,
	"GO":     GO,
	"RUN":    GO,
}

// https://golang.org/pkg/io/ioutil/#ReadAll
func (bdl *BDL) ParseReader(r io.Reader) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return bdl.Parse(buf)
}

// https://golang.org/pkg/io/ioutil/#ReadFile
func (bdl *BDL) ParseFile(flname string) error {
	buf, err := ioutil.ReadFile(flname)
	if err != nil {
		return err
	}

	return bdl.Parse(buf)
}

// https://golang.org/pkg/bufio/#ScanLines

func (bdl *BDL) Parse(buf []byte) error {
	br := bytes.NewReader(buf)
	scanner := bufio.NewScanner(br)

	for i := 0; scanner.Scan(); i++ {
		bdl.process(scanner.Text(), i)
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		return err
	}
	WaitForAll.Wait()
	return bdl.err
}
func (bdl *BDL) process(line string, l int) {
	tokens := strings.Fields(line)
	tkn := strTok[strings.ToUpper(tokens[0])]
	if !isValidToken(bdl.st, tkn) {
		bdl.err = fmt.Errorf("unexpected  keyword [%s] in line %d", tokens[0], l)
		return
	}
	switch tkn {
	case REM:
		return

	case OBJECT:
		// OBJECT <instance name> AS <Symbol Name>
		if len(tokens) != 4 {
			bdl.err = fmt.Errorf("invalid number of tokens [%d] in line %d", len(tokens), l)
		}
		//*TODO* OBJECT <instance name> AS <Symbol Name> From <Plugin.so>

		name := tokens[1]
		as := tokens[2]
		cpath := tokens[3] //sym :=  tokens[3] ; from := tokens[4] ;  cpath := tokens[5]

		if strings.EqualFold("As", as) { // case insensitive
			bdl.err = fmt.Errorf("WARNING: Expecting As got [%s] in line %d", as, l)
			return
		}
		/*if strings.EqualFold("From", from) { // case insensitive
					bdl.err = fmt.Errorf("WARNING: Expecting From got [%s] in line %d", from, l)
					return
				}
		        bdl.parseObj(l,name, sym, cpath)*/

		bdl.parseObj(l, name, cpath)

	case WITH:
		if len(tokens) != 2 {
			bdl.err = fmt.Errorf("invalid number of tokens [%d] in line %d", len(tokens), l)
			bdl.curObjCtx = ""
			return
		}
		// WITH <instance name>
		bdl.curObjCtx = tokens[1]
		bdl.st = ST_WITH
		if bdl.objMap[bdl.curObjCtx] == nil {
			bdl.err = fmt.Errorf("invalid object instance name [%s] in line %d", registry[bdl.curObjCtx], l)
			return
		}
	case END:
		bdl.st = ST_START
		bdl.curObjCtx = ""
	case SET:
		// SET key = val
		if len(tokens) != 4 {
			bdl.err = fmt.Errorf("invalid number of tokens [%d] in line %d", len(tokens), l)
		}
		key := tokens[1]
		valStr := tokens[3]
		if strings.EqualFold("=", tokens[2]) { // case insensitive
			bdl.err = fmt.Errorf("WARNING: Expecting = got [%s] in line %d", tokens[2], l)
		}
		valParam := *bdl.objMap[valStr]
		obj := bdl.objMap[bdl.curObjCtx]
		if obj == nil {
			bdl.err = fmt.Errorf("object [%s] NOT Found, in context of SET,  [%s] at line %d", bdl.curObjCtx, key, l)
			return
		}
		// check if the interface object implements the Bean interface
		if valParam == nil {
			// https://golang.org/doc/faq#convert_slice_of_interface
			valParam = valStr

		}

		// Go Idiom: https://golang.org/pkg/reflect/#TypeOf
		beanType := reflect.TypeOf((Bean)(nil)).Elem()
		if reflect.TypeOf(*obj).Implements(beanType) {
			v := reflect.ValueOf(*obj)
			setr := v.Interface().(Bean)
			setr.Set(key, &valParam)
		} else {
			found := false
			st := reflect.TypeOf(*obj)
			for i := 0; i < st.NumField(); i++ {
				field := st.Field(i)
				if bean, ok := field.Tag.Lookup("bean"); ok {
					if bean == key {
					  found = true
					  ftyp := field.Type
					  if ftyp.AssignableTo(reflect.TypeOf(valParam)) {
					     fv := reflect.ValueOf(field)
					     fv.Set(reflect.ValueOf(valParam))
					  } else {
					     // found but not assignable
					    fldPkgPath := field.PkgPath + "/" + field.Name  
					    bdl.err = fmt.Errorf("Incompatiable Type: [%s] is unassignable to OBJECT [%s] SET  [%s] at line %d",fldPkgPath, bdl.curObjCtx, key, l)
					  }
					  break;
					}	
				}
			}

			if !found {
				bdl.err = fmt.Errorf("object [%s] does not have bean property [%s] at line %d", bdl.curObjCtx, key, l)
			}
		}

	case GO:
		run := &EmptyRun{}
		go func(runme Runnable) {
		    WaitForAll.Add(1)
			defer WaitForAll.Done()
			runme.Run()
		}(run)

	default:
		bdl.err = fmt.Errorf("invalid keyword [%s] in line %d", tokens[0], l) // should never come here. caught at isVlaidToken()
		return
	}
}
func (bdl *BDL) Object(key string) Object {
	return bdl.objMap[key]
}

type EmptyRun struct {
}

func (rn *EmptyRun) Run() {}

func (bdl *BDL) parseObj(line int, name, cpath string) {
	oFact := registry[cpath]
	if oFact == nil {
		bdl.err = fmt.Errorf("invalid class path reference [%s] in line %d", cpath, line)
		return
	}

	inst := (*oFact).CreateInstance(name)
	bdl.objMap[name] = inst

	bdl.st = ST_OBJ

}
