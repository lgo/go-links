all: golinks_linux golinks_osx golinks_windows.exe

golinks_linux:
	GOOS=linux GOARCH=386 go build -o $@

golinks_osx:
	GOOS=darwin GOARCH=386 go build -o $@

golinks_windows.exe:
	GOOS=windows GOARCH=386 go build -o $@

clean:
	rm golinks_osx golinks_linux golinks_windows.exe

.PHONY: golinks_linux golinks_osx golinks_windows.exe clean
