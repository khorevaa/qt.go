package qtpositioning

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtPositioning/qgeocoordinate.h:125
// index:42
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QGeoCoordinate &, uint)

/*

 */
func QHash_42(coordinate QGeoCoordinate_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if coordinate != nil && coordinate.QGeoCoordinate_PTR() != nil {
		convArg0 = coordinate.QGeoCoordinate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK14QGeoCoordinatej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

//  body block end
