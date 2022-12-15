SHELL=cmd.exe
FRONT_END_BINARY=frontApp.exe

## build_front: builds the frone end binary
build_front:
	@echo Building front end binary...
	chdir ..\front-end && set CGO_ENABLED=0&& set GOOS=windows&& go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo Done!

## start: starts the front end
start: build_front
	@echo Starting front end
	chdir ..\front-end && start /B ${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo Stopping front end...
	@taskkill /IM "${FRONT_END_BINARY}" /F
	@echo "Stopped front end!"
