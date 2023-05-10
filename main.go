package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./checksum <file_or_directory_path>")
		return
	}

	path := os.Args[1]

	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Calculating checksums for %s...\n", path)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		calculateChecksumsForDir(path)
	case mode.IsRegular():
		calculateChecksumsForFile(path)
	default:
		fmt.Println("Unsupported input. Please provide a file or directory path.")
	}
}

func calculateChecksumsForDir(dirPath string) {
	md5Hash := md5.New()
	sha1Hash := sha1.New()
	sha256Hash := sha256.New()
	sha512Hash := sha512.New()

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(io.MultiWriter(md5Hash, sha1Hash, sha256Hash, sha512Hash), file)
		return err
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printChecksum("MD5", md5Hash)
	printChecksum("SHA-1", sha1Hash)
	printChecksum("SHA-256", sha256Hash)
	printChecksum("SHA-512", sha512Hash)
}

func calculateChecksumsForFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	md5Hash := md5.New()
	sha1Hash := sha1.New()
	sha256Hash := sha256.New()
	sha512Hash := sha512.New()

	_, err = io.Copy(io.MultiWriter(md5Hash, sha1Hash, sha256Hash, sha512Hash), file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printChecksum("MD5    ", md5Hash)
	printChecksum("SHA-1  ", sha1Hash)
	printChecksum("SHA-256", sha256Hash)
	printChecksum("SHA-512", sha512Hash)
}

func printChecksum(name string, hash hash.Hash) {
	fmt.Printf("%s:  %s\n", name, hex.EncodeToString(hash.Sum(nil)))
}
