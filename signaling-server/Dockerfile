FROM node:alpine

RUN npm install ws

USER node
WORKDIR /home/node/server
COPY server /home/node/server

EXPOSE 3001
CMD ["node", "signaling.js"]
