ARG NODE_IMAGE=node:8.11

FROM ${NODE_IMAGE} AS node-build

WORKDIR /src/Web/WebSPA

COPY src/Web/WebSPA/package.json ./
COPY src/Web/WebSPA/package-lock.json ./

RUN npm install --only=production

COPY src/Web/WebSPA .

RUN npm run build:client

CMD ["node", "./server/bootstrap.js"]