FROM nginx

COPY /nginx/certs /etc/nginx/certs/
COPY /nginx/conf.d /etc/nginx/conf.d/