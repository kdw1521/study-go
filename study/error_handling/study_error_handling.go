package errorhandling

import (
	"bufio"
	"fmt"
	"os"
)

const FILENAME string = "data.txt"

// 파일을 읽는다.
func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n') // 한줄을 읽고 string으로 반환
	return line, nil
}

// 파일을 쓴다.
func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}
