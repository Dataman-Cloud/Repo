FROM ubuntu:14.04
MAINTAINER upccup

RUN apt-get update && \
    apt-get install -yq reprepro createrepo nginx && \
    mkdir -p /var/repositories/ && \
	  mv /etc/nginx/sites-available/default /etc/nginx/sites-available/default.bak

ADD ./ /var/repositories/

CMD ["/var/repositories/build.sh"]
