FROM hugomods/hugo:latest AS builder
WORKDIR /src
COPY . .
RUN hugo --minify

FROM golang:1.23-alpine AS shotconv
WORKDIR /src
COPY tools/shotconv/ .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /shotconv .

FROM nginx:alpine
# shotconv turns the phone screenshots (PNG) into small JPEGs (nginx.conf).
COPY --from=shotconv /shotconv /usr/local/bin/shotconv
COPY tools/shotconv/start.sh /docker-entrypoint.d/40-shotconv.sh
# image_filter resizes the app screenshots proxied from GitHub (nginx.conf).
RUN sed -i '1i load_module modules/ngx_http_image_filter_module.so;' /etc/nginx/nginx.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /src/public /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
