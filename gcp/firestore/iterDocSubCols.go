package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"qz/defs"
	"qz/util"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

/*
const (
	//defs.DocID id
	defs.DocID = "id"
	//defs.DocPATH path
	defs.DocPATH = "path"
)

// defs.DocSubCols firebase document & sub columns
type defs.DocSubCols struct {
	Doc     map[string]interface{}
	SubCols map[string]map[string]interface{}
}
*/

// IterateDocSubCols iterates the batches through handlers
type IterateDocSubCols struct {
	docIter   *firestore.DocumentIterator
	subcols   []string
	processor chan *defs.DocSubCols
	errChan   chan error
}

// NewIterateDocSubCols iterates the log via handlers
func NewIterateDocSubCols(docIter *firestore.DocumentIterator, inChan chan *defs.DocSubCols, subcols []string, errChan chan error) (bh *IterateDocSubCols) {

	bh = &IterateDocSubCols{
		docIter:   docIter,
		subcols:   subcols,
		processor: inChan,
		errChan:   errChan,
	}
	return
}

// Process  collection docs via Handlers
func (bi *IterateDocSubCols) Process(ctx context.Context) {

	//fmt.Println("Process")
	defer bi.docIter.Stop()
	defer close(bi.processor)
	for {
		select {
		case <-ctx.Done():
			util.DebugInfo(ctx, "IterateDocSubCols.Process: got ctx done")
			return
		default:
		}
		doc, err := bi.docIter.Next() // DocumentSnapshot
		if err == iterator.Done {
			break
		}
		if err != nil {
			bi.errChan <- err
			return
		}
		data := &defs.DocSubCols{
			Doc:     doc.Data(),
			SubCols: make(map[string]map[string]interface{}),
		}
		//fmt.Println("after defs.DocSubCols")
		data.Doc[defs.DocPATH], data.Doc[defs.DocID] = doc.Ref.Path, doc.Ref.ID

		for i := 0; i < len(bi.subcols); i++ {
			//fmt.Println("inside for i", i)
			colref := doc.Ref.Collection(bi.subcols[i]) // *CollectionRef
			data.SubCols[bi.subcols[i]] = make(map[string]interface{})
			data.SubCols[bi.subcols[i]][defs.DocPATH] = colref.Path
			data.SubCols[bi.subcols[i]][defs.DocID] = colref.ID
			snaps, err := colref.Documents(ctx).GetAll()
			if err != nil {
				bi.errChan <- err
				return
			}
			for j := 0; j < len(snaps); j++ {
				dmap := snaps[j].Data()
				dmap[defs.DocID] = snaps[j].Ref.ID
				dmap[defs.DocPATH] = snaps[j].Ref.Path
				// collection(<name>)map[<sub col ID>]map[string]interface{}
				data.SubCols[bi.subcols[i]][snaps[j].Ref.ID] = dmap
			} // for j
		} // for i
		bi.processor <- data
	} // for
	util.DebugInfo(ctx, "IterateDocSubCols.Process: dociter ret")
	return
}

// PrintDocSubCols prints in json format to stdio
func PrintDocSubCols(ctx context.Context, dat chan *defs.DocSubCols) error {
	return FPrintDocSubCols(ctx, os.Stdout, dat)

}

// FPrintDocSubCols prints in json format to io.Writer
// TODO: add a file handle as aparam for redirection to a file
func FPrintDocSubCols(ctx context.Context, w io.Writer, dat chan *defs.DocSubCols) error {
	firstLine := true
	fmt.Fprintln(w, "[")
	for u := range dat {
		select {
		case <-ctx.Done():
			util.DebugInfo(ctx, "defs.FPrintDocSubCols: INFO ctx done\n")
			return nil
		default:
		}
		p, err := json.MarshalIndent(u, "\n", " ")
		if err == nil {
			if !firstLine {
				fmt.Fprintln(w, ",")
			} else {
				firstLine = false
			}
			fmt.Fprintln(w, string(p))
		}
		if err != nil {
			util.DebugInfo(ctx, "firebase.FPrintDocSubCols: ret err")
			return err
		}

	}
	fmt.Fprintln(w, "]")
	util.DebugInfo(ctx, "firebase.FPrintDocSubCols: ret ok\n")
	return nil
}
