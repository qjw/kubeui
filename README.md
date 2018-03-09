# 基于Kubectl命令行的`Dashboard`
类似于[`kubectl proxy dashboard`](https://github.com/kubernetes/dashboard)

目前仅仅实现了基于[`namespace`](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)分类的查看/编辑/删除[`secret`](https://kubernetes.io/docs/concepts/configuration/secret/)

前端使用[`antd-admin`](https://github.com/zuiidea/antd-admin)修改而成，

编译后的前端文件，使用[`go-bindata-assetfs`](https://github.com/elazarl/go-bindata-assetfs)嵌入到二进制程序中，代码在`router/bindata.go`

后端框架使用[`kelly`](https://github.com/qjw/kelly)

# 重新编译前端
``` bash
# 编译前端
npm run build

# 将前端生成的整个dist目录拷贝到当前Project根目录，并且重命令为frontend
cp */dist . && mv dist frontend

# 将前端代码打包到router/bindata.go
go get github.com/jteeuwen/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
# go-bindata-assetfs -o router/bindata.go -pkg router ./frontend/...
go generate

# 重新编译
go build .
```