module libgo

#flag -I@VMODROOT/go
#flag @VMODROOT/go/libgo.a

$if darwin {
	$if arch == "arm64" {
		#flag -arch arm64
	} else {
		#flag -arch amd64
	}
}

#include "libgo.h"

pub fn C.download_file(c_url &char,	c_path &char) &char