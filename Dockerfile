FROM golang:1.12.5

MAINTAINER hehong <deng2296917779@163.com>


ENV GOPATH $home/go
ENV PATH $PATH:$GOPATH/bin
ENV APP_DIR $GOPATH/src/zeus

RUN go get github.com/astaxie/beego && go get github.com/beego/bee

WORKDIR $APP_DIR
RUN mkdir -p $APP_DIR
ADD . $APP_DIR

EXPOSE 8080
CMD go run main.go






