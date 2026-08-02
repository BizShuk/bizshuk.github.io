# syntax=docker/dockerfile:1

FROM nginx:1.29-alpine
COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

COPY . /usr/share/nginx/html
RUN rm -rf \
    /usr/share/nginx/html/.git \
    /usr/share/nginx/html/Dockerfile \
    /usr/share/nginx/html/docker

EXPOSE 8306

CMD ["nginx", "-g", "daemon off;"]
