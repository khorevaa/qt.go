package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qdatetimeedit.h
// dst-file: /src/widgets/qdatetimeedit.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QMetaObject * QTimeEdit::metaObject();
extern void C_ZNK9QTimeEdit10metaObjectEv(void* qthis); // 4
  // proto:  void QTimeEdit::QTimeEdit(QWidget * parent);
extern void* C_ZN9QTimeEditC2EP7QWidget(void* arg0); // 3
  // proto:  void QTimeEdit::QTimeEdit(const QTime & time, QWidget * parent);
extern void* C_ZN9QTimeEditC2ERK5QTimeP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QTimeEdit::~QTimeEdit();
extern void C_ZN9QTimeEditD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QDateEdit::metaObject();
extern void C_ZNK9QDateEdit10metaObjectEv(void* qthis); // 4
  // proto:  void QDateEdit::QDateEdit(const QDate & date, QWidget * parent);
extern void* C_ZN9QDateEditC2ERK5QDateP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QDateEdit::QDateEdit(QWidget * parent);
extern void* C_ZN9QDateEditC2EP7QWidget(void* arg0); // 3
  // proto:  void QDateEdit::~QDateEdit();
extern void C_ZN9QDateEditD2Ev(void* qthis); // 4
  // proto:  void QDateTimeEdit::setMaximumDate(const QDate & max);
extern void C_ZN13QDateTimeEdit14setMaximumDateERK5QDate(void* qthis, void* arg0); // 4
  // proto:  void QDateTimeEdit::setCalendarPopup(bool enable);
extern void C_ZN13QDateTimeEdit16setCalendarPopupEb(void* qthis, bool arg0); // 4
  // proto:  Sections QDateTimeEdit::displayedSections();
extern void C_ZNK13QDateTimeEdit17displayedSectionsEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::setDateTime(const QDateTime & dateTime);
extern void C_ZN13QDateTimeEdit11setDateTimeERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  void QDateTimeEdit::setTimeRange(const QTime & min, const QTime & max);
extern void C_ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QDateTimeEdit::setDateRange(const QDate & min, const QDate & max);
extern void C_ZN13QDateTimeEdit12setDateRangeERK5QDateS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QDateTimeEdit::setMinimumTime(const QTime & min);
extern void C_ZN13QDateTimeEdit14setMinimumTimeERK5QTime(void* qthis, void* arg0); // 4
  // proto:  QString QDateTimeEdit::displayFormat();
extern void* C_ZNK13QDateTimeEdit13displayFormatEv(void* qthis); // 4
  // proto:  QDateTimeEdit::Section QDateTimeEdit::currentSection();
extern void C_ZNK13QDateTimeEdit14currentSectionEv(void* qthis); // 4
  // proto:  QTime QDateTimeEdit::maximumTime();
extern void* C_ZNK13QDateTimeEdit11maximumTimeEv(void* qthis); // 4
  // proto:  int QDateTimeEdit::currentSectionIndex();
extern int32_t C_ZNK13QDateTimeEdit19currentSectionIndexEv(void* qthis); // 4
  // proto:  bool QDateTimeEdit::event(QEvent * event);
extern bool C_ZN13QDateTimeEdit5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  QDate QDateTimeEdit::minimumDate();
extern void* C_ZNK13QDateTimeEdit11minimumDateEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::setDateTimeRange(const QDateTime & min, const QDateTime & max);
extern void C_ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QDateTimeEdit::clearMinimumDate();
extern void C_ZN13QDateTimeEdit16clearMinimumDateEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::clearMaximumDate();
extern void C_ZN13QDateTimeEdit16clearMaximumDateEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::setDisplayFormat(const QString & format);
extern void C_ZN13QDateTimeEdit16setDisplayFormatERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QDateTimeEdit::setMinimumDateTime(const QDateTime & dt);
extern void C_ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  void QDateTimeEdit::setMaximumTime(const QTime & max);
extern void C_ZN13QDateTimeEdit14setMaximumTimeERK5QTime(void* qthis, void* arg0); // 4
  // proto:  void QDateTimeEdit::~QDateTimeEdit();
