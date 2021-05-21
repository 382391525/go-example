package ltcp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// Encode 消息编码
func Encode(message string) ([]byte, error) {
	//消息的长度,装换为int32 占用4字节
	length := int32(len(message))
	bb := new(bytes.Buffer)
	//写入消息头
	err := binary.Write(bb, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息体
	err = binary.Write(bb, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

// Decode 消息解码
func Decode(reader *bufio.Reader) (string, error) {
	//消息体的长度
	var length int32
	//读取前4个字节的长度
	lengthByte, _ := reader.Peek(4)
	lengthBuffer := bytes.NewBuffer(lengthByte)
	err := binary.Read(lengthBuffer, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

func serverProcess(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

//服务端
func server() {
	listen, err := net.Listen("tcp", ":30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		//监听连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go serverProcess(conn)
	}
}

//客户端
func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
