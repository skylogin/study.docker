# pull
docker image pull gihyodocker/echo:latest

# start 
docker container run -t -p 9000:8080 gihyodocker/echo:latest

# test
curl http://localhost:9000/

# stop
docker stop $(docker container ls -q)


------------------------------------------------------------------

# build
docker image build -t example/echo:latest .

# query
docker image ls

# start
docker container run example/echo:latest
# start(background)
docker container run -d example/echo:latest

# stop (filter)
docker container stop $(docker container ls --filter "ancestor=exmaple/echo" -q)


# port forwarding
docker container run -d -p 9000:8080 example/echo:latest
# port forwarding (skip
docker container run -d -p 8080 example/echo:latest
