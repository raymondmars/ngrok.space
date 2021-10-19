
macos-client:
	GOBIN= GOOS=darwin GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o ../deploy/bin/macos_amd64/ngrok .

linux-client:
	GOBIN= GOOS=linux GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o ../deploy/bin/linux_amd64/ngrok .

windows-client:
	GOBIN= GOOS=windows GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o ../deploy/bin/windows_amd64/ngrok.exe .


macos-raymond-with-server-client:
	GOBIN= GOOS=darwin GOARCH=amd64 go build -tags release -ldflags "-s -w -X 'main.WithLocalServer=1' -X 'main.DefaultSubDomain=raymond' -X 'main.DefaultAuthtoken=c5n45fucie6jogm2ik9g_1zi1Ge97BJ0ibY7iAQP8Vjwl5lL'" -a -o ../deploy/bin/macos_amd64/ngrok-with-server .

windows-raymond-with-server-client:
	GOBIN= GOOS=windows GOARCH=amd64 go build -tags release -ldflags "-s -w -X 'main.WithLocalServer=1' -X 'main.DefaultSubDomain=raymond' -X 'main.DefaultAuthtoken=c5n45fucie6jogm2ik9g_1zi1Ge97BJ0ibY7iAQP8Vjwl5lL'" -a -o ../deploy/bin/windows_amd64/ngrok-server.exe .
