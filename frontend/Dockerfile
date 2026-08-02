FROM node:20 AS build

ARG NUXT_PUBLIC_API_HOST
ARG NUXT_PUBLIC_S3_HOST
ARG NUXT_PUBLIC_SITE_URL
ARG NUXT_APP_BASE_URL=/my/

WORKDIR /app

COPY . .

RUN npm ci
ENV NUXT_PUBLIC_API_HOST=$NUXT_PUBLIC_API_HOST \
    NUXT_PUBLIC_S3_HOST=$NUXT_PUBLIC_S3_HOST \
    NUXT_PUBLIC_SITE_URL=$NUXT_PUBLIC_SITE_URL \
    NUXT_APP_BASE_URL=$NUXT_APP_BASE_URL
RUN npm run generate

FROM nginx:alpine

COPY frontend.nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/.output/public /usr/share/nginx/html/my

CMD ["nginx", "-g", "daemon off;"]
