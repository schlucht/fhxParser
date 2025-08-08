
FROM node:20

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY dist/ ./dist
COPY views/ ./views
COPY public/ ./public

EXPOSE 3000

CMD ["node", "dist/server.js"]
