FROM alpine:3.2

COPY bin/apigwv2 /
#COPY ./static/ /static/

RUN chmod 777 /apigwv2

ENV PARAMS=""

ENTRYPOINT ["sh","-c","/apigwv2 $PARAMS"]