FROM alpine:3.2

COPY bin/download /
#COPY ./static/ /static/

RUN chmod 777 /download

ENV PARAMS=""

ENTRYPOINT ["sh","-c","/download $PARAMS"]