MAIN_PATH:=github.com/letusgogo/nopass
PROJECT:= nopass

.PHONY: darwin
darwin:
	GOOS=darwin go build -v -o ${PROJECT} -gcflags -l ${MAIN_PATH}

.PHONY: linux
linux:
	GOOS=linux go build -v -o ${PROJECT} -gcflags -l ${MAIN_PATH}
.PHONY: install
install: darwin
	cp ${PROJECT} ${GOPATH}/bin/

.PHONY: win
win:
	GOOS=windows go build -v -o ${PROJECT}.exe -gcflags -l ${MAIN_PATH}

.PHONY: clean
clean:
	rm ${PROJECT}