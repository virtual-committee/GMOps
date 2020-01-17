FROM centos:latest

ADD ./ /opt/GMOps
WORKDIR /opt/GMOps

RUN scripts/install.sh

EXPOSE 8080
EXPOSE 9022

CMD ["scripts/start.sh"]
