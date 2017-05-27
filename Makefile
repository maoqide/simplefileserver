build:
	docker build -t maoqide/simplefileserver .
run:
	docker run -d -p 18080:18080 -v `pwd`:/root/servedir maoqide/simplefileserver -serveDir /root/servedir
