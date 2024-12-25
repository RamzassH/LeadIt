FROM node:20-alpine

WORKDIR /app

COPY package.json ./

RUN npm install -g npm@latest && npm install --frozen-lockfile --verbose

COPY . .

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]
