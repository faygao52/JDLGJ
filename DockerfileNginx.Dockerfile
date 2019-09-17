FROM alpine

MAINTAINER Jingyi Gao <faygao52@gmail.com>

COPY /nginx/certs /etc/nginx/certs/
COPY /nginx/conf.d /etc/nginx/conf.d/
COPY /LGJ-Dashboard /var/www/localhost/htdocs
ARG NODE_ENV=development
ENV NODE_ENV=$NODE_ENV

ARG REACT_APP_API_URL=http://localhost:8080
ENV REACT_APP_API_URL=$REACT_APP_API_URL
RUN apk add nginx && \
    mkdir /run/nginx && \
    apk add nodejs && \
    apk add npm && \
    cd /var/www/localhost/htdocs && \
    npm install && \
    npm run build && \
    apk del nodejs && \
    apk del npm && \
    mv /var/www/localhost/htdocs/build /var/www/localhost && \
    cd /var/www/localhost/htdocs && \
    rm -rf * && \
    mv /var/www/localhost/build /var/www/localhost/htdocs;
CMD ["/bin/sh", "-c", "exec nginx -g 'daemon off;';"]
WORKDIR /var/www/localhost/htdocs