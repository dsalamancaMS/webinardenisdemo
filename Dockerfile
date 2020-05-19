FROM golang:latest as stage1 
RUN mkdir app
ADD Application/ /app/
WORKDIR /app
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o web .

FROM alpine:latest
WORKDIR /root/
ADD Application/ /root/
COPY --from=stage1 /app/web .
CMD ["./web"]
EXPOSE 8080
