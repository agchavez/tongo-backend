FROM golang:1.13
RUN mkdir /src

ADD /src /src


WORKDIR /src
RUN go get -d -v 
RUN go build -o main .

CMD ["/src/main"]