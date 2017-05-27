# simplefileserver
access the file of a certain directory on your machine via web.

## usage
```bash
go run simplefileserver.go

# you can also specify the directory to serve and the port it listens.
go run simplefileserver.go -serveDir '/path/to/serve' -port port
```

## docker
```bash
# following command can serve current directory you execute it:
docker run -d -p 18080:18080 -v `pwd`:/root/servedir maoqide/simplefileserver -serveDir /root/servedir

# you can also change to any directory you want:
docker run -d -p 18080:18080 -v /path/you/want:/root/servedir maoqide/simplefileserver -serveDir /root/servedir
```
