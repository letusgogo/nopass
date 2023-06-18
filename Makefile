MAIN_PATH:=main.go
PROJECT:= nopass

.PHONY: darwin
darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o ${PROJECT} ${MAIN_PATH}

.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ${PROJECT} ${MAIN_PATH}
.PHONY: install
install: darwin
	cp ${PROJECT} ${GOPATH}/bin/


.PHONY: clean
clean:
	rm ${PROJECT}