# build stage

FROM node:lts-alpine as build-stage
WORKDIR /thealth
COPY /thealth/package*.json ./
RUN npm install
COPY ./thealth/. .
RUN npm run build


# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /thealth/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx","-g", "daemon off;"]