package files

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

// CreateTar takes the name of the tar to be created and the list of files
func CreateTar(name string, files []string) error {
	// creating tar file
	tar_file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer tar_file.Close()
	// gzip writer
	gz := gzip.NewWriter(tar_file)
	defer gz.Close()

	// tar writer
	tr := tar.NewWriter(gz)
	defer tr.Close()

	// iterate over files
	for _, file := range files {
		// copy each file to tar
		err = copyFileToTar(tr,file)
		// error out in case of an error 
		if err != nil {
			return err
		}
	}

	// successfully completed
	return nil
}

func copyFileToTar(tar_writer *tar.Writer, filename string) error {
	// open
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// info about the file to pass to tar
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// creating header 
	hearder, err := tar.FileInfoHeader(info,info.Name())
	if err != nil {
		return err
	}

	// to preserve the directry structure of the original file system
	hearder.Name = filename

	err = tar_writer.WriteHeader(hearder)
	if err != nil {
		return err
	}

	// copy the file
	_, err = io.Copy(tar_writer,file)
	if err != nil {
		return err
	}

	return nil
}

// Untar will take a source tar and untar into the destination
func Untar(source, destination string) error {
	// open source
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	// gzip reader
	gz_reader, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	defer gz_reader.Close()

	//tar reader
	tar_reader := tar.NewReader(gz_reader)
	
	// loop over reader
	for {
		header, err := tar_reader.Next()

		switch{
			//exiting when you hit the end or an error
		case err == io.EOF :
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		// create the file path
		destination_path := filepath.Join(destination,header.Name)

		// based on the type of object take the right steps
		switch header.Typeflag {
		case tar.TypeDir:
			CreateDir(destination_path)
		case tar.TypeReg:
			CopyToNewFile(tar_reader,destination_path,header.Mode)
		}
	}
}