extern void C_ZN13QDateTimeEditD2Ev(void* qthis); // 4
  // proto:  void QDateTimeEdit::setCalendarWidget(QCalendarWidget * calendarWidget);
extern void C_ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget(void* qthis, void* arg0); // 4
  // proto:  void QDateTimeEdit::clearMinimumTime();
extern void C_ZN13QDateTimeEdit16clearMinimumTimeEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::setDate(const QDate & date);
extern void C_ZN13QDateTimeEdit7setDateERK5QDate(void* qthis, void* arg0); // 4
  // proto:  bool QDateTimeEdit::calendarPopup();
extern bool C_ZNK13QDateTimeEdit13calendarPopupEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::stepBy(int steps);
extern void C_ZN13QDateTimeEdit6stepByEi(void* qthis, int32_t arg0); // 4
  // proto:  QCalendarWidget * QDateTimeEdit::calendarWidget();
extern void* C_ZNK13QDateTimeEdit14calendarWidgetEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::clearMinimumDateTime();
extern void C_ZN13QDateTimeEdit20clearMinimumDateTimeEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::setMaximumDateTime(const QDateTime & dt);
extern void C_ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  void QDateTimeEdit::setMinimumDate(const QDate & min);
extern void C_ZN13QDateTimeEdit14setMinimumDateERK5QDate(void* qthis, void* arg0); // 4
  // proto:  int QDateTimeEdit::sectionCount();
extern int32_t C_ZNK13QDateTimeEdit12sectionCountEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::setTime(const QTime & time);
extern void C_ZN13QDateTimeEdit7setTimeERK5QTime(void* qthis, void* arg0); // 4
  // proto:  void QDateTimeEdit::clearMaximumDateTime();
extern void C_ZN13QDateTimeEdit20clearMaximumDateTimeEv(void* qthis); // 4
  // proto:  QDate QDateTimeEdit::date();
extern void* C_ZNK13QDateTimeEdit4dateEv(void* qthis); // 4
  // proto:  QSize QDateTimeEdit::sizeHint();
extern void* C_ZNK13QDateTimeEdit8sizeHintEv(void* qthis); // 4
  // proto:  QDate QDateTimeEdit::maximumDate();
extern void* C_ZNK13QDateTimeEdit11maximumDateEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::QDateTimeEdit(const QTime & t, QWidget * parent);
extern void* C_ZN13QDateTimeEditC2ERK5QTimeP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QDateTimeEdit::QDateTimeEdit(const QDate & d, QWidget * parent);
extern void* C_ZN13QDateTimeEditC2ERK5QDateP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QDateTimeEdit::QDateTimeEdit(const QDateTime & dt, QWidget * parent);
extern void* C_ZN13QDateTimeEditC2ERK9QDateTimeP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QDateTimeEdit::QDateTimeEdit(QWidget * parent);
extern void* C_ZN13QDateTimeEditC2EP7QWidget(void* arg0); // 3
  // proto:  const QMetaObject * QDateTimeEdit::metaObject();
extern void C_ZNK13QDateTimeEdit10metaObjectEv(void* qthis); // 4
  // proto:  Qt::TimeSpec QDateTimeEdit::timeSpec();
extern void C_ZNK13QDateTimeEdit8timeSpecEv(void* qthis); // 4
  // proto:  QDateTime QDateTimeEdit::dateTime();
extern void* C_ZNK13QDateTimeEdit8dateTimeEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::clear();
extern void C_ZN13QDateTimeEdit5clearEv(void* qthis); // 4
  // proto:  QDateTimeEdit::Section QDateTimeEdit::sectionAt(int index);
extern void C_ZNK13QDateTimeEdit9sectionAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QDateTimeEdit::clearMaximumTime();
extern void C_ZN13QDateTimeEdit16clearMaximumTimeEv(void* qthis); // 4
  // proto:  QTime QDateTimeEdit::minimumTime();
extern void* C_ZNK13QDateTimeEdit11minimumTimeEv(void* qthis); // 4
  // proto:  QDateTime QDateTimeEdit::maximumDateTime();
