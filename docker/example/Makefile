VERSION=2.6.3
REGISTRY=hub.longzhu.com
NAME=relationship-web
PORT=40016
GRCPNAME=go.micro.srv.relationship
COPYTO=${COPY_TO}
run:
	go run main.go --server_address=${BIND_IP}:$(PORT)
run-bg:
	nohup ./$(NAME) --server_address=${BIND_IP}:$(PORT) 2>&1 > $(NAME).nohup &
run-grpc:
	go run main.go --server_name=$(GRCPNAME) --server_version=$(VERSION) --client=grpc --server=grpc --transport=grpc --server_address=${BIND_IP}:$(PORT)
run-grpc-bg:
	nohup ./$(NAME) --server_name=$(GRCPNAME) --server_version=$(VERSION) --client=grpc --server=grpc --transport=grpc --server_address=${BIND_IP}:$(PORT) 2>&1 > $(NAME).nohup &
build:
	GOARCH=amd64 GOOS=linux go build -o $(NAME)
clean:
	rm -rf ./$(NAME) ./$(NAME).nohup
kill:
	killall $(NAME)
build-image:
	GOARCH=amd64 GOOS=linux go build -o $(NAME)
	docker build -f Dockerfile -t $(REGISTRY)/micro/$(NAME):$(VERSION) .
	docker push $(REGISTRY)/micro/$(NAME):$(VERSION)
	rm -rf $(NAME)
copy-to-dev:
	scp ./$(NAME) root@$(COPYTO):/home/micro/$(NAME)/
	scp ./Makefile root@$(COPYTO):/home/micro/$(NAME)/
	scp ./log*.xml root@$(COPYTO):/home/micro/$(NAME)/

copy-to-test:
	scp -P58422 ./$(NAME) root@$(COPYTO):/home/micro/$(NAME)/
	scp -P58422 ./Makefile root@$(COPYTO):/home/micro/$(NAME)/
	scp -P58422 ./log*.xml root@$(COPYTO):/home/micro/$(NAME)/
