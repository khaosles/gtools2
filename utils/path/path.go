package gpath

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/khaosles/gtools2/core/log"
)

// Exist judge whether exists filepath
func Exist(path string) bool {
	path = Format(path)
	// path stat
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// Format the path
func Format(path string) string {
	// delete the space at the both ends
	path = strings.TrimSpace(path)
	// simplified path
	path = filepath.Clean(path)
	// \\ to /
	path = filepath.ToSlash(path)
	// / to \ or /
	path = filepath.FromSlash(path)
	return path
}

// IsFile judge whether is a file
func IsFile(path string) bool {
	path = Format(path)
	// path stat
	fileStat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileStat.IsDir()
}

// IsDir judge whether is a dir
func IsDir(path string) bool {
	path = Format(path)
	// path stat
	fileStat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileStat.IsDir()
}

// FileSize obtain file size
func FileSize(path string) int64 {
	path = Format(path)
	if !IsFile(path) {
		glog.Error("FileNotExist: '" + path + "' not exist")
		return 0
	}
	fileStat, err := os.Stat(path)
	if err != nil {
		glog.Error("FileOpenError: cannot open the file" + path)
		return 0
	}
	return fileStat.Size()
}

// Basename get filename
func Basename(path string) string {
	path = Format(path)
	return filepath.Base(path)
}

// Join the path
func Join(elem ...string) string {
	return Format(filepath.Join(elem...))
}

// Dirname get file dir name
func Dirname(path string) string {
	path = Format(path)
	return filepath.Dir(path)
}

// Split get file dir name
func Split(path string) (string, string) {
	path = Format(path)
	return filepath.Split(path)
}

// Suffix get file suffix
func Suffix(path string) string {
	path = Format(path)
	return filepath.Ext(path)
}

// Mkdir create a folder
func Mkdir(path string) {
	path = Format(path)
	// if path is a file, raise an error
	if IsFile(path) {
		glog.Error("FolderCreateError: path is a file, " + path)
	}
	if !IsDir(path) {
		// create the folder
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			glog.Error("FolderCreateError: " + err.Error())
		}
	}
}

// MkParentDir create a path's parent dir
func MkParentDir(path string) {
	Mkdir(Dirname(path))
}

// Abs get file absolute path
func Abs(path string) string {
	path = Format(path)
	if filepath.IsAbs(path) {
		return path
	} else {
		absPath, err := filepath.Abs(path)
		if err != nil {
			glog.Error("AbsPathGetError" + err.Error())
		}
		return absPath
	}
}

// Remove file or folder
func Remove(path string) {
	if !Exist(path) {
		return
	}
	err := os.RemoveAll(path)
	if err != nil {
		glog.Error("RemoveFileError: " + err.Error())
	}
}

// RemoveFile remove a file
func RemoveFile(path string) {
	if !IsFile(path) {
		return
	}
	err := os.Remove(path)
	if err != nil {
		glog.Error("RemoveFileError: " + err.Error())
	}
}

func RootPath() string {
	rootPath, _ := os.Getwd()
	return rootPath
}

func Rename(src, dst string) error {
	err := os.Rename(src, dst)
	if err != nil {
		return err
	}
	return nil
}

func Rel(basepath, targpath string) (string, error) {
	relpath, err := filepath.Rel(basepath, targpath)
	if err != nil {
		return "", err
	}
	return relpath, nil
}

func FileNameWithoutExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func GenUniqueFilename(filename string, tries int, rule func(name string) string) (string, error) {
	if !Exist(filename) {
		return filename, nil
	}

	name := FileNameWithoutExt(filename)
	ext := filepath.Ext(filename)
	var newName string
	i := 1

	for {
		if rule != nil {
			newName = rule(name)
		} else {
			newName = fmt.Sprintf("%s%d", name, i)
		}
		newFilename := newName + ext
		if !Exist(newFilename) {
			return newFilename, nil
		}
		if i > tries {
			return "", errors.New("too many tries")
		}
		i++

	}
}

func CleanPath(p string) string {
	const stackBufSize = 128

	// Turn empty string into "/"
	if p == "" {
		return "/"
	}

	// Reasonably sized buffer on stack to avoid allocations in the common case.
	// If a larger buffer is required, it gets allocated dynamically.
	buf := make([]byte, 0, stackBufSize)

	n := len(p)

	// Invariants:
	//      reading from path; r is index of next byte to process.
	//      writing to buf; w is index of next byte to write.

	// path must start with '/'
	r := 1
	w := 1

	if p[0] != '/' {
		r = 0

		if n+1 > stackBufSize {
			buf = make([]byte, n+1)
		} else {
			buf = buf[:n+1]
		}
		buf[0] = '/'
	}

	trailing := n > 1 && p[n-1] == '/'

	// A bit more clunky without a 'lazybuf' like the path package, but the loop
	// gets completely inlined (bufApp calls).
	// So in contrast to the path package this loop has no expensive function
	// calls (except make, if needed).

	for r < n {
		switch {
		case p[r] == '/':
			// empty path element, trailing slash is added after the end
			r++

		case p[r] == '.' && r+1 == n:
			trailing = true
			r++

		case p[r] == '.' && p[r+1] == '/':
			// . element
			r += 2

		case p[r] == '.' && p[r+1] == '.' && (r+2 == n || p[r+2] == '/'):
			// .. element: remove to last /
			r += 3

			if w > 1 {
				// can backtrack
				w--

				if len(buf) == 0 {
					for w > 1 && p[w] != '/' {
						w--
					}
				} else {
					for w > 1 && buf[w] != '/' {
						w--
					}
				}
			}

		default:
			// Real path element.
			// Add slash if needed
			if w > 1 {
				bufApp(&buf, p, w, '/')
				w++
			}

			// Copy element
			for r < n && p[r] != '/' {
				bufApp(&buf, p, w, p[r])
				w++
				r++
			}
		}
	}

	// Re-append trailing slash
	if trailing && w > 1 {
		bufApp(&buf, p, w, '/')
		w++
	}

	// If the original string was not modified (or only shortened at the end),
	// return the respective substring of the original string.
	// Otherwise return a new string from the buffer.
	if len(buf) == 0 {
		return p[:w]
	}
	return string(buf[:w])
}

// Internal helper to lazily create a buffer if necessary.
// Calls to this function get inlined.
func bufApp(buf *[]byte, s string, w int, c byte) {
	b := *buf
	if len(b) == 0 {
		// No modification of the original string so far.
		// If the next character is the same as in the original string, we do
		// not yet have to allocate a buffer.
		if s[w] == c {
			return
		}

		// Otherwise use either the stack buffer, if it is large enough, or
		// allocate a new buffer on the heap, and copy all previous characters.
		if l := len(s); l > cap(b) {
			*buf = make([]byte, len(s))
		} else {
			*buf = (*buf)[:l]
		}
		b = *buf

		copy(b, s[:w])
	}
	b[w] = c
}
