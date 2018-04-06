FROM library/golang
FROM quay.io/goswagger/swagger

MAINTAINER Alex Bondar <abondar1992@gmail.com>

WORKDIR /app

ENV SRC_DIR=$GOPATH/src/github.com/abondar24/SocialTournamentService
ADD . $SRC_DIR

RUN cd $SRC_DIR/main; go get
RUN cd $SRC_DIR/main; go build -o social; cp social /app/
RUN cp $SRC_DIR/api/swagger.json /app/
ENTRYPOINT ["./social"]