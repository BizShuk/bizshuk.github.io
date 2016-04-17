# Http & Https

# Http

### Headers
##### Authentication
<user>:<password> ,  encode by base64






# Https
[ref]
- [傳輸層安全-認識SSL](http://blog.yogo.tw/2009/11/ssl.html)
- [How Https works](http://robertheaton.com/2014/03/27/how-does-https-actually-work/)

HTTPS Handshake
1. Hello , client say hello to server including the various cipher suites and maximum SSL version that it supports.
2. Certificate Exchange , server say hello to client with certificate, public key, cipher suite, client have, and SSL version.
3. Key Exchange , client generate a secret key and encrypt with server public key
4. Communicate with each other by secrete key

ps: 
- secret key is called session key , too.
- cipher suites = encryption algo. list


Trouble
- server private key is not private anymore , it will be eavesdropped by hacker , pretenting real server.
