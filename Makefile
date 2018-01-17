EXE = gorestable


all: ${EXE}
	@go build -v ./cmd/${EXE}
	@mv ./cmd/${EXE}/${EXE} _output/

clean:
	@rm -rf _output

fmt:
	@go fmt ./cmd/${EXE}/*.go
