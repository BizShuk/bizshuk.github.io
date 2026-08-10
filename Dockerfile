# syntax=docker/dockerfile:1

FROM nginx:1.29-alpine
COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

COPY . /usr/share/nginx/html
RUN rm -rf \
    /usr/share/nginx/html/.git \
    /usr/share/nginx/html/Dockerfile \
    /usr/share/nginx/html/docker \
    && addgroup -S -g 10001 app \
    && adduser -S -D -u 10001 -G app -h /home/app app \
    && touch /var/run/nginx.pid \
    && chown -R app:app /etc/nginx/conf.d /usr/share/nginx/html /var/cache/nginx /var/run/nginx.pid

ENV HOME=/home/app

EXPOSE 8306

USER app

CMD ["nginx", "-g", "daemon off;"]
