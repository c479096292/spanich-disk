FROM alpine:3.2

COPY bin/upload /
#COPY ./static/ /static/

RUN chmod 777 /upload

ENV PARAMS=""

ENTRYPOINT ["sh","-c","/upload $PARAMS"]