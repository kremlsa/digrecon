set CGO_ENABLED=0
set GOOS=linux
go build -a -installsuffix cgo -ldflags "-w -s" -o main ./cmd