# Database-2023Autumn
written in Golang
migrate from https://gitee.com/solaris61/database-2023-autumn

## build and run
```
git clone https://github.com/ppolariss/Database-2023Autumn.git
cd Database-2023Autumn
# install swag and generate docs
go install github.com/swaggo/swag/cmd/swag@latest
swag init --parseDependency --parseDepth 1 # to generate the latest docs, this should be run before compiling
# build for debug
go build -o main.exe
# run
./main.exe
```
