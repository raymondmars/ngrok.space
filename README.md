
### This is a fork version from [ngrok](https://github.com/inconshreveable/ngrok)     
Because the [original version](https://github.com/inconshreveable/ngrok) is no longer maintained, So I forked that project and made the following changes:    
1. use [Golang Project Layout](https://github.com/golang-standards/project-layout) to refactor project structure       
2. init these modules by go mod and upgrade golang to the latest version         
3. create a user system and add a server auth function based on user information             
4. add a newly self-sign certificate and create a new Makefile to do build and deploy               
5. add a Dockerfile to the server and let it can run in a docker image  
6. create a new Makefile to make everything            
6. modify client code, add a gin server, Let the user decide whether to build a version with its own client-server            

#### Why did I add a gin server to the client as the default server? 
If someone wants a function out of the box, or he is not familiar with writing code,     
He/She can build a client version with its own server(like this: GOBIN= GOOS=darwin GOARCH=amd64 go build -tags release -ldflags "-s -w -X 'main.WithLocalServer=1', just set build flags to: main.WithLocalServer=1).        
He only needs to start the client, then creates a site file under the startup directory, and then put the website template file in it, and he immediately has his own website :)       


[to be continue...]      

### ngrok - Introspected tunnels to localhost    
### ”I want to expose a local server behind a NAT or firewall to the internet.”
![](https://ngrok.com/static/img/overview.png)

### What is ngrok?
ngrok is a reverse proxy that creates a secure tunnel from a public endpoint to a locally running web service.
ngrok captures and analyzes all traffic over the tunnel for later inspection and replay.


### Production Use

**DO NOT RUN THIS VERSION OF NGROK (1.X) IN PRODUCTION**. Both the client and server are known to have serious reliability issues including memory and file descriptor leaks as well as crashes. There is also no HA story as the server is a SPOF. You are advised to run 2.0 for any production quality system. 

### What can I do with ngrok?
- Expose any http service behind a NAT or firewall to the internet on a subdomain of ngrok.com
- Expose any tcp service behind a NAT or firewall to the internet on a random port of ngrok.com
- Inspect all http requests/responses that are transmitted over the tunnel
- Replay any request that was transmitted over the tunnel


### What is ngrok useful for?
- Temporarily sharing a website that is only running on your development machine
- Demoing an app at a hackathon without deploying
- Developing any services which consume webhooks (HTTP callbacks) by allowing you to replay those requests
- Debugging and understanding any web service by inspecting the HTTP traffic
- Running networked services on machines that are firewalled off from the internet



