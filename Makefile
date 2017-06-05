all:golinks_linux_x86 golinks_linux_x64 golinks_osx_x86 golinks_osx_x64 golinks_windows_x86.exe golinks_windows_x64.exe

golinks_linux_x86:
	GOOS=linux GOARCH=386 go build -o $@

golinks_osx_x86:
	GOOS=darwin GOARCH=386 go build -o $@

golinks_windows_x86.exe:
	GOOS=windows GOARCH=386 go build -o $@

golinks_linux_x64:
	GOOS=linux GOARCH=amd64 go build -o $@

golinks_osx_x64:
	GOOS=darwin GOARCH=amd64 go build -o $@

golinks_windows_x64.exe:
	GOOS=windows GOARCH=amd64 go build -o $@


clean:
	rm golinks_*

.PHONY: golinks_linux_x86 golinks_linux_x64 golinks_osx_x86 golinks_osx_x64 golinks_windows_x86.exe golinks_windows_x64.exe clean
