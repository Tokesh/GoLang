FROM golang:latest as build-stage
WORKDIR /project
COPY go.* .
COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o server ./app/server.go

CMD ["./server"]


FROM postgres

ENV POSTGRES_PASSWORD=tokesh
ENV POSTGRES_DB=postgres
COPY DML.sql /docker-entrypoint-initdb.d/