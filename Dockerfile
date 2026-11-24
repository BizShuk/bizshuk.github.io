# syntax=docker/dockerfile:1

ARG NGINX_VERSION=1.29-alpine

FROM nginx:${NGINX_VERSION} AS image-build
WORKDIR /home/app/src

COPY . .
RUN mkdir -p /home/app/out/html \
    && cp -a . /home/app/out/html/ \
    && rm -rf \
        /home/app/out/html/.git \
        /home/app/out/html/.gitignore \
        /home/app/out/html/.dockerignore \
        /home/app/out/html/.claude \
        /home/app/out/html/.agents \
        /home/app/out/html/.githooks \
        /home/app/out/html/Dockerfile \
        /home/app/out/html/deployment.yml \
        /home/app/out/html/package.json \
        /home/app/out/html/pnpm-lock.yaml \
        /home/app/out/html/README.md \
        /home/app/out/html/AGENTS.md \
        /home/app/out/html/README.todo \
        /home/app/out/html/config \
        /home/app/out/html/scripts \
        /home/app/out/html/skills \
        /home/app/out/html/life \
        /home/app/out/html/tmp \
        /home/app/out/html/pkg/resume/jd \
        /home/app/out/html/pkg/resume/Resume.md \
    && find /home/app/out -type f -name '*.map' -delete \
    && cp config/nginx.conf /home/app/out/default.conf \
    && printf '%s\n' '#!/bin/sh' 'exec nginx -g "pid /home/app/.config/bizshuk.github.io/data/nginx.pid; daemon off;"' > /home/app/out/entrypoint.sh \
    && chmod +x /home/app/out/entrypoint.sh

FROM nginx:${NGINX_VERSION}
ARG APP_UID=10001
ARG APP_GID=10001
RUN u=$(awk -F: -v id="${APP_UID}" '$3 == id {print $1}' /etc/passwd) && { [ -z "$u" ] || deluser "$u"; } \
    && g=$(awk -F: -v id="${APP_GID}" '$3 == id {print $1}' /etc/group) && { [ -z "$g" ] || delgroup "$g"; } \
    && addgroup -g ${APP_GID} app \
    && adduser -u ${APP_UID} -G app -h /home/app -D app \
    && install -d -o app -g app \
        /home/app/.config/bizshuk.github.io/data \
        /home/app/.config/bizshuk.github.io/logs \
    && printf '{}\n' > /home/app/.config/bizshuk.github.io/settings.json \
    && chown app:app /home/app/.config/bizshuk.github.io/settings.json \
    && chown -R app:app /var/cache/nginx /var/log/nginx
WORKDIR /home/app

COPY --from=image-build --chown=app:app /home/app/out ./out
RUN cp /home/app/out/default.conf /etc/nginx/conf.d/default.conf \
    && ln -s /home/app/out/entrypoint.sh /home/app/app \
    && chown -h app:app /home/app/app

EXPOSE 8329

USER app

ENTRYPOINT ["/home/app/app"]
