# build stage

FROM node:lts-alphine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build


# production stage
FROM ngnix:stable-alphine as production-stage
COPY --from=build-stage /app/dist /usr/share/ngnix/html
EXPOSE 80
CMD ["ngnix","-g", "daemon off;"]