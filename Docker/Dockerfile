FROM golang:1.8.1-alpine

ENV GOBINARIES /go/bin
ENV BUILDPATH /code
ENV REFLEXURL=http://s3.amazonaws.com/wbm-raff/bin/reflex1.8a
ENV REFLEXSHA=19bdbbb68c869f85ee22a6b7fa9c73f8e5b46d0fe7a73df37e028555a6ba03e8

WORKDIR $GOBINARIES
RUN rm -rf /var/cache/apk/*
RUN wget -q "$REFLEXURL" -O reflex
RUN chmod +x /go/bin/reflex

ENV TOOLS /go/_tools

RUN mkdir -p $BUILDPATH
ENV PORT 5000
EXPOSE $PORT

RUN mkdir -p $TOOLS
ADD build.sh $TOOLS
ADD reflex.conf $TOOLS
RUN chown root $TOOLS/build.sh
RUN chmod +x $TOOLS/build.sh

WORKDIR $BUILDPATH
CMD ["reflex","-c","/go/_tools/reflex.conf"]
