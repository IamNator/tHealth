# build stage

FROM node:lts-alpine as build-stage
WORKDIR /thealth
COPY /thealth/package*.json ./
RUN npm install
COPY ./thealth/. .
RUN npm run build


# production stage
FROM nginx:stable-alpine as production-stage

COPY ./.nginx/nginx.conf /etc/nginx/nginx.conf

## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

RUN mkdir /usr/share/nginx/html/dist

COPY --from=build-stage /thealth/dist /usr/share/nginx/html/dist
COPY thealth/index.html /usr/share/nginx/html


EXPOSE 80

CMD ["nginx","-g", "daemon off;"]