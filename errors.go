package pureCL

import (
	"errors"
	"fmt"
	"github.com/opencl-pure/constantCL"
)

func StatusToErr(code Status) error {
	if code == 0 {
		return nil
	}
	if err, ok := errorMap2[code]; ok {
		return errors.New(err)
	}
	return errors.New(fmt.Sprintf("cl: error %d", int32(code)))
}

func Uninitialized(name string) error {
	return errors.New(name + " function is uninitialized, it ca not call shared, it must be initialized")
}

// Common OpenCl errors
const (
	ErrDeviceNotFound                     = "cl: Device Not Found"
	ErrDeviceNotAvailable                 = "cl: Device Not Available"
	ErrCompilerNotAvailable               = "cl: Compiler Not Available"
	ErrMemObjectAllocationFailure         = "cl: Mem Object Allocation Failure"
	ErrOutOfResources                     = "cl: Out Of Resources"
	ErrOutOfHostMemory                    = "cl: Out Of Host Memory"
	ErrProfilingInfoNotAvailable          = "cl: Profiling Info Not Available"
	ErrMemCopyOverlap                     = "cl: Mem Copy Overlap"
	ErrImageFormatMismatch                = "cl: Image Format Mismatch"
	ErrImageFormatNotSupported            = "cl: Image Format Not Supported"
	ErrBuildProgramFailure                = "cl: Build Program Failure"
	ErrMapFailure                         = "cl: Map Failure"
	ErrMisalignedSubBufferOffset          = "cl: Misaligned Sub clBuffer Offset"
	ErrExecStatusErrorForEventsInWaitList = "cl: Exec Status Error For Events In Wait List"
	ErrCompileProgramFailure              = "cl: Compile Program Failure"
	ErrLinkerNotAvailable                 = "cl: Linker Not Available"
	ErrLinkProgramFailure                 = "cl: Link Program Failure"
	ErrDevicePartitionFailed              = "cl: Device Partition Failed"
	ErrKernelArgInfoNotAvailable          = "cl: Kernel Arg Info Not Available"
	ErrInvalidValue                       = "cl: Invalid Value"
	ErrInvalidDeviceType                  = "cl: Invalid Device Type"
	ErrInvalidPlatform                    = "cl: Invalid Platform"
	ErrInvalidDevice                      = "cl: Invalid Device"
	ErrInvalidContext                     = "cl: Invalid clContext"
	ErrInvalidQueueProperties             = "cl: Invalid Queue Properties"
	ErrInvalidCommandQueue                = "cl: Invalid Command Queue"
	ErrInvalidHostPtr                     = "cl: Invalid Host Ptr"
	ErrInvalidMemObject                   = "cl: Invalid Mem Object"
	ErrInvalidImageFormatDescriptor       = "cl: Invalid Image Format Descriptor"
	ErrInvalidImageSize                   = "cl: Invalid Image Size"
	ErrInvalidSampler                     = "cl: Invalid Sampler"
	ErrInvalidBinary                      = "cl: Invalid Binary"
	ErrInvalidBuildOptions                = "cl: Invalid Build Options"
	ErrInvalidProgram                     = "cl: Invalid Program"
	ErrInvalidProgramExecutable           = "cl: Invalid Program Executable"
	ErrInvalidKernelName                  = "cl: Invalid Kernel Name"
	ErrInvalidKernelDefinition            = "cl: Invalid Kernel Definition"
	ErrInvalidKernel                      = "cl: Invalid Kernel"
	ErrInvalidArgIndex                    = "cl: Invalid Arg Index"
	ErrInvalidArgValue                    = "cl: Invalid Arg Value"
	ErrInvalidArgSize                     = "cl: Invalid Arg Size"
	ErrInvalidKernelArgs                  = "cl: Invalid Kernel Args"
	ErrInvalidWorkDimension               = "cl: Invalid Work Dimension"
	ErrInvalidWorkGroupSize               = "cl: Invalid Work Group Size"
	ErrInvalidWorkItemSize                = "cl: Invalid Work Item Size"
	ErrInvalidGlobalOffset                = "cl: Invalid Global Offset"
	ErrInvalidEventWaitList               = "cl: Invalid Event Wait List"
	ErrInvalidEvent                       = "cl: Invalid Event"
	ErrInvalidOperation                   = "cl: Invalid Operation"
	ErrInvalidGlObject                    = "cl: Invalid Gl Object"
	ErrInvalidBufferSize                  = "cl: Invalid clBuffer Size"
	ErrInvalidMipLevel                    = "cl: Invalid Mip Level"
	ErrInvalidGlobalWorkSize              = "cl: Invalid Global Work Size"
	ErrInvalidProperty                    = "cl: Invalid Property"
	ErrInvalidImageDescriptor             = "cl: Invalid Image Descriptor"
	ErrInvalidCompilerOptions             = "cl: Invalid Compiler Options"
	ErrInvalidLinkerOptions               = "cl: Invalid Linker Options"
	ErrInvalidDevicePartitionCount        = "cl: Invalid Device Partition Count"
)

