# syntax=docker/dockerfile:1

FROM nginx:1.29-alpine AS image-build
WORKDIR /home/app/src

COPY . .
RUN mkdir -p /home/app/out/html \
    && cp -a . /home/app/out/html/ \
    && rm -rf \
        /home/app/out/html/.git \
        /home/app/out/html/Dockerfile \
        /home/app/out/html/docker \
    && find /home/app/out -type f -name '*.map' -delete \
    && cp docker/nginx.conf /home/app/out/nginx.conf \
    && printf '%s\n' '#!/bin/sh' 'exec nginx -c /home/app/out/nginx.conf -g "daemon off;"' > /home/app/out/entrypoint.sh \
    && chmod +x /home/app/out/entrypoint.sh

FROM nginx:1.29-alpine
RUN addgroup -S -g 10001 app \
    && adduser -S -D -u 10001 -G app -h /home/app app \
    && install -d -o app -g app \
        /home/app/.config/bizshuk.github.io/data \
        /home/app/.config/bizshuk.github.io/logs \
    && printf '{}\n' > /home/app/.config/bizshuk.github.io/settings.json \
    && chown app:app /home/app/.config/bizshuk.github.io/settings.json \
    && chown -R app:app /var/cache/nginx
WORKDIR /home/app

COPY --from=image-build --chown=app:app /home/app/out ./out
RUN ln -s /home/app/out/entrypoint.sh /home/app/app \
    && chown -h app:app /home/app/app

EXPOSE 8306

USER app

ENTRYPOINT ["/home/app/app"]
