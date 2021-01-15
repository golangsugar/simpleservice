FROM golang:1.15-alpine AS builder

# Install git and certificates for deal with versioning system
RUN apk --update-cache upgrade && apk add --no-cache git mercurial ca-certificates

ENV APPNAME simpleservice
ENV SRCPATH /src/${APPNAME}
ENV BINPATH /${APPNAME}

RUN mkdir -p --verbose ${SRCPATH} \
    && mkdir -p --verbose ${BINPATH} \
    && git clone --single-branch --branch master --depth 1 https://github.com/miguelpragier/${APPNAME}.git ${SRCPATH}

WORKDIR ${SRCPATH}

RUN go get -u && go mod tidy && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-extldflags '-static'" -o ${BINPATH}/${APPNAME} .

# ######################################################################################################################
FROM scratch

ENV APPNAME simpleservice

COPY --from=builder /${APPNAME} /${APPNAME}/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR ${APPNAME}

ENTRYPOINT ["./simpleservice"]

# ######################################################################################################################
# docker image build --no-cache --rm -t miguelpragier/simpleservice:latest .
# docker push miguelpragier/simpleservice:latest