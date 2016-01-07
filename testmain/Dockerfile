FROM golang

ENV GOPATH /gopath
ADD . /app
WORKDIR /app
RUN go get github.com/shaalx/daomgo/testmain && go build -o tmain /app/testmain/main.go

EXPOSE 80

CMD ["/app/testmain/tmain"]
