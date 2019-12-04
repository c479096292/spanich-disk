FROM alpine:3.2

COPY bin/dbproxy /
#COPY ./static/ /static/

RUN chmod 777 /dbproxy

ENV PARAMS=""

ENTRYPOINT ["sh","-c","/dbproxy $PARAMS"]