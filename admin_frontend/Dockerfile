FROM node

MAINTAINER Alex Bondar <abondar1992@gmail.com>

WORKDIR /app

COPY . /app

RUN  npm install
RUN  npm run build

CMD ["npm", "start"]
