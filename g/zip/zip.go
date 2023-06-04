package gzip

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"

	gpath "github.com/khaosles/gtools2/g/path"
)

/*
   @File: zip.go
   @Author: khaosles
   @Time: 2023/5/30 00:12
   @Desc:
*/

// Compress 压缩文件

func Compress(zipFile string, files ...string) error {
	gpath.MkParentDir(zipFile)
	archive, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	for _, file := range files {
		log.Println("opening file -> " + file)
		fp, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fp.Close()
		log.Println("writing file to archive...")
		w, err := zipWriter.Create(gpath.Basename(file))
		if err != nil {
			return err
		}
		if _, err := io.Copy(w, fp); err != nil {
			return err
		}
	}
	log.Println("compress finish...")
	return nil
}

// Decompress 解压文件
func Decompress(zipFile string, dest string) {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), os.ModePerm)
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			_, err = io.Copy(f, rc)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
