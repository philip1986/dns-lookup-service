# pull official base image
FROM node:14.5.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# add app
COPY . ./

# install app dependencies
RUN npm install --silent
RUN npm install react-scripts@4.0.3 -g --silent

# start app
CMD ["npm", "start"]