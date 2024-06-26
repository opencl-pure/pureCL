package pureCL

import (
	"github.com/ebitengine/purego"
)

func SetHandle(handleExternal uintptr) {
	handle = handleExternal
}

func Init(version Version, paths ...string) error {
	err := loadLibrary(paths)
	if err != nil {
		return err
	}
	// Platform
	err = registerLibFuncWithoutPanic(&GetPlatformIDs, handle, "clGetPlatformIDs", nil)
	err = registerLibFuncWithoutPanic(&GetPlatformInfo, handle, "clGetPlatformInfo", err)
	// Device
	err = registerLibFuncWithoutPanic(&GetDeviceIDs, handle, "clGetDeviceIDs", err)
	err = registerLibFuncWithoutPanic(&GetDeviceInfo, handle, "clGetDeviceInfo", err)
	err = registerLibFuncWithoutPanic(&ReleaseDevice, handle, "clReleaseDevice", err)
	// Event
	err = registerLibFuncWithoutPanic(&ReleaseEvent, handle, "clReleaseEvent", err)
	err = registerLibFuncWithoutPanic(&WaitForEvents, handle, "clWaitForEvents", err)
	// Context
	err = registerLibFuncWithoutPanic(&CreateContext, handle, "clCreateContext", err)
	err = registerLibFuncWithoutPanic(&ReleaseContext, handle, "clReleaseContext", err)
	err = registerLibFuncWithoutPanic(&CreateProgramWithSource, handle, "clCreateProgramWithSource", err)
	err = registerLibFuncWithoutPanic(&CreateBuffer, handle, "clCreateBuffer", err)
	err = registerLibFuncWithoutPanic(&CreateImage2D, handle, "clCreateImage2D", err)
	err = registerLibFuncWithoutPanic(&CreateImage3D, handle, "clCreateImage3D", err)
	err = registerLibFuncWithoutPanic(&CreateImage, handle, "clCreateImage", err)
	// Command queue
	if version == Version2_0 || version == Version3_0 {
		err = registerLibFuncWithoutPanic(&CreateCommandQueueWithProperties, handle, "clCreateCommandQueueWithProperties", err)
	}
	err = registerLibFuncWithoutPanic(&CreateCommandQueue, handle, "clCreateCommandQueue", err)

	err = registerLibFuncWithoutPanic(&EnqueueBarrier, handle, "clEnqueueBarrier", err)
	err = registerLibFuncWithoutPanic(&EnqueueNDRangeKernel, handle, "clEnqueueNDRangeKernel", err)
	err = registerLibFuncWithoutPanic(&EnqueueReadBuffer, handle, "clEnqueueReadBuffer", err)
	err = registerLibFuncWithoutPanic(&EnqueueWriteBuffer, handle, "clEnqueueWriteBuffer", err)

	err = initUnsupported(handle, err)

	err = registerLibFuncWithoutPanic(&EnqueueUnmapMemObject, handle, "clEnqueueUnmapMemObject", err)
	err = registerLibFuncWithoutPanic(&FinishCommandQueue, handle, "clFinish", err)
	err = registerLibFuncWithoutPanic(&FlushCommandQueue, handle, "clFlush", err)
	err = registerLibFuncWithoutPanic(&ReleaseCommandQueue, handle, "clReleaseCommandQueue", err)
	// Program
	err = registerLibFuncWithoutPanic(&BuildProgram, handle, "clBuildProgram", err)
	err = registerLibFuncWithoutPanic(&GetProgramBuildInfo, handle, "clGetProgramBuildInfo", err)
	err = registerLibFuncWithoutPanic(&GetProgramInfo, handle, "clGetProgramInfo", err)
	err = registerLibFuncWithoutPanic(&CreateKernel, handle, "clCreateKernel", err)
	err = registerLibFuncWithoutPanic(&ReleaseProgram, handle, "clReleaseProgram", err)
	// Kernel
	err = registerLibFuncWithoutPanic(&SetKernelArg, handle, "clSetKernelArg", err)
	err = registerLibFuncWithoutPanic(&ReleaseKernel, handle, "clReleaseKernel", err)
	// Buffer
	err = registerLibFuncWithoutPanic(&GetMemObjectInfo, handle, "clGetMemObjectInfo", err)
	err = registerLibFuncWithoutPanic(&ReleaseMemObject, handle, "clReleaseMemObject", err)
	if err != nil {
		err = ErrJoin(err, purego.Dlclose(handle))
	}
	return err
}

func InitializeGLSharing(paths ...string) error {
	err := loadLibrary(paths)
	if err != nil {
		return err
	}
	err = registerLibFuncWithoutPanic(&CreateFromGLTexture, handle, "clCreateFromGLTexture", err)
	err = registerLibFuncWithoutPanic(&EnqueueAcquireGLObjects, handle, "clEnqueueAcquireGLObjects", err)
	err = registerLibFuncWithoutPanic(&EnqueueReleaseGLObjects, handle, "clEnqueueReleaseGLObjects", err)
	err = registerLibFuncWithoutPanic(&GetGLObjectInfo, handle, "clGetGLObjectInfo", err)
	err = registerLibFuncWithoutPanic(&GetGLTextureInfo, handle, "clGetGLTextureInfo", err)
	if err != nil {
		err = ErrJoin(err, purego.Dlclose(handle))
	}
	return err
}