extern void* C_ZNK13QDateTimeEdit15maximumDateTimeEv(void* qthis); // 4
  // proto:  QTime QDateTimeEdit::time();
extern void* C_ZNK13QDateTimeEdit4timeEv(void* qthis); // 4
  // proto:  QDateTime QDateTimeEdit::minimumDateTime();
extern void* C_ZNK13QDateTimeEdit15minimumDateTimeEv(void* qthis); // 4
  // proto:  void QDateTimeEdit::setCurrentSectionIndex(int index);
extern void C_ZN13QDateTimeEdit22setCurrentSectionIndexEi(void* qthis, int32_t arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTimeEdit)=1
type QTimeEdit struct {
  /*qbase*/ QDateTimeEdit;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _userTimeChanged QTimeEdit_userTimeChanged_signal;
}

// class sizeof(QDateEdit)=1
type QDateEdit struct {
  /*qbase*/ QDateTimeEdit;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _userDateChanged QDateEdit_userDateChanged_signal;
}

// class sizeof(QDateTimeEdit)=1
type QDateTimeEdit struct {
  /*qbase*/ QAbstractSpinBox;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _dateChanged QDateTimeEdit_dateChanged_signal;
//  _timeChanged QDateTimeEdit_timeChanged_signal;
//  _dateTimeChanged QDateTimeEdit_dateTimeChanged_signal;
}

// metaObject()
func (this *QTimeEdit) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QTimeEdit10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTimeEdit", "metaObject", args)
  }

  return
}

// QTimeEdit(class QWidget *)
func NewQTimeEdit(args ...interface{}) *QTimeEdit {
  // QTimeEdit(class QWidget *)
  // QTimeEdit(const class QTime &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeEditC1EP7QWidget
    // invoke: void QTimeEdit(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTimeEditC2EP7QWidget(arg0)
    return &QTimeEdit{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QTimeEditC1ERK5QTimeP7QWidget
    // invoke: void QTimeEdit(const class QTime &, class QWidget *)
    var arg0 = args[0].(*QTime).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTimeEditC2ERK5QTimeP7QWidget(arg0, arg1)
    return &QTimeEdit{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTimeEdit", "QTimeEdit", args)
  }

  return nil // QTimeEdit{Qclsinst:qthis}
}

// ~QTimeEdit()
func (this *QTimeEdit) Freeqtimeedit(args ...interface{}) () {
  // ~QTimeEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeEditD0Ev
    // invoke: void ~QTimeEdit()
    C.C_ZN9QTimeEditD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTimeEdit", "~QTimeEdit", args)
  }

  return
}

// metaObject()
func (this *QDateEdit) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QDateEdit10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateEdit", "metaObject", args)
  }

  return
}

// QDateEdit(const class QDate &, class QWidget *)
func NewQDateEdit(args ...interface{}) *QDateEdit {
  // QDateEdit(const class QDate &, class QWidget *)
  // QDateEdit(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateEditC1ERK5QDateP7QWidget
    // invoke: void QDateEdit(const class QDate &, class QWidget *)
    var arg0 = args[0].(*QDate).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QDateEditC2ERK5QDateP7QWidget(arg0, arg1)
    return &QDateEdit{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QDateEditC1EP7QWidget
    // invoke: void QDateEdit(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QDateEditC2EP7QWidget(arg0)
    return &QDateEdit{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDateEdit", "QDateEdit", args)
  }

  return nil // QDateEdit{Qclsinst:qthis}
}

// ~QDateEdit()
func (this *QDateEdit) Freeqdateedit(args ...interface{}) () {
  // ~QDateEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateEditD0Ev
    // invoke: void ~QDateEdit()
    C.C_ZN9QDateEditD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateEdit", "~QDateEdit", args)
  }

  return
}

// setMaximumDate(const class QDate &)
func (this *QDateTimeEdit) Setmaximumdate(args ...interface{}) () {
  // setMaximumDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit14setMaximumDateERK5QDate
    // invoke: void setMaximumDate(const class QDate &)
    var arg0 = args[0].(*QDate).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit14setMaximumDateERK5QDate(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumDate", args)
  }

  return
}

