# pure
This is fork and simplify of [Zyko0's opencl](https://github.com/Zyko0/go-opencl), big thank. <br> 
This package provide low level wrapper to OpenCL,
that means it is 1:1 wrapper C:GO - no GO error handling
only GO types map OpenCL function and cl_types without cgo, powered by
[purego](https://github.com/ebitengine/purego) and inspired by [libopencl-stub](https://github.com/krrishnarraj/libopencl-stub). 
This package use [constantsCL](github.com/opencl-pure/constantsCL) <br>
Thank to both of them!
# goal
- low level wrap of OpenCL for
   - [middleCL](https://github.com/opencl-pure/middleCL)
   - [highCL](https://github.com/opencl-pure/highCL)
   - and maybe others ...,
-  be minimalistic
-  
-  try to have all functions of OpenCL (so if you have some, give PR)
-  easy to multiplatform (thank [purego](https://github.com/ebitengine/purego))
-  easy find path (custumize path to openclLib shared library)
-  easy to compile, we do not need cgo and knowing link to shared library
-  try [purego](https://github.com/ebitengine/purego) and bring opencl on android without complicate link
# not goal
- be faster as cgo version, [purego](https://github.com/ebitengine/purego) is using same mechanism as cgo 
# examples
you can also use to your package independent from others ..., you can also use as this:
# example
```go
package main

import (
	"errors"
	constants "github.com/opencl-pure/constantsCL"
	pure "github.com/opencl-pure/pureCL"
	"log"
)

func main() {
	err := pure.Init(pure.Version2_0) //init with version of OpenCL
	if err != nil {
		log.Println(err)
		return
	}
	numPlatforms := uint32(0)
	st := pure.GetPlatformIDs(0, nil, &numPlatforms)
	if st != constants.CL_SUCCESS {
		log.Println(errors.New("oops platform error"))
		return
	}

	platformIDs := make([]pure.Platform, numPlatforms)
	st = pure.GetPlatformIDs(numPlatforms, platformIDs, nil)
	if st != constants.CL_SUCCESS {
		log.Println(errors.New("oops none ...."))
		return
	}
	// ....
}    
```
