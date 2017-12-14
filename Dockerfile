# Наследуемся от CentOS 7
FROM ubuntu:16.04

# Выбираем рабочую папку
WORKDIR /root

RUN apt-get -qq update
RUN apt-get -qq -y install wget

# Устанавливаем wget и скачиваем Go
RUN apt-get install -y wget && \
    wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz

# Устанавливаем Go, создаем workspace и папку проекта
RUN tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz && \
    mkdir go && mkdir go/src && mkdir go/bin && mkdir go/pkg && \
    mkdir go/src/dumb

# Задаем переменные окружения для работы Go
ENV PATH=${PATH}:/usr/local/go/bin GOROOT=/usr/local/go GOPATH=/root/go

RUN apt-get update && apt-get install -y git
RUN git --version

RUN go get github.com/gin-gonic/gin
RUN go get "github.com/jinzhu/gorm"
RUN go get "github.com/jinzhu/gorm/dialects/mysql"

RUN wget http://dev.mysql.com/get/mysql-apt-config_0.6.0-1_all.deb

RUN apt-get install -y debconf-utils \
    && echo mysql-server mysql-server/root_password password root | debconf-set-selections \
    && echo mysql-server mysql-server/root_password_again password root | debconf-set-selections \ 
    && apt-get -y install mysql-server

RUN apt-get update && apt-get install -y zip unzip

ADD /golang go/src/golang

# Компилируем и устанавливаем наш сервер
RUN export GIN_MODE=release
RUN go build golang && go install golang

# Открываем 80-й порт наружу
EXPOSE 80

RUN chmod +x go/src/golang/start.sh

# Запускаем наш сервер
CMD go/src/golang/start.sh
