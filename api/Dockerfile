# build stage
FROM golang:alpine AS build-env
WORKDIR '/go/src/DockerGoNginx/api'
ADD . .
RUN apk add git && go get -u github.com/kardianos/govendor && govendor sync
RUN go build -o APPSERVICE

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/DockerGoNginx/api/APPSERVICE /app/
COPY --from=build-env /go/src/DockerGoNginx/api/.env /app/
RUN ["mkdir","logs"]
RUN ["chmod", "+x", "APPSERVICE"]
RUN ["ls", "-al"]
EXPOSE 8000
# ENTRYPOINT ./APPSERVICE
CMD ./APPSERVICE
