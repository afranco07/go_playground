package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func recursiveFiles(path string, connection net.Conn) []string {
	var returnArr []string
	returnArr = recursiveFiles2(path, &returnArr, connection)
	return returnArr
}

func recursiveFiles2(path string, returnList *[]string, connection net.Conn) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		newPath := path + "/" + file.Name()
		if file.IsDir() {
			recursiveFiles2(newPath, returnList, connection)
		} else {
			*returnList = append(*returnList, newPath)
			connection.Write([]byte(newPath + "\n"))
		}
	}
	return *returnList
}

func listFiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		newPath := path + "/" + file.Name()
		if file.IsDir() {
			listFiles(newPath)
		} else {
			fmt.Println(newPath)
		}
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("ERROR READING:", err.Error())
	}

	s := string(buf[:reqLen])
	fmt.Println("MESSAGE =", s, "WITH LENGTH =", reqLen)
	listFiles(s)
	// var newArr []string
	recursiveFiles(s, conn)
	// fmt.Println("RETURNED ARR", newArr)
	conn.Write([]byte("Message received\n"))
	conn.Write([]byte("Message received, again!\n"))
	conn.Close()
}

func main() {
	fmt.Println("Starting server on port 8080 (localhost)...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("ERROR =>", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("2ND ERROR =>", err)
		}
		go handleRequest(conn)
	}
}
