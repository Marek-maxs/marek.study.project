FROM 766917153811.dkr.ecr.ap-east-1.amazonaws.com/golang:1.20.2

ENV OS linux
ENV ARCH amd64

WORKDIR /app

COPY main /usr/local/bin/energy

ENTRYPOINT [ "energy" ]