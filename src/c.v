module libgo

#flag -I@VMODROOT/go
#flag @VMODROOT/go/libgo.a

#include "libgo.h"

pub fn C.download_file(c_url &char,	c_path &char) &char