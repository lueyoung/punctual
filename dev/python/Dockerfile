FROM python:3
ADD requirements.txt /tmp/
RUN pip install -r /tmp/requirements.txt
RUN apt-get update
RUN apt-get install -y vim --allow-unauthenticated
WORKDIR /workspace 
