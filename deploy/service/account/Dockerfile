FROM alpine:3.2

COPY bin/account /
COPY ./static/ /static/

RUN chmod 777 /account

#ENTRYPOINT [ "/account-service" ]
ENV PARAMS=""
#
ENTRYPOINT ["sh","-c","/account $PARAMS"]