
* 1. index.html + swagger.json = ReDoc UI
    * 1.1 使用 Dockerfile 构建 docker 镜像
    * 1.2 使用 `make serve-api-doc` 通过 http-server 本地访问

* 2. swagger.json = swagger 官方的 UI
    * 2.1 使用 Dockerfile-swaggerui 构建 docker 镜像
    * 2.2 使用 [https://petstore.swagger.io](https://petstore.swagger.io/) 网页直接访问 `https://raw.githubusercontent.com/xiaosongfu/freeswitch-helper/master/docs/api/swagger.json`

