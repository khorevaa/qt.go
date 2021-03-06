package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QGradient struct {
	*qtrt.CObject
}
type QGradient_ITF interface {
	QGradient_PTR() *QGradient
}

func (ptr *QGradient) QGradient_PTR() *QGradient { return ptr }

func (this *QGradient) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGradient) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGradientFromPointer(cthis unsafe.Pointer) *QGradient {
	return &QGradient{&qtrt.CObject{cthis}}
}
func (*QGradient) NewFromPointer(cthis unsafe.Pointer) *QGradient {
	return NewQGradientFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGradient()

/*

 */
func NewQGradient() *QGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradientC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QGradient::Type type() const

/*

 */
func (this *QGradient) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradient4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:210
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSpread(QGradient::Spread)

/*

 */
func (this *QGradient) SetSpread(spread int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradient9setSpreadENS_6SpreadE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spread)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:211
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QGradient::Spread spread() const

/*

 */
func (this *QGradient) Spread() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradient6spreadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorAt(qreal, const QColor &)

/*

 */
func (this *QGradient) SetColorAt(pos float64, color QColor_ITF) {
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradient10setColorAtEdRK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:218
// index:0
// Public Visibility=Default Availability=Available
// [4] QGradient::CoordinateMode coordinateMode() const

/*

 */
func (this *QGradient) CoordinateMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradient14coordinateModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCoordinateMode(QGradient::CoordinateMode)

/*

 */
func (this *QGradient) SetCoordinateMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradient17setCoordinateModeENS_14CoordinateModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:221
// index:0
// Public Visibility=Default Availability=Available
// [4] QGradient::InterpolationMode interpolationMode() const

/*

 */
func (this *QGradient) InterpolationMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradient17interpolationModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qbrush.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInterpolationMode(QGradient::InterpolationMode)

/*

 */
func (this *QGradient) SetInterpolationMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradient20setInterpolationModeENS_17InterpolationModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:224
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QGradient &) const

/*

 */
func (this *QGradient) Operator_equal_equal(gradient QGradient_ITF) bool {
	var convArg0 unsafe.Pointer
	if gradient != nil && gradient.QGradient_PTR() != nil {
		convArg0 = gradient.QGradient_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradienteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qbrush.h:225
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QGradient &) const

/*

 */
func (this *QGradient) Operator_not_equal(other QGradient_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGradient_PTR() != nil {
		convArg0 = other.QGradient_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGradientneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQGradient(this *QGradient) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGradientD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QGradient__Type = int

//
const QGradient__LinearGradient QGradient__Type = 0

//
const QGradient__RadialGradient QGradient__Type = 1

//
const QGradient__ConicalGradient QGradient__Type = 2

//
const QGradient__NoGradient QGradient__Type = 3

func (this *QGradient) TypeItemName(val int) string {
	switch val {
	case QGradient__LinearGradient: // 0
		return "LinearGradient"
	case QGradient__RadialGradient: // 1
		return "RadialGradient"
	case QGradient__ConicalGradient: // 2
		return "ConicalGradient"
	case QGradient__NoGradient: // 3
		return "NoGradient"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_TypeItemName(val int) string {
	var nilthis *QGradient
	return nilthis.TypeItemName(val)
}

/*


 */
type QGradient__Spread = int

//
const QGradient__PadSpread QGradient__Spread = 0

//
const QGradient__ReflectSpread QGradient__Spread = 1

//
const QGradient__RepeatSpread QGradient__Spread = 2

func (this *QGradient) SpreadItemName(val int) string {
	switch val {
	case QGradient__PadSpread: // 0
		return "PadSpread"
	case QGradient__ReflectSpread: // 1
		return "ReflectSpread"
	case QGradient__RepeatSpread: // 2
		return "RepeatSpread"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_SpreadItemName(val int) string {
	var nilthis *QGradient
	return nilthis.SpreadItemName(val)
}

/*


 */
type QGradient__CoordinateMode = int

//
const QGradient__LogicalMode QGradient__CoordinateMode = 0

//
const QGradient__StretchToDeviceMode QGradient__CoordinateMode = 1

//
const QGradient__ObjectBoundingMode QGradient__CoordinateMode = 2

func (this *QGradient) CoordinateModeItemName(val int) string {
	switch val {
	case QGradient__LogicalMode: // 0
		return "LogicalMode"
	case QGradient__StretchToDeviceMode: // 1
		return "StretchToDeviceMode"
	case QGradient__ObjectBoundingMode: // 2
		return "ObjectBoundingMode"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_CoordinateModeItemName(val int) string {
	var nilthis *QGradient
	return nilthis.CoordinateModeItemName(val)
}

/*


 */
type QGradient__InterpolationMode = int

//
const QGradient__ColorInterpolation QGradient__InterpolationMode = 0

//
const QGradient__ComponentInterpolation QGradient__InterpolationMode = 1

func (this *QGradient) InterpolationModeItemName(val int) string {
	switch val {
	case QGradient__ColorInterpolation: // 0
		return "ColorInterpolation"
	case QGradient__ComponentInterpolation: // 1
		return "ComponentInterpolation"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGradient_InterpolationModeItemName(val int) string {
	var nilthis *QGradient
	return nilthis.InterpolationModeItemName(val)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
