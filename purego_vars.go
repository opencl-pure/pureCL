package pureCL

import "unsafe"

// functions
var (
	// GetPlatformIDs this wrap opencl clGetPlatformIDs and do returns a list of available OpenCL platforms
	GetPlatformIDs func(numEntries uint32, platforms []Platform, numPlatforms *uint32) Status = nil
	// GetPlatformInfo this wrap opencl clGetPlatformInfo and do queries information about a specific platform
	GetPlatformInfo func(platform Platform, platformInfo PlatformInfo, paramValueSize Size, paramValue []byte, paramValueSizeRet *Size) Status = nil

	// GetDeviceIDs this wrap opencl clGetDeviceIDs and do returns a list of available OpenCL devices
	GetDeviceIDs func(platform Platform, deviceType DeviceType, numEntries uint32, devices []Device, numDevices *uint32) Status = nil
	// GetDeviceInfo this wrap opencl clGetDeviceInfo and do queries information about a specific device
	GetDeviceInfo func(device Device, deviceInfo DeviceInfo, paramValueSize Size, paramValue []byte, paramValueSizeRet *Size) Status = nil
	// ReleaseDevice this wrap opencl clReleaseDevice and do releases the OpenCL device
	ReleaseDevice func(id Device) Status = nil

	// ReleaseEvent this wrap opencl clReleaseEvent and do releases an OpenCL event
	ReleaseEvent func(event Event) Status = nil
	// WaitForEvents this wrap opencl clWaitForEvents and do waits on the host thread for commands identified by num_events to complete
	WaitForEvents func(numEvents uint32, eventList []Event) Status = nil

	// CreateContext this wrap opencl clCreateContext and do creates an OpenCL context
	CreateContext func(properties unsafe.Pointer, numDevices uint32, devices []Device, pfnNotify *CreateContextNotifyFuncType, userData []byte, errCodeRet *Status) Context = nil
	// ReleaseContext this wrap opencl clReleaseContext and do releases the OpenCL context
	ReleaseContext func(ctx Context) Status = nil
	// CreateProgramWithSource this wrap opencl clCreateProgramWithSource and do creates a program object for a context
	CreateProgramWithSource func(ctx Context, count Size, strings []string, lengths []Size, errCodeRet *Status) Program = nil
	// CreateBuffer this wrap opencl clCreateBuffer and do creates a buffer object
	CreateBuffer func(ctx Context, memFlags MemFlag, size Size, hostPtr unsafe.Pointer, errCodeRet *Status) Buffer = nil
	// CreateImage2D this wrap opencl clCreateImage2D and do creates a 2D image object
	CreateImage2D func(ctx Context, memFlags MemFlag, imageFormat *ImageFormat, imageWidth, imageHeight, imageRowPitch Size, hostPtr unsafe.Pointer, errCodeRet *Status) Buffer = nil

	// CreateCommandQueue this wrap opencl clCreateCommandQueue and do creates a command-queue on a specific device
	CreateCommandQueue func(context Context, device Device, properties CommandQueueProperty, errCodeRet *Status) CommandQueue = nil
	// CreateCommandQueueWithProperties this wrap opencl clCreateCommandQueueWithProperties and do creates a command-queue on a specific device with specified properties
	CreateCommandQueueWithProperties func(context Context, device Device, properties CommandQueueProperty, errCodeRet *Status) CommandQueue = nil
	// EnqueueBarrier this wrap opencl clEnqueueBarrier and do inserts a barrier command
	EnqueueBarrier func(queue CommandQueue) Status = nil
	// EnqueueNDRangeKernel this wrap opencl clEnqueueNDRangeKernel and do enqueue a kernel to execute on a device
	EnqueueNDRangeKernel func(queue CommandQueue, kernel Kernel, workDim uint, globalWorkOffset, globalWorkSize, localWorkSize []Size, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil
	// EnqueueReadBuffer this wrap opencl clEnqueueReadBuffer and do enqueue a command to read from a buffer object to host memory
	EnqueueReadBuffer func(queue CommandQueue, buffer Buffer, blockingRead bool, offset, cb Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil
	// EnqueueReadImage this wrap opencl clEnqueueReadImage and do enqueue a command to read from a 2D or 3D image object to host memory
	//TODO eventWaitList is ignored due to syscall init
	EnqueueReadImage func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil
	// EnqueueWriteImage this wrap opencl clEnqueueWriteImage and do enqueue a command to write from host memory to a 2D or 3D image object
	//TODO eventWaitList is ignored due to syscall init
	EnqueueWriteImage func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil
	// EnqueueWriteBuffer this wrap opencl clEnqueueWriteBuffer and do enqueue a command to write to a buffer object from host memory
	EnqueueWriteBuffer func(queue CommandQueue, buffer Buffer, blockingWrite bool, offset, cb Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil
	// EnqueueMapBuffer this wrap opencl clEnqueueMapBuffer and do enqueue a command to map a buffer object into the host address space
	//TODO eventWaitList is ignored due to syscall init
	EnqueueMapBuffer func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset, size Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr = nil
	// EnqueueUnmapMemObject this wrap opencl clEnqueueUnmapMemObject and do enqueue a command to unmap a previously mapped buffer object
	EnqueueUnmapMemObject func(queue CommandQueue, buffer Buffer, mappedPtr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil
	// EnqueueMapImage this wrap opencl clEnqueueMapImage and do enqueue a command to map a 2D or 3D image object into the host address space
	//TODO eventWaitList is ignored due to syscall init
	EnqueueMapImage func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin, region [3]Size, imageRowPitch, imageSlicePitch *Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr = nil
	// FinishCommandQueue this wrap opencl clFinish and do issues all previously queued OpenCL commands in a command-queue to the device
	FinishCommandQueue func(queue CommandQueue) Status = nil
	// FlushCommandQueue this wrap opencl clFlush and do ensures that all previously queued OpenCL commands in a command-queue are submitted to the device
	FlushCommandQueue func(queue CommandQueue) Status = nil
	// ReleaseCommandQueue this wrap opencl clReleaseCommandQueue and do releases a kernel object
	ReleaseCommandQueue func(queue CommandQueue) Status = nil

	// BuildProgram this wrap opencl clBuildProgram and do builds (compiles & links) a program executable from the program source or binary
	BuildProgram func(program Program, numDevices uint32, devices []Device, options []byte, pfnNotify *BuildProgramNotifyFuncType, userData []byte) Status = nil
	// GetProgramBuildInfo this wrap opencl clGetProgramBuildInfo and do returns build information for each device in the program object
	GetProgramBuildInfo func(program Program, device Device, info ProgramBuildInfo, paramSize Size, paramValue unsafe.Pointer, paramSizeRet *Size) Status = nil
	// CreateKernel this wrap opencl clCreateKernel and do creates a kernel object
	CreateKernel func(program Program, kernelName string, errCodeRet *Status) Kernel = nil
	// ReleaseProgram this wrap opencl clReleaseProgram and do releases the OpenCL program object
	ReleaseProgram func(program Program) Status = nil
	// GetProgramInfo this wrap opencl clGetProgramInfo and do queries information about a program object
	GetProgramInfo func(program Program, info ProgramBuildInfo, size Size, pointer unsafe.Pointer, t *Size) Status = nil

	// SetKernelArg this wrap opencl clSetKernelArg and do sets the argument value for a specific argument of a kernel
	SetKernelArg func(kernel Kernel, argIndex uint32, argSize Size, argValue unsafe.Pointer) Status = nil
	// ReleaseKernel this wrap opencl clReleaseKernel and do releases a kernel object
	ReleaseKernel func(kernel Kernel) Status = nil

	// GetMemObjectInfo this wrap opencl clGetMemObjectInfo and do queries information about a memory object
	GetMemObjectInfo func(buffer Buffer, memInfo MemInfo, paramValueSize Size, paramValue unsafe.Pointer, paramValueSizeRet *Size) Status = nil
	// ReleaseMemObject this wrap opencl clReleaseMemObject and do releases an OpenCL memory object
	ReleaseMemObject func(buffer Buffer) Status = nil

	// CreateFromGLTexture this wrap opencl clCreateFromGLTexture and do creates a 2D image object from an OpenGL texture object
	CreateFromGLTexture func(ctx Context, memFlags MemFlag, textureTarget GLEnum, mipLevel GLInt, texture GLUint, errCodeRet *Status) Buffer = nil
	// GetGLObjectInfo this wrap opencl clGetGLObjectInfo and do returns information about the OpenCL memory object and OpenGL object
	GetGLObjectInfo func(memObj Buffer, objectType *CLGLObjectType, objectName *GLUint) Status = nil
	// GetGLTextureInfo this wrap opencl clGetGLTextureInfo and do returns information about the OpenGL texture object associated with a memory object
	GetGLTextureInfo func(memObj Buffer, paramName CLGLTextureInfo, paramValueSize Size, paramValue unsafe.Pointer, paramValueSizeRet *Size) Status = nil
	// EnqueueAcquireGLObjects this wrap opencl clEnqueueAcquireGLObjects and do enqueue commands to acquire OpenCL memory objects that have been created from OpenGL objects
	EnqueueAcquireGLObjects func(queue CommandQueue, numObjects uint32, memObjects unsafe.Pointer, numEventsInWaitList uint32, eventWaitList []Event, event *Event) Status = nil
	// EnqueueReleaseGLObjects this wrap opencl clEnqueueReleaseGLObjects and do enqueue commands to release OpenCL memory objects that have been created from OpenGL objects
	EnqueueReleaseGLObjects func(queue CommandQueue, numObjects uint32, memObjects unsafe.Pointer, numEventsInWaitList uint32, eventWaitList []Event, event *Event) Status = nil
)
