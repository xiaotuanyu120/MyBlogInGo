package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func ExecuteDir() string {
	// return binary execute dir
	ret, _ := os.Executable()
	return filepath.Join(ret, "..")
}

/*
CreateNonExistDir

create not existed directory
*/
func CreateNonExistDir(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Printf("Folder does not exist >> [%s]", dir)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		log.Printf("Folder create success >> [%s]", dir)
	}
	return nil
}

/*
ReadFile

read content from file
*/
func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string
	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}
	return textLines, nil
}

/*
AppendToFile

If the file doesn't exist, create it, or append to the file
*/
/*
>>>>>>>>>>>>>>>>>>>>>> USELESS, SO COMMENT <<<<<<<<<<<<<<<<<<<<<<<<<<<
func AppendToFile(filePath string, content []byte) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("open file failed, message: [" + err.Error() + "]")
	}
	if _, err := f.Write(content); err != nil {
		err := f.Close()
		if err != nil {
			return errors.New("write file failed and close file failed, message: [" + err.Error() + "]")
		}
		return errors.New("write file failed, message: [" + err.Error() + "]")
	}
	if err := f.Close(); err != nil {
		return errors.New("close file failed, message: [" + err.Error() + "]")
	}

	return nil
}
*/
