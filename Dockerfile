FROM golang:1.15.13

ENV REPO_URL=github.com/abhilashdk2016/bookstore_items_api
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o items-api .
EXPOSE 8082

CMD ["./items-api"]
