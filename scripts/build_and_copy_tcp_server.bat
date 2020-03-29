set GOOS=linux
set GOARCH=arm
set GOARM=5
set binaryName=test_tcp_server_9002
go build -o %binaryName% .
scp .\%binaryName% pi@192.168.0.34:/home/pi/tcp
rm %binaryName%

