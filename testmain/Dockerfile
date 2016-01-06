FROM golang

ENV GOPATH /gopath
RUN go get github.com/shaalx/daomgo/testmain

EXPOSE 80

CMD ["/gopath/bin/testmain"]