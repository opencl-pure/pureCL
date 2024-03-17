package pureCL

type Extension = string
type Version string

// TODO: Not exhaustive
const (
	// OpenCL
	Extension_khr_gl_sharing Extension = "cl_khr_gl_sharing"
	Extension_khr_fp64       Extension = "cl_khr_fp64"
	// Nvidia
	Extension_nv_pragma_unroll    Extension = "cl_nv_pragma_unroll"
	Extension_nv_compiler_options Extension = "cl_nv_compiler_options"

	Version1_0 Version = "CL1.0"
	Version1_1 Version = "CL1.1"
	Version1_2 Version = "CL1.2"
	Version2_0 Version = "CL2.0"
	Version3_0 Version = "CL3.0"
)
