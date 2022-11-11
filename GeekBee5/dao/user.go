package dao

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 假数据库，用 map 实现
var database = map[string]string{}

var PhoneNum = map[string]string{}

func LoadAll() {
	LoadDatabase()
	LoadphoneNum()
	LoadComments()
}

func LoadDatabase() {
	println("start Load")
	fileName := "Database.txt"
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
		database[part[0]] = part[1]

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

func LoadphoneNum() {
	println("start Load phone Num")
	fileName := "PhoneNum.txt"
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
		PhoneNum[part[0]] = part[1]
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

func SaveAll() {
	SaveDatabse()
	SavePhoneNum()
}

func SaveDatabse() {
	fileName := "Database.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		println("OpenFile wrong")
		return
	}
	writer := bufio.NewWriter(file)
	for name, password := range database {
		writer.WriteString(name + " " + password + "\n")
	}
	writer.Flush()
	defer file.Close()
}

func SavePhoneNum() {
	fileName := "PhoneNum.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		println("OpenFile wrong")
		return
	}
	writer := bufio.NewWriter(file)
	for name, phoneNum := range PhoneNum {
		writer.WriteString(name + " " + phoneNum + "\n")
	}
	writer.Flush()
	defer file.Close()
}

func AddUser(username, password, phone string) {
	database[username] = password
	fileName := "Database.txt"
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(username+" "+password+"\n"), n)
	}
	defer f.Close()

	PhoneNum[username] = phone
	fileName = "PhoneNum.txt"
	f, err = os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(username+" "+phone+"\n"), n)
	}
	defer f.Close()
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}

func DelUser(username string) {
	delete(database, username)
	delete(PhoneNum, username)
	return
}

func SelectPhoneFromUsername(username, phone string) bool {
	if PhoneNum[username] == phone {
		return true
	} else {
		return false
	}
}
