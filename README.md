# version
Reusable wrappers/utils to extract version(s) from golang Buildinfo included in `go install`ed binaries

```golang
import fortio.org/version

var shortVersion, longVersion, fullVersion := version.FromBuildInfo()
```

Complete example in [sample/simpleMain.go](sample/simpleMain.go)
