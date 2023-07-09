FROM golang:alpine3.17

WORKDIR /data

COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk --no-cache --no-progress add make \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && make

CMD ["/jnote-server", "server"]
