module libgo

pub fn download_file(url string, path string) string {
	res := C.download_file(&char(url.str), &char(path.str))
	return unsafe { cstring_to_vstring(res) }
}