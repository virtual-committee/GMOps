FROM node:alpine

ADD . /opt/GMOps
RUN npm install
