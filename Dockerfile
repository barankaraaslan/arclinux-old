FROM golang

WORKDIR /app

COPY package-manager/go.mod ./
RUN go mod download

COPY **/*.go ./

RUN go build -o /packagemanager

FROM scratch

COPY --from=0 /packagemanager /packagemanager
CMD [ "/packagemanager" ]