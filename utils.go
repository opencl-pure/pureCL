package pureCL

import (
	"errors"
	"fmt"
	"github.com/ebitengine/purego"
	"image"
	"unsafe"
)

type BufferType interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func GetBufferData[T BufferType](data []T) *BufferData {
	size := unsafe.Sizeof(data[0])
	return &BufferData{
		TypeSize: size,
		DataSize: uintptr(len(data)) * size,
		Pointer:  unsafe.Pointer(&data[0]),
	}
}

func GetImageBufferData(img image.RGBA) *ImageData {
	bounds := img.Bounds()
	return &ImageData{
		BufferData: GetBufferData(img.Pix),
		Origin: [3]uint{
			uint(bounds.Min.X), uint(bounds.Min.Y), 0,
		},
		Region: [3]uint{
			uint(bounds.Dx()), uint(bounds.Dy()), 1,
		},
		RowPitch:   0,
		SlicePitch: 0,
	}
}

func registerLibFuncWithoutPanic(fptr interface{}, handle uintptr, name string, error0 error) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = ErrJoin(error0, errors.New(fmt.Sprint(r)))
		} else {
			e = error0
		}
	}()
	purego.RegisterLibFunc(fptr, handle, name)
	return
}

func ErrJoin(e1, e2 error) error {
	if e1 != nil && e2 != nil {
		return errors.New(e1.Error() + ";\n" + e2.Error())
	}
	if e1 != nil {
		return e1
	}
	return e2
}
