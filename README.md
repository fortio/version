# version
Reusable wrappers/utils to extract version(s) from golang Buildinfo included in `go install`ed binaries

```golang
import fortio.org/version

var shortVersion, longVersion, fullVersion := version.FromBuildInfo()
```

Complete example in [sample/simpleMain.go](sample/simpleMain.go)

Yielding:
```shell
$ go install fortio.org/version/sample@latest
go: downloading fortio.org/version v1.0.0
$ ~/go/bin/sample -h
fortio.org/version sample main 1.0.0 usage:
	/Users/dl/go/bin/sample [flags]
flags:
  -buildinfo
    	Show full build info and exit.
$ ~/go/bin/sample -buildinfo
1.0.0 h1:JbGoGiNQ0883KoVPDsYhQCQ32QkAVTtECn86XVRRYi4= go1.19.5 arm64 darwin
go	go1.19.5
path	fortio.org/version/sample
mod	fortio.org/version	v1.0.0	h1:JbGoGiNQ0883KoVPDsYhQCQ32QkAVTtECn86XVRRYi4=
build	-compiler=gc
build	CGO_ENABLED=1
build	CGO_CFLAGS=
build	CGO_CPPFLAGS=
build	CGO_CXXFLAGS=
build	CGO_LDFLAGS=
build	GOARCH=arm64
build	GOOS=darwin
$ ~/go/bin/sample
2023/02/12 13:34:37 fortio.org/version sample main started 1.0.0 h1:JbGoGiNQ0883KoVPDsYhQCQ32QkAVTtECn86XVRRYi4= go1.19.5 arm64 darwin
```

Libraries can embed their own version using for instance
```golang
func init() { //nolint:gochecknoinits // we do need an init for this
	shortVersion, longVersion, fullVersion = version.FromBuildInfoPath("fortio.org/fortio")
}
```

and the `longVersion` in this example can then show both the module/library version and the version of the containing binary. eg

fortio 1.40.0 h1:jSDO/jGcyC/qTpMZZ84EZbn9BQawsWM9/RMQ9s6Cn3w= go1.19.5 arm64 darwin (in fortio.org/fortiotel 1.3.0)
