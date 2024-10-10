package db

import (
	"context"
	"database/sql"
	"reflect"

	gnovm "github.com/gnolang/gno/gnovm/pkg/gnolang"
	"github.com/gnolang/gno/tm2/pkg/std"
	"github.com/gnolang/gno/tm2/pkg/store"
)

// GnoDoltStore implements Store interface of gnovm/pkg/gnolang/store.go
// does not supports nested transaction
type GnoDoltStore struct {
	Param        *Tm2DoltParam
	ErrorHandler func(error) bool // if true ignore error
	db           *sql.DB
	ctx          context.Context // ideally context should not be stored, but gno store interface, does not support context
	tx           *sql.Tx
	pg           gnovm.PackageGetter
	pkgCache     map[string]*gnovm.PackageValue
	realmCache   map[string]*gnovm.Realm
}

func NewGnoDoltStore(ctx context.Context, param *Tm2DoltParam, errorHandler func(error) bool) (*GnoDoltStore, error) {
	td := &GnoDoltStore{
		Param:        param,
		ErrorHandler: errorHandler,
		ctx:          ctx,
		pkgCache:     make(map[string]*gnovm.PackageValue),
		realmCache:   make(map[string]*gnovm.Realm),
	}
	param.setDefaults()
	if err := td.open(); err != nil {
		return nil, err
	}

	go func() {
		select {
		case <-ctx.Done():
			td.Close()
			return
		}

	}()

	return td, nil
}
func (gds *GnoDoltStore) open() error {
	if gds.db != nil {
		return nil
	}
	return nil
}
func (gds *GnoDoltStore) Close() {
	if gds.tx != nil {
		gds.tx.Rollback()
		gds.tx = nil
	}
	if gds.db != nil {
		gds.db.Close()
		gds.db = nil
	}

}
func (gds *GnoDoltStore) BeginTransaction(_, _ store.Store) gnovm.TransactionStore {
	td := gds
	if gds.tx == nil {
		var err error
		if gds.tx, err = gds.db.BeginTx(gds.ctx, nil); err != nil {
			gds.ErrorHandler(err)
			td = nil
		}
	} else {
		if gd, err := NewGnoDoltStore(gds.ctx, gds.Param, gds.ErrorHandler); err == nil {
			td = gd
			if gd.tx, err = gd.db.BeginTx(gds.ctx, nil); err != nil {
				gds.ErrorHandler(err)
				td = nil
			}
		} else {
			gds.ErrorHandler(err)
			td = nil
		}

	}
	return td
}
func (gds *GnoDoltStore) ensureTxn() *sql.Tx {
	if gds.tx == nil {
		var err error
		if gds.tx, err = gds.db.BeginTx(gds.ctx, nil); err != nil {
			gds.ErrorHandler(err)
			return nil
		}
	}
	return gds.tx
}
func (gds *GnoDoltStore) SetPackageGetter(pg gnovm.PackageGetter) {
	gds.pg = pg
	return
}

func (gds *GnoDoltStore) GetPackage(pkgPath string, _ bool) *gnovm.PackageValue {
	pv, ok := gds.pkgCache[pkgPath]
	if ok {
		return pv
	}
	if gds.pg != nil {
		_, pv = gds.pg(pkgPath, gds)
	}
	if pv == nil {
		pv = gds.getPackage(pkgPath)
	}
	return pv
}

func (gds *GnoDoltStore) getPackage(pkgPath string) (pv *gnovm.PackageValue) {

	return nil
}

func (gds *GnoDoltStore) SetCachePackage(pv *gnovm.PackageValue) {
	gds.pkgCache[pv.PkgPath] = pv
	return
}

func (gds *GnoDoltStore) GetPackageRealm(pkgPath string) *gnovm.Realm {

	return nil
}
func (gds *GnoDoltStore) SetPackageRealm(*gnovm.Realm) {

	return
}
func (gds *GnoDoltStore) GetObject(oid gnovm.ObjectID) gnovm.Object {

	return nil
}
func (gds *GnoDoltStore) GetObjectSafe(oid gnovm.ObjectID) gnovm.Object {

	return nil
}
func (gds *GnoDoltStore) SetObject(gnovm.Object) {

	return
}
func (gds *GnoDoltStore) DelObject(gnovm.Object) {

	return
}
func (gds *GnoDoltStore) GetType(tid gnovm.TypeID) gnovm.Type {

	return nil
}
func (gds *GnoDoltStore) GetTypeSafe(tid gnovm.TypeID) gnovm.Type {

	return nil
}
func (gds *GnoDoltStore) SetCacheType(gnovm.Type) {

}
func (gds *GnoDoltStore) SetType(gnovm.Type) {

}
func (gds *GnoDoltStore) GetBlockNode(gnovm.Location) gnovm.BlockNode {

	return nil
}
func (gds *GnoDoltStore) GetBlockNodeSafe(gnovm.Location) gnovm.BlockNode {

	return nil
}
func (gds *GnoDoltStore) SetBlockNode(gnovm.BlockNode) {

	return
}

// UNSTABLE
func (gds *GnoDoltStore) SetStrictGo2GnoMapping(bool) {

	return
}
func (gds *GnoDoltStore) Go2GnoType(rt reflect.Type) gnovm.Type {

	return nil
}
func (gds *GnoDoltStore) GetAllocator() *gnovm.Allocator {

	return nil
}
func (gds *GnoDoltStore) NumMemPackages() int64 {

	return 0
}

// Upon restart, all packages will be re-preprocessed; This
// loads BlockNodes and Types onto the store for persistence
// version 1.
func (gds *GnoDoltStore) AddMemPackage(memPkg *std.MemPackage) {

}
func (gds *GnoDoltStore) GetMemPackage(path string) *std.MemPackage {

	return nil
}
func (gds *GnoDoltStore) GetMemFile(path string, name string) *std.MemFile {

	return nil
}
func (gds *GnoDoltStore) IterMemPackage() <-chan *std.MemPackage {

	return nil
}
func (gds *GnoDoltStore) ClearObjectCache() { // run before processing a message

	return
}
func (gds *GnoDoltStore) SetPackageInjector(gnovm.PackageInjector) {

	return
} // for natives
func (gds *GnoDoltStore) SetNativeStore(gnovm.NativeStore) {

	return
} // for "new" natives XXX
func (gds *GnoDoltStore) GetNative(pkgPath string, name gnovm.Name) func(m *gnovm.Machine) {

	return nil
} // for "new" natives XXX
func (gds *GnoDoltStore) SetLogStoreOps(enabled bool) {

	return
}
func (gds *GnoDoltStore) SprintStoreOps() string {

	return ""
}
func (gds *GnoDoltStore) LogSwitchRealm(rlmpath string) {

	return
} // to mark change of realm boundaries
func (gds *GnoDoltStore) ClearCache() {

	return
}
func (gds *GnoDoltStore) Print() {

	return
}

func (gds *GnoDoltStore) Write() {

}
