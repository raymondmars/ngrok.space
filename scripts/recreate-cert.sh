#!/bin/bash
# export NGROK_DOMAIN=ngrok.quest
#
#参考：https://www.cnblogs.com/yjmyzz/p/openssl-tutorial.html
#https://www.cnblogs.com/jackluo/p/13841286.html
#TLS：传输层安全协议 Transport Layer Security的缩写
#SSL：安全套接字层 Secure Socket Layer的缩写
#TLS与SSL对于不是专业搞安全的开发人员来讲，可以认为是差不多的，这二者是并列关系，详细差异见 http://kb.cnblogs.com/page/197396/
#KEY 通常指私钥。
#CSR 是Certificate Signing Request的缩写，即证书签名请求，这不是证书，可以简单理解成公钥，生成证书时要把这个提交给权威的证书颁发机构。
#CRT 即 certificate的缩写，即证书。
#X.509 是一种证书格式.对X.509证书来说，认证者总是CA或由CA指定的人，一份X.509证书是一些标准字段的集合，这些字段包含有关用户或设备及其相应公钥的信息。
#什么是 SAN
#SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。

mkdir -p ./certs

#1，生成CA证书
openssl genrsa -out ./certs/ca.key 2048 
openssl req -x509 -new -nodes -key ./certs/ca.key -subj "/CN=ngrok.quest" -days 5000 -out ./certs/ca.crt 

#2，生成证书签名请求文件（公钥）,并且指定最终是生成SAN类型的证书
openssl genrsa -out ./certs/server.key 2048 
openssl req -new -sha256 \
    -key ./certs/server.key \
    -subj "/CN=ngrok.quest" \
    -reqexts SAN \
    -config <(cat ./scripts/openssl.cnf \
        <(printf "[SAN]\nsubjectAltName=DNS:localhost,DNS:ngrok.quest")) \
    -out ./certs/server.csr 

#3，生成 x.509 格式证书
openssl x509 -req -days 365000 \
    -in ./certs/server.csr -CA ./certs/ca.crt -CAkey ./certs/ca.key -CAcreateserial \
    -extfile <(printf "subjectAltName=DNS:localhost,DNS:ngrok.quest") \
    -out ./certs/server.crt

cp ./certs/ca.crt ./assets/client/tls/ngrokroot.crt 
cp ./certs/server.crt ./assets/server/tls/snakeoil.crt 
cp ./certs/server.key ./assets/server/tls/snakeoil.key 

# openssl req -new -key ./certs/rootCA.key -subj "/CN=$NGROK_DOMAIN" -reqexts SAN -config ./scripts/openssl.cnf -out ./certs/server.csr && \
