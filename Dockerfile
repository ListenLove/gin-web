# 第一阶段：构建环境
# 使用官方的Go镜像作为构建环境
FROM golang:1.20 AS builder

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /build

# 将Go模块的依赖文件拷贝到容器中
COPY go.mod .
COPY go.sum .

# 下载Go模块依赖
RUN go mod download

# 将源代码拷贝到容器中
COPY . .

# 编译应用
RUN go build -o main .

# 第二阶段：构建运行环境
# 使用scratch作为基础镜像，它是一个空白的镜像
FROM scratch

# 从构建阶段拷贝编译好的应用、配置文件等
COPY --from=builder /build/main /
COPY --from=builder /build/config.yaml /

# 暴露端口
EXPOSE 8081

# 运行应用
# CMD ["/main"]
