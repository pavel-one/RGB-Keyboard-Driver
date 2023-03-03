wails.dev:
	sudo wails dev
build:
	wails build

build.frontend:
	cd frontend && npm install && npm run build

build.manual.dev:
	make build.frontend
	go build -tags dev -gcflags "all=-N -l" -o build/bin/kd87a
build.manual:
	make build.frontend
	go build -tags desktop,production -ldflags "-w -s -H windowsgui" -o build/bin/kd87a