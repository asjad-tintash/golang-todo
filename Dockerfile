FROM golang:1.16
RUN mkdir /todo-app
ADD . /todo-app
WORKDIR /todo-app
RUN go mod download
RUN go build
#EXPOSE 8000
RUN echo "ASJAD"
CMD ["./golang-todo"]