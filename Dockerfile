#Set the builder image
FROM golang:latest as GOLANG-BUILDER

#Copy in our source file
COPY /src/* /src/

WORKDIR /src

RUN go mod tidy

#Compile the Code
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /hwassign/main main.go

#Select a base image for final image
FROM alpine

#Copy the compiled file from the base image to the final image
COPY --from=GOLANG-BUILDER /hwassign/main /app/main

EXPOSE 5309

#Set the default entry point
CMD ["/app/main"]
