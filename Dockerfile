FROM node:14

WORKDIR /usr/src/app

COPY package*.json ./
COPY yarn.lock ./
COPY ./dist ./

RUN yarn install

EXPOSE 3000

ENTRYPOINT ["node", "index.js"]