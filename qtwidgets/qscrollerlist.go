package qtwidgets

// /usr/include/qt/QtWidgets/qscroller.h
// #include <qscroller.h>
// #include <QtWidgets>

//  header block end

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
import "github.com/kitech/qt.go/qtgui"

//  ext block end

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

//  body block begin
type QScrollerList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QScrollerList) Operator_equal_0() *QScrollerList {
	// QScrollerList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QScrollerList) Operator_equal_1() *QScrollerList {
	// QScrollerList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QScrollerList) Swap_0() {
	// QScrollerList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QScrollerList) Operator_equal_equal_0() bool {
	// QScrollerList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QScrollerList) Operator_not_equal_0() bool {
	// QScrollerList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QScrollerList) Size_0() int {
	// QScrollerList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QScrollerList) Detach_0() {
	// QScrollerList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QScrollerList) DetachShared_0() {
	// QScrollerList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QScrollerList) IsDetached_0() bool {
	// QScrollerList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QScrollerList) SetSharable_0() {
	// QScrollerList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QScrollerList) IsSharedWith_0() bool {
	// QScrollerList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QScrollerList) IsEmpty_0() bool {
	// QScrollerList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QScrollerList) Clear_0() {
	// QScrollerList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QScrollerList) At_0() *QScroller {
	// QScrollerList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// const T & operator[](int)
func (this *QScrollerList) Operator_get_index_0() *QScroller {
	// QScrollerList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// T & operator[](int)
func (this *QScrollerList) Operator_get_index_1() *QScroller {
	// QScrollerList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// void reserve(int)
func (this *QScrollerList) Reserve_0() {
	// QScrollerList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QScrollerList) Append_0() {
	// QScrollerList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QScrollerList) Append_1() {
	// QScrollerList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QScrollerList) Prepend_0() {
	// QScrollerList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QScrollerList) Insert_0() {
	// QScrollerList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QScrollerList) Replace_0() {
	// QScrollerList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QScrollerList) RemoveAt_0() {
	// QScrollerList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QScrollerList) RemoveAll_0() int {
	// QScrollerList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QScrollerList) RemoveOne_0() bool {
	// QScrollerList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QScrollerList) TakeAt_0() *QScroller {
	// QScrollerList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// T takeFirst()
func (this *QScrollerList) TakeFirst_0() *QScroller {
	// QScrollerList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// T takeLast()
func (this *QScrollerList) TakeLast_0() *QScroller {
	// QScrollerList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// void move(int, int)
func (this *QScrollerList) Move_0() {
	// QScrollerList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QScrollerList) Swap_1() {
	// QScrollerList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QScrollerList) IndexOf_0() int {
	// QScrollerList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QScrollerList) LastIndexOf_0() int {
	// QScrollerList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QScrollerList) Contains_0() bool {
	// QScrollerList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QScrollerList) Count_0() int {
	// QScrollerList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QScrollerList) Begin_0() {
	// QScrollerList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QScrollerList) Begin_1() {
	// QScrollerList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QScrollerList) Cbegin_0() {
	// QScrollerList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QScrollerList) ConstBegin_0() {
	// QScrollerList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QScrollerList) End_0() {
	// QScrollerList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QScrollerList) End_1() {
	// QScrollerList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QScrollerList) Cend_0() {
	// QScrollerList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QScrollerList) ConstEnd_0() {
	// QScrollerList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QScrollerList) Rbegin_0() {
	// QScrollerList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QScrollerList) Rend_0() {
	// QScrollerList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QScrollerList) Rbegin_1() {
	// QScrollerList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QScrollerList) Rend_1() {
	// QScrollerList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QScrollerList) Crbegin_0() {
	// QScrollerList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QScrollerList) Crend_0() {
	// QScrollerList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QScrollerList) Insert_1() {
	// QScrollerList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QScrollerList) Erase_0() {
	// QScrollerList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QScrollerList) Erase_1() {
	// QScrollerList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QScrollerList) Count_1() int {
	// QScrollerList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QScrollerList) Length_0() int {
	// QScrollerList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QScrollerList) First_0() *QScroller {
	// QScrollerList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// const T & constFirst()
func (this *QScrollerList) ConstFirst_0() *QScroller {
	// QScrollerList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// const T & first()
func (this *QScrollerList) First_1() *QScroller {
	// QScrollerList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// T & last()
func (this *QScrollerList) Last_0() *QScroller {
	// QScrollerList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// const T & last()
func (this *QScrollerList) Last_1() *QScroller {
	// QScrollerList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// const T & constLast()
func (this *QScrollerList) ConstLast_0() *QScroller {
	// QScrollerList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// void removeFirst()
func (this *QScrollerList) RemoveFirst_0() {
	// QScrollerList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QScrollerList) RemoveLast_0() {
	// QScrollerList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QScrollerList) StartsWith_0() bool {
	// QScrollerList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QScrollerList) EndsWith_0() bool {
	// QScrollerList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QScrollerList) Mid_0() *QScrollerList {
	// QScrollerList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QScrollerList) Value_0() *QScroller {
	// QScrollerList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// T value(int, const T &)
func (this *QScrollerList) Value_1() *QScroller {
	// QScrollerList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// void push_back(const T &)
func (this *QScrollerList) Push_back_0() {
	// QScrollerList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QScrollerList) Push_front_0() {
	// QScrollerList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QScrollerList) Front_0() *QScroller {
	// QScrollerList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// const T & front()
func (this *QScrollerList) Front_1() *QScroller {
	// QScrollerList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// T & back()
func (this *QScrollerList) Back_0() *QScroller {
	// QScrollerList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// const T & back()
func (this *QScrollerList) Back_1() *QScroller {
	// QScrollerList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScroller{}
}

// void pop_front()
func (this *QScrollerList) Pop_front_0() {
	// QScrollerList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QScrollerList) Pop_back_0() {
	// QScrollerList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QScrollerList) Empty_0() bool {
	// QScrollerList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QScrollerList) Operator_add_equal_0() *QScrollerList {
	// QScrollerList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QScrollerList) Operator_add_0() *QScrollerList {
	// QScrollerList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QScrollerList) Operator_add_equal_1() *QScrollerList {
	// QScrollerList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QScrollerList) Operator_left_shift_0() *QScrollerList {
	// QScrollerList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QScrollerList) Operator_left_shift_1() *QScrollerList {
	// QScrollerList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QScrollerList) ToVector_0() {
	// QScrollerList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QScrollerList) ToSet_0() {
	// QScrollerList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QScrollerList) FromVector_0() *QScrollerList {
	// QScrollerList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QScrollerList) FromSet_0() *QScrollerList {
	// QScrollerList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QScrollerList) FromStdList_0() *QScrollerList {
	// QScrollerList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QScrollerList) ToStdList_0() {
	// QScrollerList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QScrollerList) Detach_helper_grow_0() {
	// QScrollerList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QScrollerList) Detach_helper_0() {
	// QScrollerList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QScrollerList) Detach_helper_1() {
	// QScrollerList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QScrollerList) Dealloc_0() {
	// QScrollerList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QScrollerList) Node_construct_0() {
	// QScrollerList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QScrollerList) Node_destruct_0() {
	// QScrollerList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QScrollerList) Node_copy_0() {
	// QScrollerList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QScrollerList) Node_destruct_1() {
	// QScrollerList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QScrollerList) IsValidIterator_0() bool {
	// QScrollerList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QScrollerList) Op_eq_impl_0() bool {
	// QScrollerList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QScrollerList) Op_eq_impl_1() bool {
	// QScrollerList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QScrollerList) Contains_impl_0() bool {
	// QScrollerList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QScrollerList) Contains_impl_1() bool {
	// QScrollerList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QScrollerList) Count_impl_0() int {
	// QScrollerList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QScrollerList) Count_impl_1() int {
	// QScrollerList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScrollerList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
