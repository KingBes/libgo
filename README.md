# libgo

Export the clang library from GoLang for use in VLang.
Use the Go library to address the existing bugs in Vlang until the bugs are fixed.

## install

```bash
v install KingBes.libgo
```

```bash
# The Golang environment needs to be installed.
~/.vmodules/kingbes/libgo/go/linux.sh  # linux
~/.vmodules/kingbes/libgo/go/macos.sh  # macos
~/.vmodules/kingbes/libgo/go/windows.cmd # windows
```

```v

import kingbe.libgo

libgo.download_file(url string, path string) string 

```