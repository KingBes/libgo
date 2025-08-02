module libgo

#flag -I@VMODROOT/go
#flag windows @VMODROOT/go/download.a

#include "download.h"

pub fn C.download_file(c_url &char,	c_path &char) &char