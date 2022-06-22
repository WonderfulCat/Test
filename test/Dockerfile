FROM golang:1.18

ENV GO111MODULE=on

COPY . /opt/test

WORKDIR /opt/test/src

RUN go build -o app .

RUN cp /opt/test/config/testItem.data /opt/test/src

EXPOSE 8080

ENTRYPOINT [ "/opt/test/src/app" ]






