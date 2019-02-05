rm src/simple/*.go 2> /dev/null
protoc -I src/ --go_out=src src/simple/simple.proto