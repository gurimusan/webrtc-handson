FROM node:alpine

RUN npm install ws

USER node
WORKDIR /home/node
COPY signaling.js /home/node/signaling.js

EXPOSE 3001
CMD ["node", "signaling.js"]
