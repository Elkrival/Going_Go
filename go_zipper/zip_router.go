package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {

	files := []string{"bat_api/_user.json"}
	output := "user.zip"

	if err := ZipFiles(output, files); err != nil {
		panic(err)
	}
	fmt.Println("Zipped file:", output)
}
/*
ZipFiles compresses one or more files into a zip archive
Param1: filename is the output zip file's name
Param2: files is alist of files to add to the zip
*/

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}