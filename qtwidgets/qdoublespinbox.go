package qtwidgets

// /usr/include/qt/QtWidgets/qspinbox.h
// #include <qspinbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QDoubleSpinBox struct {
	*QAbstractSpinBox
}
type QDoubleSpinBox_ITF interface {
	QAbstractSpinBox_ITF
	QDoubleSpinBox_PTR() *QDoubleSpinBox
}

func (ptr *QDoubleSpinBox) QDoubleSpinBox_PTR() *QDoubleSpinBox { return ptr }

func (this *QDoubleSpinBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSpinBox.GetCthis()
	}
}
func (this *QDoubleSpinBox) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSpinBox = NewQAbstractSpinBoxFromPointer(cthis)
}
func NewQDoubleSpinBoxFromPointer(cthis unsafe.Pointer) *QDoubleSpinBox {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QDoubleSpinBox{bcthis0}
}
func (*QDoubleSpinBox) NewFromPointer(cthis unsafe.Pointer) *QDoubleSpinBox {
	return NewQDoubleSpinBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qspinbox.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDoubleSpinBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qspinbox.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDoubleSpinBox(QWidget *)

/*

 */
func NewQDoubleSpinBox(parent QWidget_ITF /*777 QWidget **/) *QDoubleSpinBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDoubleSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDoubleSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDoubleSpinBox(QWidget *)

/*

 */
func NewQDoubleSpinBox__() *QDoubleSpinBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDoubleSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDoubleSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDoubleSpinBox()

/*

 */
func DeleteQDoubleSpinBox(this *QDoubleSpinBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qspinbox.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] double value() const

/*

 */
func (this *QDoubleSpinBox) Value() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox5valueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QString prefix() const

/*

 */
func (this *QDoubleSpinBox) Prefix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox6prefixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrefix(const QString &)

/*

 */
func (this *QDoubleSpinBox) SetPrefix(prefix string) {
	var tmpArg0 = qtcore.NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox9setPrefixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QString suffix() const

/*

 */
func (this *QDoubleSpinBox) Suffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox6suffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSuffix(const QString &)

/*

 */
func (this *QDoubleSpinBox) SetSuffix(suffix string) {
	var tmpArg0 = qtcore.NewQString_5(suffix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox9setSuffixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cleanText() const

/*

 */
func (this *QDoubleSpinBox) CleanText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox9cleanTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] double singleStep() const

/*

 */
func (this *QDoubleSpinBox) SingleStep() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox10singleStepEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleStep(double)

/*

 */
func (this *QDoubleSpinBox) SetSingleStep(val float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox13setSingleStepEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] double minimum() const

/*

 */
func (this *QDoubleSpinBox) Minimum() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox7minimumEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(double)

/*

 */
func (this *QDoubleSpinBox) SetMinimum(min float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox10setMinimumEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] double maximum() const

/*

 */
func (this *QDoubleSpinBox) Maximum() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox7maximumEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(double)

/*

 */
func (this *QDoubleSpinBox) SetMaximum(max float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox10setMaximumEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(double, double)

/*
Convenience function to set the minimum, and maximum values with a single function call.


  setRange(minimum, maximum);



is equivalent to:


  setMinimum(minimum);
  setMaximum(maximum);



See also minimum and maximum.
*/
func (this *QDoubleSpinBox) SetRange(min float64, max float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox8setRangeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] int decimals() const

/*

 */
func (this *QDoubleSpinBox) Decimals() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox8decimalsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDecimals(int)

/*

 */
func (this *QDoubleSpinBox) SetDecimals(prec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox11setDecimalsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), prec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:153
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &) const

/*
Reimplemented from QAbstractSpinBox::validate().
*/
func (this *QDoubleSpinBox) Validate(input string, pos int) int {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] double valueFromText(const QString &) const

/*
This virtual function is used by the spin box whenever it needs to interpret text entered by the user as a value.

Subclasses that need to display spin box values in a non-numeric way need to reimplement this function.

Note: QSpinBox handles specialValueText() separately; this function is only concerned with the other values.

See also textFromValue() and validate().
*/
func (this *QDoubleSpinBox) ValueFromText(text string) float64 {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox13valueFromTextERK7QString", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:155
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString textFromValue(double) const

/*
This virtual function is used by the spin box whenever it needs to display the given value. The default implementation returns a string containing value printed in the standard way using QWidget::locale().toString(), but with the thousand separator removed unless setGroupSeparatorShown() is set. Reimplementations may return anything. (See the example in the detailed description.)

Note: QSpinBox does not call this function for specialValueText() and that neither prefix() nor suffix() should be included in the return value.

If you reimplement this, you may also need to reimplement valueFromText() and validate()

See also valueFromText(), validate(), and QLocale::groupSeparator().
*/
func (this *QDoubleSpinBox) TextFromValue(val float64) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox13textFromValueEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:156
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &) const

/*
Reimplemented from QAbstractSpinBox::fixup().
*/
func (this *QDoubleSpinBox) Fixup(str string) {
	var tmpArg0 = qtcore.NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDoubleSpinBox5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(double)

/*

 */
func (this *QDoubleSpinBox) SetValue(val float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox8setValueEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(double)

/*
This signal is emitted whenever the spin box's value is changed. The new value's integer value is passed in i.

Note: Signal valueChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(spinBox, QOverload<int>::of(&QSpinBox::valueChanged),
      [=](int i){ \/* ... *\/ });



Note: Notifier signal for property value.
*/
func (this *QDoubleSpinBox) ValueChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox12valueChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:163
// index:1
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(const QString &)

/*
This signal is emitted whenever the spin box's value is changed. The new value's integer value is passed in i.

Note: Signal valueChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(spinBox, QOverload<int>::of(&QSpinBox::valueChanged),
      [=](int i){ \/* ... *\/ });



Note: Notifier signal for property value.
*/
func (this *QDoubleSpinBox) ValueChanged_1(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDoubleSpinBox12valueChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
