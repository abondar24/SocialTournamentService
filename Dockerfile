FROM iron/go:dev

WORKDIR /app

ENV SRC_DIR=$GOPATH/src/github.com/abondar24/SocialTournamentService
ADD . $SRC_DIR

RUN cd $SRC_DIR/main; go get
RUN cd $SRC_DIR/main; go build -o social; cp social /app/
ENTRYPOINT ["./social"]