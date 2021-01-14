FROM gcr.io/tetratelabs/tetrate-base:v0.1

ADD bin/ext-authz /usr/local/bin/ext-authz

WORKDIR /usr/local/bin
COPY certs certs/

ENTRYPOINT [ "/usr/local/bin/ext-authz" ]