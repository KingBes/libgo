# libgo

Export the clang library from GoLang for use in VLang.
Use the Go library to address the existing bugs in Vlang until the bugs are fixed.

## install

```bash
v install KingBes.libgo
```

```bash
# The Golang environment needs to be installed.
~/.vmodules/kingbes/libgo/go/build.sh  # linux
~/.vmodules/kingbes/libgo/go/build.cmd # windows
```

```v

import kingbe.libgo

libgo.download_file(url string, path string) string 

```