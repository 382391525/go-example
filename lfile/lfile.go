package lfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FileTest() {
	//创建文件
	file, err := os.Create("./lfile/a.txt")
	if err != nil {
		return
	}
	//关闭文件句柄
	defer file.Close()
	fmt.Println("create file ok")
}

func FileOpen() {
	file, err := os.OpenFile("./lfile/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString("abc123")
	_, err = file.WriteString("abc123")
	if err != nil {
		fmt.Println(err)
		return
	}
	//seek, err := file.Seek(5, io.SeekEnd)
	//if err != nil {
	//	return
	//}
	//fmt.Println(seek)
	//
	//_, err = file.WriteString("edfg")

}

func FileBuf() {
	file, err := os.OpenFile("./lfile/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完毕")
			} else {
				fmt.Println("读取出错:", err)
			}
			return
		}
		fmt.Println(string(buf))
	}
}

func FileCopy() {
	//读取文件
	readFile, err := os.Open("./lfile/a.txt")
	if err != nil {
		fmt.Println("Open err")
		return
	}
	defer readFile.Close()
	//创建一个写文件
	createFile, err := os.Create("./lfile/b.txt")
	if err != nil {
		fmt.Println("Create err")
		return
	}
	defer createFile.Close()
	//创建缓冲区
	buf := make([]byte, 4096)
	for {
		n, err := readFile.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("文件 读完")
			return
		}
		_, err = createFile.Write(buf[:n])
		if err != nil {
			fmt.Println("文件 写 错误:", err)
			return
		}
	}

}

// OpenDir 目录相关
func OpenDir() {
	fmt.Println("输入需要查询的目录:")
	var path string
	fmt.Scan(&path)
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err: ", err)
		return
	}
	defer file.Close()
	dirs, err := file.ReadDir(-1)
	if err != nil {
		fmt.Println("ReadDir err: ", err)
		return
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println("目录: ", dir.Name())
		} else {
			fmt.Println("文件: ", dir.Name())
		}
	}
}
