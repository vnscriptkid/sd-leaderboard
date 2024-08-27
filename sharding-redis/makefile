up:
	docker compose up -d redis-node-1 redis-node-2 redis-node-3 redis-node-4 redis-node-5 redis-node-6
	make cluster
	docker compose up --build go-app

up_build:
	docker compose up -d --build	

down:
	docker compose down --remove-orphans --volumes

cluster:
	docker compose exec redis-node-1 redis-cli --cluster create \
	redis-node-1:7001 redis-node-2:7002 redis-node-3:7003  redis-node-4:7004 redis-node-5:7005 redis-node-6:7006 \
		--cluster-replicas 1

nodes:
	docker compose exec -it redis-node-1 redis-cli -p 7001 cluster nodes

# possible err: (error) MOVED 6918 172.25.0.2:7002
# -c: enable cluster mode, automatically redirect to the correct node
cli:
	docker compose exec -it redis-node-1 redis-cli -c -p 7001

logs:
	docker compose logs -f go-app

ping:
	docker-compose exec redis-node-5 redis-cli -c -p 7005 ping
	docker-compose exec redis-node-6 redis-cli -c -p 7006 ping
	docker-compose exec redis-node-1 redis-cli -c -p 7001 ping
	docker-compose exec redis-node-2 redis-cli -c -p 7002 ping
	docker-compose exec redis-node-3 redis-cli -c -p 7003 ping
	docker-compose exec redis-node-4 redis-cli -c -p 7004 ping

