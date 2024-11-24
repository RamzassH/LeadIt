FROM node:20-alpine

RUN corepack enable

RUN corepack prepare yarn@4.5.2 --activate

WORKDIR /app

COPY .yarnrc.yml ./
COPY package.json yarn.lock ./

RUN yarn config set npmRegistryServer https://registry.yarnpkg.com

RUN yarn install --frozen-lockfile

COPY . .

RUN yarn build

EXPOSE 3000

CMD ["yarn", "start"]
