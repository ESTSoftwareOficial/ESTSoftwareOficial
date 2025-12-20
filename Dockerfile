FROM ubuntu:latest
LABEL authors="ameth"

ENTRYPOINT ["top", "-b"]