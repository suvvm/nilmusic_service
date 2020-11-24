RUN_NAME="nilmusic.service"

help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

build:
	@sh build.sh

run: build
	output/bin/${RUN_NAME}

clean:
	rm -rf output