// setCalendarPopup(_Bool)
func (this *QDateTimeEdit) Setcalendarpopup(args ...interface{}) () {
  // setCalendarPopup(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16setCalendarPopupEb
    // invoke: void setCalendarPopup(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit16setCalendarPopupEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCalendarPopup", args)
  }

  return
}

// displayedSections()
func (this *QDateTimeEdit) Displayedsections(args ...interface{}) () {
  // displayedSections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit17displayedSectionsEv
    // invoke: Sections displayedSections()
    C.C_ZNK13QDateTimeEdit17displayedSectionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "displayedSections", args)
  }

  return
}

// setDateTime(const class QDateTime &)
func (this *QDateTimeEdit) Setdatetime(args ...interface{}) () {
  // setDateTime(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit11setDateTimeERK9QDateTime
    // invoke: void setDateTime(const class QDateTime &)
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit11setDateTimeERK9QDateTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateTime", args)
  }

  return
}

// setTimeRange(const class QTime &, const class QTime &)
func (this *QDateTimeEdit) Settimerange(args ...interface{}) () {
  // setTimeRange(const class QTime &, const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[0][1] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_
    // invoke: void setTimeRange(const class QTime &, const class QTime &)
    var arg0 = args[0].(*QTime).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTime).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setTimeRange", args)
  }

  return
}

// setDateRange(const class QDate &, const class QDate &)
func (this *QDateTimeEdit) Setdaterange(args ...interface{}) () {
  // setDateRange(const class QDate &, const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[0][1] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit12setDateRangeERK5QDateS2_
    // invoke: void setDateRange(const class QDate &, const class QDate &)
    var arg0 = args[0].(*QDate).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QDate).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QDateTimeEdit12setDateRangeERK5QDateS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateRange", args)
  }

  return
}

// setMinimumTime(const class QTime &)
func (this *QDateTimeEdit) Setminimumtime(args ...interface{}) () {
  // setMinimumTime(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit14setMinimumTimeERK5QTime
    // invoke: void setMinimumTime(const class QTime &)
    var arg0 = args[0].(*QTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit14setMinimumTimeERK5QTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumTime", args)
  }

  return
}

// displayFormat()
func (this *QDateTimeEdit) Displayformat(args ...interface{}) (ret interface{}) {
  // displayFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit13displayFormatEv
    // invoke: QString displayFormat()
    var ret0 = C.C_ZNK13QDateTimeEdit13displayFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "displayFormat", args)
  }

  return
}

// currentSection()
func (this *QDateTimeEdit) Currentsection(args ...interface{}) () {
  // currentSection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit14currentSectionEv
    // invoke: QDateTimeEdit::Section currentSection()
    C.C_ZNK13QDateTimeEdit14currentSectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "currentSection", args)
  }

  return
}

// maximumTime()
func (this *QDateTimeEdit) Maximumtime(args ...interface{}) (ret interface{}) {
  // maximumTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit11maximumTimeEv
    // invoke: QTime maximumTime()
    var ret0 = C.C_ZNK13QDateTimeEdit11maximumTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumTime", args)
  }

  return
}

// currentSectionIndex()
func (this *QDateTimeEdit) Currentsectionindex(args ...interface{}) (ret interface{}) {
  // currentSectionIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit19currentSectionIndexEv
    // invoke: int currentSectionIndex()
    var ret0 = C.C_ZNK13QDateTimeEdit19currentSectionIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "currentSectionIndex", args)
  }

  return
}

// event(class QEvent *)
func (this *QDateTimeEdit) Event(args ...interface{}) (ret interface{}) {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(*QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QDateTimeEdit5eventEP6QEvent(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "event", args)
  }

  return
}

// minimumDate()
func (this *QDateTimeEdit) Minimumdate(args ...interface{}) (ret interface{}) {
  // minimumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit11minimumDateEv
    // invoke: QDate minimumDate()
    var ret0 = C.C_ZNK13QDateTimeEdit11minimumDateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumDate", args)
  }

  return
}

