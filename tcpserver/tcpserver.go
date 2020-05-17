package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

	for {
		n, err = c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) //데이터 출력

		_, err = c.Write(data[:n]) // 클라이언트로 데이터를 보낸다.
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}

func main() {
	ln, err := net.Listen("tcp", "8000") // tcp 프로토콜에 8000 포트로 연결을 받음 , 프로토콜,ip주소, 포트번호를 설정하여 네트워크 연결 대기
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close() // main 함수가 끝나기 직전에 연결을 닫음

	for {
		conn, err := ln.Accent() // 클라이언트가 연결되면 tcp 연결을 리턴
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer conn.Close() // main 함수가 끝나기 직전에 tcp 연결을 닫음

		go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 실행
	}
}
