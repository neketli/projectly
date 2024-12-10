FROM node:18 AS build

WORKDIR /app

COPY . .

RUN npm ci
RUN npm run generate

FROM nginx:alpine

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
