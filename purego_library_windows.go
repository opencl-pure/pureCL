//go:build windows

package pureCL

import (
	"errors"
	"syscall"
	"unsafe"
)

func loadLibrary(paths []string) (uintptr, error) {
	paths = append(paths, "opencl.dll")
	var resultError error
	for i := 0; i < len(paths); i++ {
		handle, err := syscall.LoadLibrary(paths)
		if err == nil {
			return uintptr(handle), err
		}
		resultError = ErrJoin(resultError, err)
	}
	return uintptr(0), ErrJoin(errors.New("no path has passed: "), resultError)
}

func initUnsupported(handle syscall.Handle, errIn error) error {
	// purego unsupported functions
	dll := syscall.DLL{
		Name:   "opencl.dll",
		Handle: handle,
	}

	// Note: Functions with unsupported arguments requiring syscall loading
	readImg, err := dll.FindProc("clEnqueueReadImage")
	if err != nil {
		errIn = errors.Join(err)
	}
	mapImg, err := dll.FindProc("clEnqueueMapImage")
	if err != nil {
		errIn = errors.Join(err)
	}
	mapBuffer, err := dll.FindProc("clEnqueueMapBuffer")
	if err != nil {
		errIn = errors.Join(err)
	}
	writeImg, err := dll.FindProc("clEnqueueWriteImage")
	if err != nil {
		errIn = errors.Join(err)
	}
	EnqueueReadImage = func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status {
		block := uintptr(0)
		if blockingRead {
			block = 1
		}
		eventList := uintptr(0)
		if eventWaitList != nil {
			eventList = uintptr(unsafe.Pointer(&eventWaitList[0]))
		}
		r1, _, _ := readImg.Call(
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
	EnqueueMapImage = func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin, region [3]Size, imageRowPitch, imageSlicePitch *Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr {
		block := uintptr(0)
		if blockingMap {
			block = 1
		}
		eventList := uintptr(0)
		if eventWaitList != nil {
			eventList = uintptr(unsafe.Pointer(&eventWaitList[0]))
		}
		r1, _, _ := mapImg.Call(
			uintptr(queue),
			uintptr(image),
			uintptr(block),
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
		r1, _, _ := mapBuffer.Call(
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
		r1, _, _ := writeImg.Call(
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
