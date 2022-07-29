# syntax=docker/dockerfile:1
FROM node:16-alpine
RUN mkdir -p /home/project-dev/beta-flag-platform/flag-bearer
WORKDIR /home/project-dev/beta-flag-platform/flag-bearer
COPY package.json .
RUN npm install
COPY . .
CMD npm start
EXPOSE 3001