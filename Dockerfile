FROM golang:1.14 as build
COPY . /home/workspace/SKB
WORKDIR /home/workspace/SKB
RUN go build -o server cmd/*.go

FROM ubuntu:18.04

RUN sed -i 's#http://archive.ubuntu.com/#http://mirrors.tuna.tsinghua.edu.cn/#' /etc/apt/sources.list \
&& sed -i 's#http://security.ubuntu.com/#http://mirrors.tuna.tsinghua.edu.cn/#' /etc/apt/sources.list &&  apt update \
&& apt install -y procps tzdata vim && rm -f /etc/localtime && echo "Asia/Shanghai" > /etc/timezone \
&& dpkg-reconfigure -f noninteractive tzdata &&  sed -i "s?# alias ll?alias ll?" ~/.bashrc

COPY --from=build /home/workspace/SKB/server /server

CMD ["/server"]