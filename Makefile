bin/consumer:
	go build -o bin/consumer ./consumer
	chmod +x bin/consumer

bin/producer:
	go build -o bin/producer ./producer
	chmod +x bin/producer

clean:
	rm -rf bin out

fmt:
	go fmt ./...

.PHONY: clean fmt