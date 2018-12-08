FROM golang:latest 
EXPOSE 5000
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
RUN cd /app ; pwd;ls ;go install payments
CMD ["/app/bin/payments"]
