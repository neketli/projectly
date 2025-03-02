FROM node:20 AS build

WORKDIR /app

COPY . .

RUN npm ci
RUN NUXT_APP_BASE_URL=/my/ npm run generate

FROM nginx:alpine

COPY frontend.nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/.output/public /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
