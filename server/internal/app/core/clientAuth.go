package core

import (
	"bufio"
	"log"
	"os"
	"strings"

	"raymond.com/common/msg"
)

var authTokens []string

//Custom client authorization
type ClientAuth interface {
	IsValid() bool
}

//Default use file to store username and password and to check client auth info
type fileAuthDb struct {
	Auth *msg.Auth
}

func NewClientAuth(auth *msg.Auth) ClientAuth {
	return &fileAuthDb{Auth: auth}
}

func (fa *fileAuthDb) IsValid() bool {
	return true
	//To checkout user token is valid
	if len(fa.Auth.User) < 16 || !strings.Contains(fa.Auth.User, "@") {
		log.Println("client user token too short or invalid")
	}

	if len(authTokens) == 0 {
		tokens, err := fa.readLines("./tokens.txt")
		if err != nil {
			log.Fatalf("failed opening directory: %s", err)
		} else {
			authTokens = tokens
			if len(authTokens) == 0 {
				return false
			}
		}
	}

	for _, token := range authTokens {
		// fmt.Println("======")
		// fmt.Println(fa.Auth.User)
		// fmt.Println(token)
		if !strings.Contains(token, "@") {
			log.Printf("invalid server token: %s\n, token must contains @", token)
			return false
		}
		serverToken := strings.Split(strings.TrimSpace(token), "@")
		clientToken := strings.Split(strings.TrimSpace(fa.Auth.User), "@")
		if serverToken[0] == clientToken[0] && serverToken[1] == clientToken[1] {
			return true
		}

	}
	return false
}

func (fa *fileAuthDb) readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
