##### 1.Mac下编译Linux, Windows(如编译到树莓派4B则GOARCH=arm64)
```code
  # Mac
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o cmd/dev/pea_server -ldflags="-linkmode internal -X main.BuildEnv=dev"
  # Linux
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/qa/pea_server -ldflags="-linkmode internal -X main.BuildEnv=qa"
  # Windows
  go generate
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o cmd/dev/pea_server.exe -ldflags="-linkmode internal -X main.BuildEnv=dev"
```
##### 2.Linux下编译Mac, Windows(如编译到树莓派4B则GOARCH=arm64)
```code
  # Mac
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o cmd/dev/pea_server -ldflags="-linkmode internal -X main.BuildEnv=dev"
  # Linux
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/dev/pea_server -ldflags="-linkmode internal -X main.BuildEnv=dev"
  # Windows
  go generate
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o cmd/dev/pea_server.exe -ldflags="-linkmode internal -X main.BuildEnv=dev"
```
##### 3.Windows下编译Mac, Linux(如编译到树莓派4B则GOARCH=arm64)
windows环境的编译语句需要逐行执行
```code
  # Mac
  SET CGO_ENABLED=0
  SET GOOS=darwin
  SET GOARCH=amd64
  go build -o cmd/dev/pea_server -ldflags="-linkmode internal -X main.BuildEnv=dev"
  # Linux
  SET CGO_ENABLED=0
  SET GOOS=linux
  SET GOARCH=amd64
  go build -o cmd/dev/pea_server -ldflags="-linkmode internal -X main.BuildEnv=dev"
  # Windows
  SET CGO_ENABLED=0
  SET GOOS=windows
  SET GOARCH=amd64
  go generate
  go build -o cmd/dev/pea_server.exe -ldflags="-linkmode internal -X main.BuildEnv=dev"
```