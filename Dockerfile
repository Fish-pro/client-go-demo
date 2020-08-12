FROM golang AS build-env
ENV GO111MODULE on
ENV GOFLAGS -mod=vendor
ENV GOPROXY https://goproxy.cn
ARG appRoot=/go/src/github.com/daocloud/dsp-appserver
ADD . $appRoot
RUN cd $appRoot && go build -o /go/appserver ./cmd/appserver

FROM daocloud.io/atsctoo/rhel8-go-toolset:1.11.6-22
COPY --from=build-env /go/appserver /appserver

ENV TZ Asia/Shanghai
EXPOSE 8989
CMD ["/appserver"]
