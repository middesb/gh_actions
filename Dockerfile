# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

RUN mkdir /app
ADD . /app

WORKDIR /app
  

RUN go build -o middesb/gh_actions_app .

 

CMD [ "/app/middesb/gh_actions_app" ]

ENTRYPOINT [ "go", "run" ,"main.go" ]