FROM node:lts-alpine3.18 as builder

WORKDIR /build

COPY . .

RUN npm install

RUN npm run build

FROM nginx:alpine3.18

COPY --from=builder /build/dist /usr/share/nginx/html

ADD server/nginx.conf /etc/nginx/nginx.conf

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]