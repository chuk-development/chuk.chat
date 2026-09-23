FROM hugomods/hugo:latest AS builder
WORKDIR /src
COPY . .
RUN hugo --minify

FROM nginx:alpine
# image_filter resizes the app screenshots proxied from GitHub (nginx.conf).
RUN sed -i '1i load_module modules/ngx_http_image_filter_module.so;' /etc/nginx/nginx.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /src/public /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
