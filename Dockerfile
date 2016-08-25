FROM gett/golang_base

ENV APP_NAME {{PROJECTNAME}}

RUN mkdir -p /go/src/github.com/gtforge/$APP_NAME
WORKDIR /go/src/github.com/gtforge/$APP_NAME
ADD . /go/src/github.com/gtforge/$APP_NAME

RUN echo -e "machine github.com\n  login 54de8434f752617a7b939636a48999a1ceb11f78" >> ~/.netrc && glide install && \
    GOGC=off go build -v -ldflags "-X github.com/gtforge/$APP_NAME/vendor/github.com/gtforge/services_common_go/gett-alive.Buildstamp=`date -u +%Y/%m/%d_%H:%M:%S` -X github.com/gtforge/$APP_NAME/vendor/github.com/gtforge/services_common_go/gett-alive.Commit=`git rev-parse HEAD`"

CMD /go/src/github.com/gtforge/$APP_NAME/$APP_NAME
EXPOSE 80

