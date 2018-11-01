FROM library/golang
FROM quay.io/goswagger/swagger

MAINTAINER Alex Bondar <abondar1992@gmail.com>

WORKDIR /app

ENV SRC_DIR=$GOPATH/src/github.com/abondar24/SocialTournamentService/api
ADD . $SRC_DIR

WORKDIR $SRC_DIR
RUN  swagger generate spec -o ./swagger.json

EXPOSE 8024
CMD ["serve","-p","8024","swagger.json", "--no-open"]
