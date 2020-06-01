run:
	#go-webserver
	docker rm -f go-webserver || true
	docker build -t poc/webserver webserver
	docker run -d \
		--name go-webserver \
		-p "2112:2112" \
		-m 512m \
		--cpus="0.5" \
		--network="prom-poc" \
		poc/webserver
	#prometheus
	docker rm -f prom || true
	docker run -d \
		--name prom \
    -p 9090:9090 \
		-m 512m \
		--cpus="0.5" \
		--network="prom-poc" \
		-v $(shell pwd)/prom:/etc/prometheus \
    prom/prometheus 
	#alertmanager
	docker rm -f alertmanager || true
	docker run -d \
		--name alertmanager \
		-m 512m \
		--cpus="0.5" \
		--network="prom-poc" \
		-v $(shell pwd)/alertmanager:/etc/alertmanager \
    prom/alertmanager


