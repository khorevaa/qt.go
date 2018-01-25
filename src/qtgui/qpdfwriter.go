package qtgui

// /usr/include/qt/QtGui/qpdfwriter.h
// #include <qpdfwriter.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QPdfWriter struct {
	*qtcore.QObject
	*QPagedPaintDevice
}

func (this *QPdfWriter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QPdfWriter) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QPagedPaintDevice = NewQPagedPaintDeviceFromPointer(cthis)
}
func NewQPdfWriterFromPointer(cthis unsafe.Pointer) *QPdfWriter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQPagedPaintDeviceFromPointer(cthis)
	return &QPdfWriter{bcthis0, bcthis1}
}
func (*QPdfWriter) NewFromPointer(cthis unsafe.Pointer) *QPdfWriter {
	return NewQPdfWriterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpdfwriter.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPdfWriter) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpdfwriter.h:60
// index:0
// Public
// void QPdfWriter(const class QString &)
func NewQPdfWriter(filename *qtcore.QString) *QPdfWriter {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = filename.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:61
// index:1
// Public
// void QPdfWriter(class QIODevice *)
func NewQPdfWriter_1(device *qtcore.QIODevice /*444 QIODevice **/) *QPdfWriter {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:62
// index:0
// Public virtual
// void ~QPdfWriter()
func DeleteQPdfWriter(*QPdfWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:64
// index:0
// Public
// void setPdfVersion(enum QPagedPaintDevice::PdfVersion)
func (this *QPdfWriter) SetPdfVersion(version int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setPdfVersionEN17QPagedPaintDevice10PdfVersionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:65
// index:0
// Public
// QPagedPaintDevice::PdfVersion pdfVersion()
func (this *QPdfWriter) PdfVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10pdfVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:67
// index:0
// Public
// QString title()
func (this *QPdfWriter) Title() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter5titleEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpdfwriter.h:68
// index:0
// Public
// void setTitle(const class QString &)
func (this *QPdfWriter) SetTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter8setTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:70
// index:0
// Public
// QString creator()
func (this *QPdfWriter) Creator() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter7creatorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpdfwriter.h:71
// index:0
// Public
// void setCreator(const class QString &)
func (this *QPdfWriter) SetCreator(creator *qtcore.QString) {
	var convArg0 = creator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter10setCreatorERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:73
// index:0
// Public virtual
// bool newPage()
func (this *QPdfWriter) NewPage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter7newPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpdfwriter.h:75
// index:0
// Public
// void setResolution(int)
func (this *QPdfWriter) SetResolution(resolution int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setResolutionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:76
// index:0
// Public
// int resolution()
func (this *QPdfWriter) Resolution() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10resolutionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpdfwriter.h:89
// index:0
// Public virtual
// void setPageSize(enum QPagedPaintDevice::PageSize)
func (this *QPdfWriter) SetPageSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter11setPageSizeEN17QPagedPaintDevice8PageSizeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:90
// index:0
// Public virtual
// void setPageSizeMM(const class QSizeF &)
func (this *QPdfWriter) SetPageSizeMM(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setPageSizeMMERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:95
// index:0
// Protected virtual
// QPaintEngine * paintEngine()
func (this *QPdfWriter) PaintEngine() *QPaintEngine /*444 QPaintEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter11paintEngineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpdfwriter.h:96
// index:0
// Protected virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPdfWriter) Metric(id int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end