var errorMap2 = map[Status]string{
	constantsCL.CL_SUCCESS:                                   "",
	constantsCL.CL_DEVICE_NOT_FOUND:                          ErrDeviceNotFound,
	constantsCL.CL_DEVICE_NOT_AVAILABLE:                      ErrDeviceNotAvailable,
	constantsCL.CL_COMPILER_NOT_AVAILABLE:                    ErrCompilerNotAvailable,
	constantsCL.CL_MEM_OBJECT_ALLOCATION_FAILURE:             ErrMemObjectAllocationFailure,
	constantsCL.CL_OUT_OF_RESOURCES:                          ErrOutOfResources,
	constantsCL.CL_OUT_OF_HOST_MEMORY:                        ErrOutOfHostMemory,
	constantsCL.CL_PROFILING_INFO_NOT_AVAILABLE:              ErrProfilingInfoNotAvailable,
	constantsCL.CL_MEM_COPY_OVERLAP:                          ErrMemCopyOverlap,
	constantsCL.CL_IMAGE_FORMAT_MISMATCH:                     ErrImageFormatMismatch,
	constantsCL.CL_IMAGE_FORMAT_NOT_SUPPORTED:                ErrImageFormatNotSupported,
	constantsCL.CL_BUILD_PROGRAM_FAILURE:                     ErrBuildProgramFailure,
	constantsCL.CL_MAP_FAILURE:                               ErrMapFailure,
	constantsCL.CL_MISALIGNED_SUB_BUFFER_OFFSET:              ErrMisalignedSubBufferOffset,
	constantsCL.CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST: ErrExecStatusErrorForEventsInWaitList,
	constantsCL.CL_INVALID_VALUE:                             ErrInvalidValue,
	constantsCL.CL_INVALID_DEVICE_TYPE:                       ErrInvalidDeviceType,
	constantsCL.CL_INVALID_PLATFORM:                          ErrInvalidPlatform,
	constantsCL.CL_INVALID_DEVICE:                            ErrInvalidDevice,
	constantsCL.CL_INVALID_CONTEXT:                           ErrInvalidContext,
	constantsCL.CL_INVALID_QUEUE_PROPERTIES:                  ErrInvalidQueueProperties,
	constantsCL.CL_INVALID_COMMAND_QUEUE:                     ErrInvalidCommandQueue,
	constantsCL.CL_INVALID_HOST_PTR:                          ErrInvalidHostPtr,
	constantsCL.CL_INVALID_MEM_OBJECT:                        ErrInvalidMemObject,
	constantsCL.CL_INVALID_IMAGE_FORMAT_DESCRIPTOR:           ErrInvalidImageFormatDescriptor,
	constantsCL.CL_INVALID_IMAGE_SIZE:                        ErrInvalidImageSize,
	constantsCL.CL_INVALID_SAMPLER:                           ErrInvalidSampler,
	constantsCL.CL_INVALID_BINARY:                            ErrInvalidBinary,
	constantsCL.CL_INVALID_BUILD_OPTIONS:                     ErrInvalidBuildOptions,
	constantsCL.CL_INVALID_PROGRAM:                           ErrInvalidProgram,
	constantsCL.CL_INVALID_PROGRAM_EXECUTABLE:                ErrInvalidProgramExecutable,
	constantsCL.CL_INVALID_KERNEL_NAME:                       ErrInvalidKernelName,
	constantsCL.CL_INVALID_KERNEL_DEFINITION:                 ErrInvalidKernelDefinition,
	constantsCL.CL_INVALID_KERNEL:                            ErrInvalidKernel,
	constantsCL.CL_INVALID_ARG_INDEX:                         ErrInvalidArgIndex,
	constantsCL.CL_INVALID_ARG_VALUE:                         ErrInvalidArgValue,
	constantsCL.CL_INVALID_ARG_SIZE:                          ErrInvalidArgSize,
	constantsCL.CL_INVALID_KERNEL_ARGS:                       ErrInvalidKernelArgs,
	constantsCL.CL_INVALID_WORK_DIMENSION:                    ErrInvalidWorkDimension,
	constantsCL.CL_INVALID_WORK_GROUP_SIZE:                   ErrInvalidWorkGroupSize,
	constantsCL.CL_INVALID_WORK_ITEM_SIZE:                    ErrInvalidWorkItemSize,
	constantsCL.CL_INVALID_GLOBAL_OFFSET:                     ErrInvalidGlobalOffset,
	constantsCL.CL_INVALID_EVENT_WAIT_LIST:                   ErrInvalidEventWaitList,
	constantsCL.CL_INVALID_EVENT:                             ErrInvalidEvent,
	constantsCL.CL_INVALID_OPERATION:                         ErrInvalidOperation,
	constantsCL.CL_INVALID_GL_OBJECT:                         ErrInvalidGlObject,
	constantsCL.CL_INVALID_BUFFER_SIZE:                       ErrInvalidBufferSize,
	constantsCL.CL_INVALID_MIP_LEVEL:                         ErrInvalidMipLevel,
	constantsCL.CL_INVALID_GLOBAL_WORK_SIZE:                  ErrInvalidGlobalWorkSize,
	constantsCL.CL_INVALID_PROPERTY:                          ErrInvalidProperty,
}
