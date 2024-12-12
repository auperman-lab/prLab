package manager

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

type FTPClient struct {
	Address   string
	Username  string
	Password  string
	LeaderURL string
}

// NewFTPClient initializes the FTPClient
func NewFTPClient(address string, username string, password string, leaderURL string) *FTPClient {
	return &FTPClient{Address: address, Username: username, Password: password, LeaderURL: leaderURL}
}

// PollAndProcess retrieves the file and ID every 30 seconds
func (client *FTPClient) PollAndProcess() {
	for {
		conn, err := ftp.Dial(client.Address, ftp.DialWithTimeout(5*time.Second))
		if err != nil {
			log.Printf("Failed to connect to FTP server: %s", err)
			time.Sleep(30 * time.Second)
			continue
		}

		err = conn.Login(client.Username, client.Password)
		if err != nil {
			log.Printf("Failed to login to FTP server: %s", err)
			time.Sleep(30 * time.Second)
			continue
		}

		entries, err := conn.List("/")
		if err != nil {
			log.Printf("Failed to list files: %s", err)
			time.Sleep(30 * time.Second)
			continue
		}

		for _, entry := range entries {
			if entry.Type == ftp.EntryTypeFile {
				log.Printf("Processing file: %s", entry.Name)

				resp, err := conn.Retr(entry.Name)
				if err != nil {
					log.Printf("Failed to retrieve file: %s", err)
					continue
				}

				buf := new(bytes.Buffer)
				_, _ = buf.ReadFrom(resp)

				err = client.sendToLeader(entry.Name, buf.Bytes())
				if err != nil {
					log.Printf("Failed to send file to leader: %s", err)
					continue
				}

				log.Printf("File %s sent successfully to leader", entry.Name)

				err = conn.Delete(entry.Name)
				if err != nil {
					log.Printf("Failed to delete file: %s", err)
				}
				resp.Close()
			}
		}
		conn.Quit() // Explicitly close the connection after processing

		time.Sleep(30 * time.Second)
	}
}

func (client *FTPClient) sendToLeader(fileName string, fileData []byte) error {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("image", fileName)
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}
	_, err = io.Copy(part, bytes.NewReader(fileData))
	if err != nil {
		return fmt.Errorf("failed to write file data to form: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/products/20/upload", client.LeaderURL), body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	clientResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer clientResp.Body.Close()

	if clientResp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(clientResp.Body)
		return fmt.Errorf("leader responded with status %d: %s", clientResp.StatusCode, responseBody)
	}

	return nil
}

func (client *FTPClient) UpdateURL(newURL string) {
	client.LeaderURL = newURL
}
