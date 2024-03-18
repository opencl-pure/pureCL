//go:build !windows && !wasm

package pureCL

import (
	"errors"
	"github.com/ebitengine/purego"
	"runtime"
	"unsafe"
)

func getOpenCLPath() ([]string, error) {
	if runtime.GOOS == "linux" {
		return []string{
			"/usr/lib/libOpenCL.so",
			"/usr/local/lib/libOpenCL.so",
			"/usr/local/lib/libpocl.so",
			"/usr/lib64/libOpenCL.so",
			"/usr/lib32/libOpenCL.so",
			"libOpenCL.so"}, nil
	} else if runtime.GOOS == "darwin" {
		return []string{
			"libOpenCL.so",
			"/System/Library/Frameworks/OpenCL.framework/OpenCL"}, nil
	} else if runtime.GOOS == "android" {
		return []string{
			"/system/lib64/libOpenCL.so",
			"/system/vendor/lib64/libOpenCL.so",
			"/system/vendor/lib64/egl/libGLES_mali.so",
			"/system/vendor/lib64/libPVROCL.so",
			"/data/data/org.pocl.libs/files/lib64/libpocl.so",
			"/system/lib/libOpenCL.so",
			"/system/vendor/lib/libOpenCL.so",
			"/system/vendor/lib/egl/libGLES_mali.so",
			"/system/lib64/egl/libGLES_mali.so",
			"/system/vendor/lib/libPVROCL.so",
			"/data/data/org.pocl.libs/files/lib/libpocl.so",
			"libOpenCL.so"}, nil
	}
	return nil, errors.New("unknown system paths")
}

func loadLibrary(paths []string) error {
	if handle != 0 {
		return nil
	}
	paths2, err := getOpenCLPath()
	if err != nil {
		return err
	}
	paths = append(paths, paths2...)
	var resultError error
	for i := 0; i < len(paths); i++ {
		libOpenCl, err0 := purego.Dlopen(paths[i], purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err0 == nil {
			handle = libOpenCl
			return nil
		}
		resultError = ErrJoin(resultError, err0)
	}
	return errors.New("no path has passed:\n" + resultError.Error())
}

func initUnsupported(handle uintptr, errIn error) error {
	readImg, err := purego.Dlsym(handle, "clEnqueueReadImage")
	if err != nil {
		errIn = err
	}
	mapImg, err := purego.Dlsym(handle, "clEnqueueMapImage")
	if err != nil {
		errIn = errors.Join(errIn, err)
	}
	mapBuffer, err := purego.Dlsym(handle, "clEnqueueMapBuffer")
	if err != nil {
		errIn = errors.Join(errIn, err)
	}
	writeImg, err := purego.Dlsym(handle, "clEnqueueWriteImage")
	if err != nil {
		errIn = errors.Join(errIn, err)
	}
	// Note: Functions with unsupported arguments requiring syscall loading
	EnqueueReadImage = func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status {
		block := uintptr(0)
		if blockingRead {
			block = 1
		}
		eventList := uintptr(0)
		if eventWaitList != nil {
			eventList = uintptr(unsafe.Pointer(&eventWaitList[0]))
		}
		r1, _, _ := purego.SyscallN(readImg,
			uintptr(queue),
			uintptr(image),
			block,
			uintptr(unsafe.Pointer(&origin[0])),
			uintptr(unsafe.Pointer(&region[0])),
			uintptr(row_pitch),
			uintptr(slice_pitch),
			uintptr(ptr),
			uintptr(numEventsWaitList),
			eventList,
			uintptr(unsafe.Pointer(event)),
			0)
		return Status(r1)
	}
	EnqueueMapImage = func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin, region [3]Size, imageRowPitch, imageSlicePitch *Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr {
		block := uintptr(0)
		if blockingMap {
			block = 1
		}
		eventList := uintptr(0)
		if eventWaitList != nil {
			eventList = uintptr(unsafe.Pointer(&eventWaitList[0]))
		}
		r1, _, _ := purego.SyscallN(mapImg,
			uintptr(queue),
			uintptr(image),
			block,
			uintptr(mapFlags),
			uintptr(unsafe.Pointer(&origin[:][0])),
			uintptr(unsafe.Pointer(&region[:][0])),
			uintptr(unsafe.Pointer(imageRowPitch)),
			uintptr(unsafe.Pointer(imageSlicePitch)),
			uintptr(numEventsWaitList),
			eventList,
			uintptr(unsafe.Pointer(event)),
			uintptr(unsafe.Pointer(errCodeRet)),
		)
		return r1
	}

	EnqueueMapBuffer = func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset, size Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr {
		block := uintptr(0)
		if blockingMap {
			block = 1
		}
		eventList := uintptr(0)
		if eventWaitList != nil {
			eventList = uintptr(unsafe.Pointer(&eventWaitList[0]))
		}
		r1, _, _ := purego.SyscallN(mapBuffer,
			uintptr(queue),
			uintptr(buffer),
			uintptr(block),
			uintptr(mapFlags),
			uintptr(offset),
			uintptr(size),
			uintptr(numEventsWaitList),
			eventList,
			uintptr(unsafe.Pointer(event)),
			uintptr(unsafe.Pointer(errCodeRet)),
		)

		return r1
	}
	EnqueueWriteImage = func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status {
		block := uintptr(0)
		if blockingRead {
			block = 1
		}
		eventList := uintptr(0)
		if eventWaitList != nil {
			eventList = uintptr(unsafe.Pointer(&eventWaitList[0]))
		}
		r1, _, _ := purego.SyscallN(writeImg,
			uintptr(queue),
			uintptr(image),
			uintptr(block),
			uintptr(unsafe.Pointer(&origin[0])),
			uintptr(unsafe.Pointer(&region[0])),
			uintptr(row_pitch),
			uintptr(slice_pitch),
			uintptr(ptr),
			uintptr(numEventsWaitList),
			eventList,
			uintptr(unsafe.Pointer(event)),
		)
		return Status(r1)
	}
	return errIn
}