// setDateTimeRange(const class QDateTime &, const class QDateTime &)
func (this *QDateTimeEdit) Setdatetimerange(args ...interface{}) () {
  // setDateTimeRange(const class QDateTime &, const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[0][1] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_
    // invoke: void setDateTimeRange(const class QDateTime &, const class QDateTime &)
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QDateTime).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateTimeRange", args)
  }

  return
}

// clearMinimumDate()
func (this *QDateTimeEdit) Clearminimumdate(args ...interface{}) () {
  // clearMinimumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16clearMinimumDateEv
    // invoke: void clearMinimumDate()
    C.C_ZN13QDateTimeEdit16clearMinimumDateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumDate", args)
  }

  return
}

// clearMaximumDate()
func (this *QDateTimeEdit) Clearmaximumdate(args ...interface{}) () {
  // clearMaximumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16clearMaximumDateEv
    // invoke: void clearMaximumDate()
    C.C_ZN13QDateTimeEdit16clearMaximumDateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumDate", args)
  }

  return
}

// setDisplayFormat(const class QString &)
func (this *QDateTimeEdit) Setdisplayformat(args ...interface{}) () {
  // setDisplayFormat(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16setDisplayFormatERK7QString
    // invoke: void setDisplayFormat(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit16setDisplayFormatERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDisplayFormat", args)
  }

  return
}

// setMinimumDateTime(const class QDateTime &)
func (this *QDateTimeEdit) Setminimumdatetime(args ...interface{}) () {
  // setMinimumDateTime(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime
    // invoke: void setMinimumDateTime(const class QDateTime &)
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumDateTime", args)
  }

  return
}

// setMaximumTime(const class QTime &)
func (this *QDateTimeEdit) Setmaximumtime(args ...interface{}) () {
  // setMaximumTime(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit14setMaximumTimeERK5QTime
    // invoke: void setMaximumTime(const class QTime &)
    var arg0 = args[0].(*QTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit14setMaximumTimeERK5QTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumTime", args)
  }

  return
}

// ~QDateTimeEdit()
func (this *QDateTimeEdit) Freeqdatetimeedit(args ...interface{}) () {
  // ~QDateTimeEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEditD0Ev
    // invoke: void ~QDateTimeEdit()
    C.C_ZN13QDateTimeEditD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "~QDateTimeEdit", args)
  }

  return
}

// setCalendarWidget(class QCalendarWidget *)
func (this *QDateTimeEdit) Setcalendarwidget(args ...interface{}) () {
  // setCalendarWidget(class QCalendarWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCalendarWidget{}) // "QCalendarWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget
    // invoke: void setCalendarWidget(class QCalendarWidget *)
    var arg0 = args[0].(*QCalendarWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCalendarWidget", args)
  }

  return
}

// clearMinimumTime()
func (this *QDateTimeEdit) Clearminimumtime(args ...interface{}) () {
  // clearMinimumTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16clearMinimumTimeEv
    // invoke: void clearMinimumTime()
    C.C_ZN13QDateTimeEdit16clearMinimumTimeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumTime", args)
  }

  return
}

// setDate(const class QDate &)
func (this *QDateTimeEdit) Setdate(args ...interface{}) () {
  // setDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit7setDateERK5QDate
    // invoke: void setDate(const class QDate &)
    var arg0 = args[0].(*QDate).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit7setDateERK5QDate(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDate", args)
  }

  return
}

// calendarPopup()
func (this *QDateTimeEdit) Calendarpopup(args ...interface{}) (ret interface{}) {
  // calendarPopup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit13calendarPopupEv
    // invoke: bool calendarPopup()
    var ret0 = C.C_ZNK13QDateTimeEdit13calendarPopupEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "calendarPopup", args)
  }

  return
}

// stepBy(int)
func (this *QDateTimeEdit) Stepby(args ...interface{}) () {
  // stepBy(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit6stepByEi
    // invoke: void stepBy(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit6stepByEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "stepBy", args)
  }

  return
}

