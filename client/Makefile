
macos-client:
	GOBIN= GOOS=darwin GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o ../deploy/bin/macos_amd64/ngrok .

linux-client:
	GOBIN= GOOS=linux GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o ../deploy/bin/linux_amd64/ngrok .

windows-client:
	GOBIN= GOOS=windows GOARCH=amd64 go build -tags release -ldflags "-s -w" -a -o ../deploy/bin/windows_amd64/ngrok.exe .


macos-raymond-with-server-client:
	GOBIN= GOOS=darwin GOARCH=amd64 go build -tags release -ldflags "-s -w -X 'main.WithLocalServer=1' -X 'main.DefaultSubDomain=raymond' -X 'main.DefaultAuthtoken=hellolaojiang@gmail.com:Laojiang12@897'" -a -o ../deploy/bin/macos_amd64/ngrok-with-server .
