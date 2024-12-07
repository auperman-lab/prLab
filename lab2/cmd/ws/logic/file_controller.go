package logic

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"
)

func CreateFile(name string) (error, *os.File) {
	file, err := os.Create(fmt.Sprintf("./cmd/ws/repo/%s.json", name))
	if err != nil {
		slog.Error("Failed to create file ", "error", err, "file name", name)
		return err, nil
	}
	fmt.Println("File is created successfully.") //print the s
	return nil, file
}

func GetFile(name string) *os.File {
	file, err := os.OpenFile(fmt.Sprintf("./cmd/ws/repo/%s", name), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		slog.Error("Failed to get file ", "error", err, "file name", name)
		return nil
	}
	return file
}

func WriteMessage(f *os.File, message string) {
	msg := []byte(fmt.Sprintf("%s\n", message))
	_, err := f.Write(msg)
	if err != nil {
		if err.Error() == "EOF" {
		} else {
			slog.Error("Failed to write message to file", "error", err, "file name", f.Name())
		}
	}
}

func ReadRooms() []string {
	dirPath := "./cmd/ws/repo"

	dir, err := os.Open(dirPath)
	if err != nil {
		slog.Error("Failed to open dir", "error", err, "dirPath", dirPath)
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	roomsFiles := make([]string, len(files))
	for i, file := range files {
		fmt.Println(file.Name())
		roomsFiles[i] = file.Name()
	}
	return roomsFiles
}

func ReadChat(f *os.File) string {
	var buffer [512]byte
	var content string

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		log.Println("Error resetting file pointer:", err)
		return ""
	}

	for {
		n, err := f.Read(buffer[:])
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return ""
		}
		content += string(buffer[:n])

		if err == nil {
			break
		}
	}

	lines := strings.Split(content, "\n")
	var chatHistory string

	for _, line := range lines {
		if line != "" {
			chatHistory += line + "\n"
		}
	}

	return chatHistory
}
