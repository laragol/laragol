#FROM scratch
FROM golang:latest
##ADD ca-certificates.crt /etc/ssl/certs/
RUN mkdir -p /code
WORKDIR /code
ENV HOME /code
ADD . .
EXPOSE 8080
#ENTRYPOINT ["/main"]
CMD ["make", "dev"]
#CMD ["/bin/bash"]
