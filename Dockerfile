FROM golang:1.16 as builder
COPY . /src
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct CGO_ENABLED=0
RUN cd /src && \
 go build -o ss5 .

FROM scratch
COPY --from=builder /src/ss5 /src/ss5
WORKDIR /src
ENTRYPOINT ["/src/ss5"]

EXPOSE 1080


