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
		fmt.Println("Usage: ./checksum <file_or_directory_path> [file_or_directory_path]")
		return
	}

	path1 := os.Args[1]
	path2 := ""

	if len(os.Args) > 2 {
		path2 = os.Args[2]
	}

	fi1, err := os.Stat(path1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if path2 != "" {
		fi2, err := os.Stat(path2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if fi1.Mode().IsRegular() && fi2.Mode().IsRegular() {
			compareFiles(path1, path2)
		} else if fi1.Mode().IsDir() && fi2.Mode().IsDir() {
			compareDirectories(path1, path2)
		} else {
			fmt.Println("Error: Both input paths should be either files or directories.")
		}
	} else {
		fmt.Printf("Calculating checksums for %s...\n", path1)

		md5Hash := md5.New()
		sha1Hash := sha1.New()
		sha256Hash := sha256.New()
		sha512Hash := sha512.New()

		switch mode := fi1.Mode(); {
		case mode.IsDir():
			calculateChecksumsForDir(path1, md5Hash, sha1Hash, sha256Hash, sha512Hash)
		case mode.IsRegular():
			calculateChecksumsForFile(path1, md5Hash, sha1Hash, sha256Hash, sha512Hash)
		default:
			fmt.Println("Unsupported input. Please provide a file or directory path.")
		}
	}
}

func compareFiles(path1, path2 string) {
	md5Hash1, md5Hash2 := md5.New(), md5.New()
	sha1Hash1, sha1Hash2 := sha1.New(), sha1.New()
	sha256Hash1, sha256Hash2 := sha256.New(), sha256.New()
	sha512Hash1, sha512Hash2 := sha512.New(), sha512.New()

	calculateChecksumsForFile(path1, md5Hash1, sha1Hash1, sha256Hash1, sha512Hash1)
	calculateChecksumsForFile(path2, md5Hash2, sha1Hash2, sha256Hash2, sha512Hash2)

	printComparisonResults(path1, path2, md5Hash1, md5Hash2, sha1Hash1, sha1Hash2, sha256Hash1, sha256Hash2, sha512Hash1, sha512Hash2)
}

func compareDirectories(path1, path2 string) {
	md5Hash1, md5Hash2 := md5.New(), md5.New()
	sha1Hash1, sha1Hash2 := sha1.New(), sha1.New()
	sha256Hash1, sha256Hash2 := sha256.New(), sha256.New()
	sha512Hash1, sha512Hash2 := sha512.New(), sha512.New()

	calculateChecksumsForDir(path1, md5Hash1, sha1Hash1, sha256Hash1, sha512Hash1)
	calculateChecksumsForDir(path2, md5Hash2, sha1Hash2, sha256Hash2, sha512Hash2)

	printComparisonResults(path1, path2, md5Hash1, md5Hash2, sha1Hash1, sha1Hash2, sha256Hash1, sha256Hash2, sha512Hash1, sha512Hash2)
}

func calculateChecksumsForDir(dirPath string, md5Hash, sha1Hash, sha256Hash, sha512Hash hash.Hash) {
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

	printChecksum("MD5    ", md5Hash)
	printChecksum("SHA-1  ", sha1Hash)
	printChecksum("SHA-256", sha256Hash)
	printChecksum("SHA-512", sha512Hash)
}

func calculateChecksumsForFile(filePath string, md5Hash, sha1Hash, sha256Hash, sha512Hash hash.Hash) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

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
	fmt.Printf("%s  %s\n", name, hex.EncodeToString(hash.Sum(nil)))
}

func printComparisonResults(path1, path2 string, md5Hash1, md5Hash2, sha1Hash1, sha1Hash2, sha256Hash1, sha256Hash2, sha512Hash1, sha512Hash2 hash.Hash) {
	fmt.Printf("Comparing checksums for %s and %s...\n", path1, path2)

	checksumsMatch := compareHashes(md5Hash1, md5Hash2) &&
		compareHashes(sha1Hash1, sha1Hash2) &&
		compareHashes(sha256Hash1, sha256Hash2) &&
		compareHashes(sha512Hash1, sha512Hash2)

	if checksumsMatch {
		fmt.Println("Checksums match")
	} else {
		fmt.Println("Checksums do not match")
	}
}

func compareHashes(hash1, hash2 hash.Hash) bool {
	return hex.EncodeToString(hash1.Sum(nil)) == hex.EncodeToString(hash2.Sum(nil))
}
