FROM alpine:3.16

ARG IS_CHINA
ARG DEV_ENVIRONMENT
ENV REPOSITORY_DIR=/data/repository
ENV BOOKS_DIR=/data/books
ENV CONFIG_DIR=/data/config
COPY booker /opt/

RUN if [ "$IS_CHINA" == "true" ]; then sed -i 's@dl-cdn.alpinelinux.org@mirrors.aliyun.com@g' /etc/apk/repositories; fi
RUN apk update \
    && if [ "$IS_CHINA" == "true" ]; then apk add --no-cache --repository http://mirrors.aliyun.com/alpine/v3.12/main/ nodejs=12.22.12-r0; else apk add --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/v3.12/main/ nodejs=12.22.12-r0; fi \
    && apk add npm \
    && apk add go \
    && apk add git \
    && echo "export GO111MODULE=on" >> ~/.profile \
    && echo "export GOPROXY=https://goproxy.cn" >> ~/.profile \
    && source ~/.profile \
    && chmod 755 /opt/booker

# install gitbook
RUN npm install gitbook-cli -g 
RUN sed -i '/fs.stat / s/^\(.*\)$/\/\/\1/g' /usr/local/lib/node_modules/gitbook-cli/node_modules/npm/node_modules/graceful-fs/polyfills.js \
    && sed -i '/fs.fstat / s/^\(.*\)$/\/\/\1/g' /usr/local/lib/node_modules/gitbook-cli/node_modules/npm/node_modules/graceful-fs/polyfills.js \
    && sed -i '/fs.lstat / s/^\(.*\)$/\/\/\1/g' /usr/local/lib/node_modules/gitbook-cli/node_modules/npm/node_modules/graceful-fs/polyfills.js
RUN gitbook -V
RUN find / -name copyPluginAssets.js | xargs sed -i 's@confirm: true@confirm: false@g'

EXPOSE 5454

CMD ["sh", "-c", "/opt/booker -r $REPOSITORY_DIR -b $BOOKS_DIR -c $CONFIG_DIR"]