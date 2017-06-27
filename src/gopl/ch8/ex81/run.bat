set TZ=Asia/Shanghai
start /b go run clock.go -p 8001
set TZ=Asia/Tokyo
start /b go run clock.go -p 8002
set TZ=US/Eastern
start /b go run clock.go -p 8003
set TZ=Europe/London
start /b go run clock.go -p 8005
start go run clockwall.go Beijing=localhost:8001 NewYork=localhost:8003 London=localhost:8005 Tokyo=localhost:8002
start go run termbox_clock.go Beijing=localhost:8001 NewYork=localhost:8003 London=localhost:8005 Tokyo=localhost:8002
