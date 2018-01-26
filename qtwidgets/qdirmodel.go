package qtwidgets

// /usr/include/qt/QtWidgets/qdirmodel.h
// #include <qdirmodel.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QDirModel struct {
	*qtcore.QAbstractItemModel
}

func (this *QDirModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QDirModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = qtcore.NewQAbstractItemModelFromPointer(cthis)
}
func NewQDirModelFromPointer(cthis unsafe.Pointer) *QDirModel {
	bcthis0 := qtcore.NewQAbstractItemModelFromPointer(cthis)
	return &QDirModel{bcthis0}
}
func (*QDirModel) NewFromPointer(cthis unsafe.Pointer) *QDirModel {
	return NewQDirModelFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDirModel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:70
// index:0
// Public
// void QDirModel(class QObject *)
func NewQDirModel(parent *qtcore.QObject /*777 QObject **/) *QDirModel {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdirmodel.h:71
// index:0
// Public virtual
// void ~QDirModel()
func DeleteQDirModel(*QDirModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:73
// index:0
// Public virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QDirModel) Index(row int, column int, parent *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:118
// index:1
// Public
// QModelIndex index(const class QString &, int)
func (this *QDirModel) Index_1(path *qtcore.QString, column int) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5indexERK7QStringi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:74
// index:0
// Public virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QDirModel) Parent(child *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:76
// index:0
// Public virtual
// int rowCount(const class QModelIndex &)
func (this *QDirModel) RowCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qdirmodel.h:77
// index:0
// Public virtual
// int columnCount(const class QModelIndex &)
func (this *QDirModel) ColumnCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qdirmodel.h:79
// index:0
// Public virtual
// QVariant data(const class QModelIndex &, int)
func (this *QDirModel) Data(index *qtcore.QModelIndex, role int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:80
// index:0
// Public virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QDirModel) SetData(index *qtcore.QModelIndex, value *qtcore.QVariant, role int) bool {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:82
// index:0
// Public virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QDirModel) HeaderData(section int, orientation int, role int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), section, orientation, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:84
// index:0
// Public virtual
// bool hasChildren(const class QModelIndex &)
func (this *QDirModel) HasChildren(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:85
// index:0
// Public virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QDirModel) Flags(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:87
// index:0
// Public virtual
// void sort(int, Qt::SortOrder)
func (this *QDirModel) Sort(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:91
// index:0
// Public virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QDirModel) DropMimeData(data *qtcore.QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *qtcore.QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:93
// index:0
// Public virtual
// Qt::DropActions supportedDropActions()
func (this *QDirModel) SupportedDropActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:97
// index:0
// Public
// void setIconProvider(class QFileIconProvider *)
func (this *QDirModel) SetIconProvider(provider *QFileIconProvider /*777 QFileIconProvider **/) {
	var convArg0 = provider.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel15setIconProviderEP17QFileIconProvider", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:98
// index:0
// Public
// QFileIconProvider * iconProvider()
func (this *QDirModel) IconProvider() *QFileIconProvider /*777 QFileIconProvider **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel12iconProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQFileIconProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:100
// index:0
// Public
// void setNameFilters(const class QStringList &)
func (this *QDirModel) SetNameFilters(filters *qtcore.QStringList) {
	var convArg0 = filters.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:104
// index:0
// Public
// QDir::Filters filter()
func (this *QDirModel) Filter() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel6filterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:106
// index:0
// Public
// void setSorting(class QDir::SortFlags)
func (this *QDirModel) SetSorting(sort int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel10setSortingE6QFlagsIN4QDir8SortFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sort)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:107
// index:0
// Public
// QDir::SortFlags sorting()
func (this *QDirModel) Sorting() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel7sortingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:109
// index:0
// Public
// void setResolveSymlinks(_Bool)
func (this *QDirModel) SetResolveSymlinks(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel18setResolveSymlinksEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:110
// index:0
// Public
// bool resolveSymlinks()
func (this *QDirModel) ResolveSymlinks() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel15resolveSymlinksEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:112
// index:0
// Public
// void setReadOnly(_Bool)
func (this *QDirModel) SetReadOnly(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:113
// index:0
// Public
// bool isReadOnly()
func (this *QDirModel) IsReadOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:115
// index:0
// Public
// void setLazyChildCount(_Bool)
func (this *QDirModel) SetLazyChildCount(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel17setLazyChildCountEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:116
// index:0
// Public
// bool lazyChildCount()
func (this *QDirModel) LazyChildCount() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel14lazyChildCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:120
// index:0
// Public
// bool isDir(const class QModelIndex &)
func (this *QDirModel) IsDir(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5isDirERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:121
// index:0
// Public
// QModelIndex mkdir(const class QModelIndex &, const class QString &)
func (this *QDirModel) Mkdir(parent *qtcore.QModelIndex, name *qtcore.QString) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = parent.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel5mkdirERK11QModelIndexRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:122
// index:0
// Public
// bool rmdir(const class QModelIndex &)
func (this *QDirModel) Rmdir(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel5rmdirERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:123
// index:0
// Public
// bool remove(const class QModelIndex &)
func (this *QDirModel) Remove(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel6removeERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:125
// index:0
// Public
// QString filePath(const class QModelIndex &)
func (this *QDirModel) FilePath(index *qtcore.QModelIndex) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8filePathERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:126
// index:0
// Public
// QString fileName(const class QModelIndex &)
func (this *QDirModel) FileName(index *qtcore.QModelIndex) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileNameERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:127
// index:0
// Public
// QIcon fileIcon(const class QModelIndex &)
func (this *QDirModel) FileIcon(index *qtcore.QModelIndex) *qtgui.QIcon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileIconERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:128
// index:0
// Public
// QFileInfo fileInfo(const class QModelIndex &)
func (this *QDirModel) FileInfo(index *qtcore.QModelIndex) *qtcore.QFileInfo /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileInfoERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:133
// index:0
// Public
// void refresh(const class QModelIndex &)
func (this *QDirModel) Refresh(parent *qtcore.QModelIndex) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel7refreshERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QDirModel__Roles = int

const QDirModel__FileIconRole QDirModel__Roles = 1
const QDirModel__FilePathRole QDirModel__Roles = 257
const QDirModel__FileNameRole QDirModel__Roles = 258

//  body block end