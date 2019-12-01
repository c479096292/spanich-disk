FROM alpine:3.2

COPY bin/transfer /
COPY ./static/ /static/

RUN chmod 777 /transfer

ENV PARAMS=""

ENTRYPOINT ["sh","-c","/transfer $PARAMS"]