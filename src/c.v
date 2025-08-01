module libgo

#flag -I@VMODROOT/include
#flag windows @VMODROOT/lib/download.a

#include "download.h"

pub fn C.download_file(c_url &char,	c_path &char) &char