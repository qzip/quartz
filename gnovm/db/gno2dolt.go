package db

import (
	"database/sql"
	"reflect"

	gnovm "github.com/gnolang/gno/gnovm/pkg/gnolang"
	"github.com/gnolang/gno/tm2/pkg/std"
	"github.com/gnolang/gno/tm2/pkg/store"
)

// GnoDoltStore implements Store interface of gnovm/pkg/gnolang/store.go
type GnoDoltStore struct {
	Param        *Tm2DoltParam
	ErrorHandler func(error) bool // if true ignore error
	db           *sql.DB
}

func NewGnoDoltStore(param *Tm2DoltParam, errorHandler func(error) bool) (*GnoDoltStore, error) {
	td := &GnoDoltStore{
		Param:        param,
		ErrorHandler: errorHandler,
	}
	param.setDefaults()
	if err := td.open(); err != nil {
		return nil, err
	}
	return td, nil
}
func (gds *GnoDoltStore) open() error {

	return nil
}
func (gds *GnoDoltStore) BeginTransaction(_, _ store.Store) gnovm.TransactionStore {

	return nil
}

func (gds *GnoDoltStore) SetPackageGetter(pg gnovm.PackageGetter) {

	return
}

func (gds *GnoDoltStore) GetPackage(pkgPath string, isImport bool) *gnovm.PackageValue {

	return nil
}

func (gds *GnoDoltStore) SetCachePackage(*gnovm.PackageValue) {

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
