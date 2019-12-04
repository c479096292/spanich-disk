FROM alpine:3.2

COPY bin/apigw /
#COPY ./static/ /static/

RUN chmod 777 /apigw

ENV PARAMS=""

ENTRYPOINT ["sh","-c","/apigw $PARAMS"]