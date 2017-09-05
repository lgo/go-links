all: golinks_linux_x86 golinks_linux_x64 golinks_osx_x86 golinks_osx_x64 golinks_windows_x86.exe golinks_windows_x64.exe

dev: export ADMIN_KEY := test
dev:
	go run main.go

golinks_linux_x86:
	GOOS=linux GOARCH=386 go build -o build/$@

golinks_osx_x86:
	GOOS=darwin GOARCH=386 go build -o build/$@

golinks_windows_x86.exe:
	GOOS=windows GOARCH=386 go build -o build/$@

golinks_linux_x64:
	GOOS=linux GOARCH=amd64 go build -o build/$@

golinks_osx_x64:
	GOOS=darwin GOARCH=amd64 go build -o build/$@

golinks_windows_x64.exe:
	GOOS=windows GOARCH=amd64 go build -o build/$@


clean:
	rm -rf build/

.PHONY: golinks_linux_x86 golinks_linux_x64 golinks_osx_x86 golinks_osx_x64 golinks_windows_x86.exe golinks_windows_x64.exe clean
