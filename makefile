.PHONY: server stop client

server:
	go build go_server.go && ./go_server &
	python py_server.py &

stop:
	@echo "Closing servers..."
	@pkill -f go_server || true
	@pkill -f py_server.py || true

client:
	@echo "Calling Python API..."
	curl http://127.0.0.1:8000/py/hello

	@echo "Calling Go API..."
	curl http://127.0.0.1:8080/go/hello