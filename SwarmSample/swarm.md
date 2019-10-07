# registry (1)
port 5000
/var/lib/registry

# manager (1)
dind

# worker (3)
dind


# run
docker-compose -up -d

# check
docker container ls

---


# manager init
docker container exec -it manager docker swarm init

# registry worker01
docker container exec -it worker01 docker swarm join \
  --token SWMTKN-1-5nponat5xqptrj4jyjjlrqgias1woltc81xa0xh7m33m7goyyr-d0t8c1ixdijjro8fnzxo6k8tp manager:2377
# registry worker02
docker container exec -it worker02 docker swarm join \
  --token SWMTKN-1-5nponat5xqptrj4jyjjlrqgias1woltc81xa0xh7m33m7goyyr-d0t8c1ixdijjro8fnzxo6k8tp manager:2377
# registry worker03
docker container exec -it worker03 docker swarm join \
  --token SWMTKN-1-5nponat5xqptrj4jyjjlrqgias1woltc81xa0xh7m33m7goyyr-d0t8c1ixdijjro8fnzxo6k8tp manager:2377
