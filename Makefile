build:
	@go build -buildmode=c-shared
	@rm -f main.h
