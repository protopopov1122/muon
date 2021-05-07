FROM alpine:latest

RUN apk add python3 nodejs yarn nginx py3-pip
RUN mkdir -p /srv/api/webapi /srv/app /srv/public /srv/files/static
ADD start.sh /srv
# Setup API
ADD api/webapi /srv/api/webapi
ADD api/requirements.txt /srv/api
ADD config/webapi.json /srv/api
WORKDIR /srv/api
RUN pip3 install --ignore-installed six -r requirements.txt
# Setup application
ADD app /srv/app/build
ADD config/app.json /srv/app/build/src/config.json
WORKDIR /srv/app/build
RUN rm -rf node_modules dist
RUN yarn
RUN yarn build
RUN mv dist/* /srv/app
WORKDIR /srv
RUN rm -rf /srv/app/build
# Setup nginx
ADD config/nginx.conf /etc/nginx/nginx.conf
# Setup user
RUN addgroup -S www && adduser -S www -G www
RUN chown -R root:www /srv
RUN chmod -R 750 /srv
RUN touch /run/nginx.pid
RUN mkdir /var/cache/nginx
RUN mkdir /run/sockets
RUN chown www:www /run/nginx.pid
RUN chown www:www /run/sockets
RUN chown -R www:www /var/cache/nginx
RUN chown -R www:www /var/log/nginx
RUN chown -R www:www /var/lib/nginx
USER www
ENTRYPOINT ["./start.sh"]
