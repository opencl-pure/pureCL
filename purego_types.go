package pureCL

import (
	"unsafe"
)

type ImageFormat struct {
	ChannelOrder ImageChannelOrder
	ChannelType  ImageChannelType
}

type BufferData struct {
	TypeSize uintptr
	DataSize uintptr
	Pointer  unsafe.Pointer
}

type ImageData struct {
	*BufferData
	Origin     [3]uint
	Region     [3]uint
	RowPitch   uint
	SlicePitch uint
}

type (
	// primitive types
	Size                   uint
	Status                 int32
	Program                uint
	ProgramBuildInfo       uint32
	DeviceType             uint32
	PlatformInfo           uint
	Context                uint
	Platform               uint
	Device                 uint
	DeviceInfo             uint32
	MemFlag                uint32
	ImageChannelOrder      uint32
	ImageChannelType       uint32
	CommandQueueProperties uint32
	CommandQueue           uint
	Buffer                 uint
	CommandQueueProperty   uint32
	MemInfo                uint32
	MapFlag                uint32
	Kernel                 uint
	Event                  uint

	// notify function types
	CreateContextNotifyFuncType func(errinfo, privateInfo []byte, cb Size, userData []byte)
	BuildProgramNotifyFuncType  func(program Program, userData []byte)

	// structs
	KernelArg struct {
		ptr  unsafe.Pointer
		size Size
	}

	// GL
	GLEnum          uint32
	GLInt           int32
	GLUint          uint32
	CLGLObjectType  uint32
	CLGLTextureInfo uint32

	// untested and unintergrateed
	PipeProperties         uint32
	KernelWorkGroupInfo    uint32
	EventInfo              uint32
	EventCommandExecStatus uint32
	ImageInfo              uint32
	DeviceAffinityDomain   uint64
	ContextProperty        uint
	ContextProperties      struct {
		Platform *Platform
		// Interop
		InteropUserSync *bool
		// OpenGL
		GLContextKHR *uint
		// Windows
		WGL_HDC_KHR *uint
	}

	DevicePartitionProperty struct {
		Type  uint32
		Flags uint32
	}
	DevicePartition struct {
		Properties     []DevicePartitionProperty
		AffinityDomain DeviceAffinityDomain
	}
	Sampler       uint
	MemObjectType uint32
	MapPointer    unsafe.Pointer
	BufferRect    struct {
		Origin [3]Size
		Region [3]Size
	}
	ImageDesc struct {
		Type        MemObjectType
		Width       Size
		Height      Size
		Depth       Size
		ArraySize   Size
		RowPitch    Size
		SlisePitch  Size
		Buffer      Buffer
		NumMipLevel uint32
		NumSamples  uint32
	}
)

func NewKernelArg[T any](arg *T) KernelArg {
	return KernelArg{
		ptr:  unsafe.Pointer(arg),
		size: Size(unsafe.Sizeof(*arg)),
	}
}

func (k Kernel) SetArg(index uint, arg KernelArg) error {
	if SetKernelArg == nil {
		return Uninitialized("SetKernelArg")
	}
	return StatusToErr(SetKernelArg(k, uint32(index), arg.size, arg.ptr)) //TODO uint or uint32
}

func (k Kernel) Release() error {
	if ReleaseKernel == nil {
		return Uninitialized("ReleaseKernel")
	}
	return StatusToErr(ReleaseKernel(k))
}