// calendarWidget()
func (this *QDateTimeEdit) Calendarwidget(args ...interface{}) (ret interface{}) {
  // calendarWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit14calendarWidgetEv
    // invoke: QCalendarWidget * calendarWidget()
    var ret0 = C.C_ZNK13QDateTimeEdit14calendarWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCalendarWidget{}) // "QCalendarWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "calendarWidget", args)
  }

  return
}

// clearMinimumDateTime()
func (this *QDateTimeEdit) Clearminimumdatetime(args ...interface{}) () {
  // clearMinimumDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit20clearMinimumDateTimeEv
    // invoke: void clearMinimumDateTime()
    C.C_ZN13QDateTimeEdit20clearMinimumDateTimeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumDateTime", args)
  }

  return
}

// setMaximumDateTime(const class QDateTime &)
func (this *QDateTimeEdit) Setmaximumdatetime(args ...interface{}) () {
  // setMaximumDateTime(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime
    // invoke: void setMaximumDateTime(const class QDateTime &)
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumDateTime", args)
  }

  return
}

// setMinimumDate(const class QDate &)
func (this *QDateTimeEdit) Setminimumdate(args ...interface{}) () {
  // setMinimumDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit14setMinimumDateERK5QDate
    // invoke: void setMinimumDate(const class QDate &)
    var arg0 = args[0].(*QDate).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit14setMinimumDateERK5QDate(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumDate", args)
  }

  return
}

// sectionCount()
func (this *QDateTimeEdit) Sectioncount(args ...interface{}) (ret interface{}) {
  // sectionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit12sectionCountEv
    // invoke: int sectionCount()
    var ret0 = C.C_ZNK13QDateTimeEdit12sectionCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "sectionCount", args)
  }

  return
}

// setTime(const class QTime &)
func (this *QDateTimeEdit) Settime(args ...interface{}) () {
  // setTime(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit7setTimeERK5QTime
    // invoke: void setTime(const class QTime &)
    var arg0 = args[0].(*QTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit7setTimeERK5QTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setTime", args)
  }

  return
}

// clearMaximumDateTime()
func (this *QDateTimeEdit) Clearmaximumdatetime(args ...interface{}) () {
  // clearMaximumDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit20clearMaximumDateTimeEv
    // invoke: void clearMaximumDateTime()
    C.C_ZN13QDateTimeEdit20clearMaximumDateTimeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumDateTime", args)
  }

  return
}

// date()
func (this *QDateTimeEdit) Date(args ...interface{}) (ret interface{}) {
  // date()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit4dateEv
    // invoke: QDate date()
    var ret0 = C.C_ZNK13QDateTimeEdit4dateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "date", args)
  }

  return
}

// sizeHint()
func (this *QDateTimeEdit) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK13QDateTimeEdit8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "sizeHint", args)
  }

  return
}

// maximumDate()
func (this *QDateTimeEdit) Maximumdate(args ...interface{}) (ret interface{}) {
  // maximumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit11maximumDateEv
    // invoke: QDate maximumDate()
    var ret0 = C.C_ZNK13QDateTimeEdit11maximumDateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumDate", args)
  }

  return
}

// QDateTimeEdit(const class QTime &, class QWidget *)
func NewQDateTimeEdit(args ...interface{}) *QDateTimeEdit {
  // QDateTimeEdit(const class QTime &, class QWidget *)
  // QDateTimeEdit(const class QDate &, class QWidget *)
  // QDateTimeEdit(const class QDateTime &, class QWidget *)
  // QDateTimeEdit(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[2][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEditC1ERK5QTimeP7QWidget
    // invoke: void QDateTimeEdit(const class QTime &, class QWidget *)
    var arg0 = args[0].(*QTime).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QDateTimeEditC2ERK5QTimeP7QWidget(arg0, arg1)
    return &QDateTimeEdit{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QDateTimeEditC1ERK5QDateP7QWidget
    // invoke: void QDateTimeEdit(const class QDate &, class QWidget *)
    var arg0 = args[0].(*QDate).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QDateTimeEditC2ERK5QDateP7QWidget(arg0, arg1)
    return &QDateTimeEdit{Qclsinst:qthis}
  case 2:
    // invoke: _ZN13QDateTimeEditC1ERK9QDateTimeP7QWidget
    // invoke: void QDateTimeEdit(const class QDateTime &, class QWidget *)
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QDateTimeEditC2ERK9QDateTimeP7QWidget(arg0, arg1)
    return &QDateTimeEdit{Qclsinst:qthis}
  case 3:
    // invoke: _ZN13QDateTimeEditC1EP7QWidget
    // invoke: void QDateTimeEdit(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QDateTimeEditC2EP7QWidget(arg0)
    return &QDateTimeEdit{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "QDateTimeEdit", args)
  }

  return nil // QDateTimeEdit{Qclsinst:qthis}
}

