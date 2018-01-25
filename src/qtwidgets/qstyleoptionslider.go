package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QStyleOptionSlider struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionSlider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionSlider) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionSliderFromPointer(cthis unsafe.Pointer) *QStyleOptionSlider {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSlider{bcthis0}
}
func (*QStyleOptionSlider) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionSlider {
	return NewQStyleOptionSliderFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:533
// index:0
// Public
// void QStyleOptionSlider()
func NewQStyleOptionSlider() *QStyleOptionSlider {
	cthis := qtrt.Calloc(1, 256) // 128
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStyleOptionSliderC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSliderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:537
// index:1
// Protected
// void QStyleOptionSlider(int)
func NewQStyleOptionSlider_1(version int) *QStyleOptionSlider {
	cthis := qtrt.Calloc(1, 256) // 128
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStyleOptionSliderC2Ei", ffiqt.FFI_TYPE_VOID, cthis, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSliderFromPointer(cthis)
	return gothis
}

type QStyleOptionSlider__StyleOptionType = int

const QStyleOptionSlider__Type QStyleOptionSlider__StyleOptionType = 983041

type QStyleOptionSlider__StyleOptionVersion = int

const QStyleOptionSlider__Version QStyleOptionSlider__StyleOptionVersion = 1

//  body block end