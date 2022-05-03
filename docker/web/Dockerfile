# ref: https://v2.vuejs.org/v2/cookbook/dockerize-vuejs-app.html#Real-World-Example
# ref: https://docs.docker.com/develop/develop-images/multistage-build/
FROM node:14.19.1 as build-stage
WORKDIR /app
COPY package*.json ./
COPY . .
# RUN npm install -g yarn
RUN yarn install
RUN yarn build

FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY ./docker/web/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
