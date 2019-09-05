protoc --plugin=protoc-gen=protoc-gen-go.exe --go_out=./goout *.proto
git add *.proto
git add goout/*.go
git commit -m proto
pause