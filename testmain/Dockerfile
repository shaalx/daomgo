FROM golang

ENV GOPATH /gopath
RUN go get github.com/shaalx/daomgo/testmain && go build -o tmain

EXPOSE 80

CMD ["tmain"]