// metaObject()
func (this *QDateTimeEdit) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QDateTimeEdit10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "metaObject", args)
  }

  return
}

// timeSpec()
func (this *QDateTimeEdit) Timespec(args ...interface{}) () {
  // timeSpec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit8timeSpecEv
    // invoke: Qt::TimeSpec timeSpec()
    C.C_ZNK13QDateTimeEdit8timeSpecEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "timeSpec", args)
  }

  return
}

// dateTime()
func (this *QDateTimeEdit) Datetime(args ...interface{}) (ret interface{}) {
  // dateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit8dateTimeEv
    // invoke: QDateTime dateTime()
    var ret0 = C.C_ZNK13QDateTimeEdit8dateTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "dateTime", args)
  }

  return
}

// clear()
func (this *QDateTimeEdit) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit5clearEv
    // invoke: void clear()
    C.C_ZN13QDateTimeEdit5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clear", args)
  }

  return
}

// sectionAt(int)
func (this *QDateTimeEdit) Sectionat(args ...interface{}) () {
  // sectionAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit9sectionAtEi
    // invoke: QDateTimeEdit::Section sectionAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK13QDateTimeEdit9sectionAtEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "sectionAt", args)
  }

  return
}

// clearMaximumTime()
func (this *QDateTimeEdit) Clearmaximumtime(args ...interface{}) () {
  // clearMaximumTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16clearMaximumTimeEv
    // invoke: void clearMaximumTime()
    C.C_ZN13QDateTimeEdit16clearMaximumTimeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumTime", args)
  }

  return
}

// minimumTime()
func (this *QDateTimeEdit) Minimumtime(args ...interface{}) (ret interface{}) {
  // minimumTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit11minimumTimeEv
    // invoke: QTime minimumTime()
    var ret0 = C.C_ZNK13QDateTimeEdit11minimumTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumTime", args)
  }

  return
}

// maximumDateTime()
func (this *QDateTimeEdit) Maximumdatetime(args ...interface{}) (ret interface{}) {
  // maximumDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit15maximumDateTimeEv
    // invoke: QDateTime maximumDateTime()
    var ret0 = C.C_ZNK13QDateTimeEdit15maximumDateTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumDateTime", args)
  }

  return
}

// time()
func (this *QDateTimeEdit) Time(args ...interface{}) (ret interface{}) {
  // time()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit4timeEv
    // invoke: QTime time()
    var ret0 = C.C_ZNK13QDateTimeEdit4timeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTime{}) // "QTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "time", args)
  }

  return
}

// minimumDateTime()
func (this *QDateTimeEdit) Minimumdatetime(args ...interface{}) (ret interface{}) {
  // minimumDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit15minimumDateTimeEv
    // invoke: QDateTime minimumDateTime()
    var ret0 = C.C_ZNK13QDateTimeEdit15minimumDateTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDateTime{}) // "QDateTime"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumDateTime", args)
  }

  return
}

// setCurrentSectionIndex(int)
func (this *QDateTimeEdit) Setcurrentsectionindex(args ...interface{}) () {
  // setCurrentSectionIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit22setCurrentSectionIndexEi
    // invoke: void setCurrentSectionIndex(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QDateTimeEdit22setCurrentSectionIndexEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCurrentSectionIndex", args)
  }

  return
}

// <= body block end

