package qtwidgets

// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
type QSpacerItem struct {
	*QLayoutItem
}

func (this *QSpacerItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QLayoutItem.GetCthis()
	}
}
func (this *QSpacerItem) SetCthis(cthis unsafe.Pointer) {
	this.QLayoutItem = NewQLayoutItemFromPointer(cthis)
}
func NewQSpacerItemFromPointer(cthis unsafe.Pointer) *QSpacerItem {
	bcthis0 := NewQLayoutItemFromPointer(cthis)
	return &QSpacerItem{bcthis0}
}
func (*QSpacerItem) NewFromPointer(cthis unsafe.Pointer) *QSpacerItem {
	return NewQSpacerItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:95
// index:0
// Public inline
// void QSpacerItem(int, int, class QSizePolicy::Policy, class QSizePolicy::Policy)
func NewQSpacerItem(w int, h int, hData int, vData int) *QSpacerItem {
	cthis := qtrt.Calloc(1, 256) // 40
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItemC2EiiN11QSizePolicy6PolicyES1_", ffiqt.FFI_TYPE_VOID, cthis, w, h, hData, vData)
	gopp.ErrPrint(err, rv)
	gothis := NewQSpacerItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:99
// index:0
// Public virtual
// void ~QSpacerItem()
func DeleteQSpacerItem(*QSpacerItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:101
// index:0
// Public
// void changeSize(int, int, class QSizePolicy::Policy, class QSizePolicy::Policy)
func (this *QSpacerItem) ChangeSize(w int, h int, hData int, vData int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItem10changeSizeEiiN11QSizePolicy6PolicyES1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h, hData, vData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:104
// index:0
// Public virtual
// QSize sizeHint()
func (this *QSpacerItem) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:105
// index:0
// Public virtual
// QSize minimumSize()
func (this *QSpacerItem) MinimumSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:106
// index:0
// Public virtual
// QSize maximumSize()
func (this *QSpacerItem) MaximumSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:107
// index:0
// Public virtual
// Qt::Orientations expandingDirections()
func (this *QSpacerItem) ExpandingDirections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:108
// index:0
// Public virtual
// bool isEmpty()
func (this *QSpacerItem) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:109
// index:0
// Public virtual
// void setGeometry(const class QRect &)
func (this *QSpacerItem) SetGeometry(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItem11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:110
// index:0
// Public virtual
// QRect geometry()
func (this *QSpacerItem) Geometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem8geometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:111
// index:0
// Public virtual
// QSpacerItem * spacerItem()
func (this *QSpacerItem) SpacerItem() *QSpacerItem /*444 QSpacerItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSpacerItem10spacerItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSpacerItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:112
// index:0
// Public inline
// QSizePolicy sizePolicy()
func (this *QSpacerItem) SizePolicy() *QSizePolicy /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSpacerItem10sizePolicyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end