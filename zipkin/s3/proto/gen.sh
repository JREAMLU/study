protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. s3.proto

#protoc -I ./ --go_out=plugins=micro,service_name=go.micro.srv.greeter:. ./greeter.proto
# 因为ios中类型的默认值是nil，所以要去掉生成*.proto.go文件中omitempty。
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"