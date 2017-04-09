FROM gett/golang_base

ENV APP_NAME {|PROJECTNAME|}

RUN mkdir -p /go/src/github.com/gtforge/$APP_NAME
WORKDIR /go/src/github.com/gtforge/$APP_NAME
RUN echo -e "machine github.com\n  login b08b064f7d82cfb3996562e74e125160e076f6be" >> ~/.netrc
ADD glide.yaml /go/src/github.com/gtforge/$APP_NAME
ADD glide.lock /go/src/github.com/gtforge/$APP_NAME
RUN glide install
ADD . /go/src/github.com/gtforge/$APP_NAME
RUN GOGC=off go build -v -i -ldflags "-X github.com/gtforge/$APP_NAME/vendor/github.com/gtforge/services_common_go/gett-alive.Buildstamp=`date -u +%Y/%m/%d_%H:%M:%S` -X github.com/gtforge/$APP_NAME/vendor/github.com/gtforge/services_common_go/gett-alive.Commit=`git rev-parse HEAD`"

CMD /go/src/github.com/gtforge/$APP_NAME/$APP_NAME
EXPOSE 80
