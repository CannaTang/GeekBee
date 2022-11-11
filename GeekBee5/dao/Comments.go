package dao

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var Comments = map[string][]string{}

func LoadComments() {
	println("start Load")
	fileName := "Comments.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		println("OpenFile wrong")
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		part := strings.SplitN(line, " ", 2)
		Comments[part[0]] = append(Comments[part[0]], part[1])

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
				return
			}
		}
	}
}

func AddComment(username, comment string) {
	Comments[username] = append(Comments[username], comment)
	SaveComments()
}

func SaveComments() {
	fileName := "Comments.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		println("OpenFile wrong")
		return
	}
	writer := bufio.NewWriter(file)
	for name, _ := range Comments {
		for _, comment := range Comments[name] {
			writer.WriteString(name + " " + comment + "\n")
		}
	}
	writer.Flush()
	defer file.Close()
}

func FindComment(username string) []string {
	var lst = []string{}
	for _, comment := range Comments[username] {
		lst = append(lst, comment)
	}
	return lst
}

func DelComment(username, comment string) {
	var idx int
	for i, element := range Comments[username] {
		if element == comment {
			idx = i
			break
		}
	}
	Comments[username] = append(Comments[username][:idx], Comments[username][idx+1:]...)
	SaveComments()
}
