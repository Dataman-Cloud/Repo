package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

// fileExists return flag whether a given file exists
// and operation error if an unclassified failure occurs.
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	log.Println("AgrsWithProg: ", argsWithProg)
	log.Println("ArgsWithoutProg: ", argsWithoutProg)

	if len(argsWithoutProg) < 1 {
		log.Panicln("No engouth params")
	}

	filePath := argsWithoutProg[0]
	log.Println("Target file path", filePath)

	exists, err := FileExists(filePath)
	if err != nil {
		log.Panicln(err)
	}

	if !exists {
		log.Panicln("File not found")
	}

	fileInfo, err := os.Stat(filePath)

	fileName := fileInfo.Name()
	log.Println("FileName: ", fileName)

	file, err := os.Open(filePath)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	hash := md5.New()

	if _, err = io.Copy(hash, file); err != nil {
		log.Panicln(err)
	}

	md5Code := hash.Sum(nil)
	md5CodeStr := fmt.Sprintf("%x", md5Code)
	log.Printf("md5Code of %s is %s ", filePath, md5CodeStr)

	md5FileName := fileName + ".md5"

	md5File, err := os.OpenFile(md5FileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Panicln("Error for open md5File to write md5 code in")
	}
	defer md5File.Close()

	md5File.Write([]byte(md5CodeStr))
	log.Printf("Write md5code: %s to %s success", md5CodeStr, md5FileName)

	sha256hash := sha256.New()
	sha256Code := sha256hash.Sum(nil)
	sha256CodeStr := fmt.Sprintf("%x", sha256Code)
	log.Printf("sha256code of %s is %s ", filePath, sha256CodeStr)

	sha256FileName := fileName + ".sha256"
	sha256File, err := os.OpenFile(sha256FileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Panicln("Error for open sha256File to write sha256 code in")
	}
	defer sha256File.Close()

	sha256File.Write([]byte(sha256CodeStr))
	log.Println("Write sha256code: %s to %s success", sha256CodeStr, sha256FileName)
}
