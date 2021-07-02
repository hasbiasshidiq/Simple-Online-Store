# set base image (host OS)
FROM python:3.8

# set the working directory in the container
WORKDIR /code

# install dependencies
RUN pip3 install requests

# copy the content of the local src directory to the working directory
COPY . .