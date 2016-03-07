FROM ubuntu:14.04
MAINTAINER upccup

RUN apt-get update
RUN apt-get install -yq reprepro createrepo nginx
RUN mkdir -p /var/repositories/
RUN mv /etc/nginx/sites-available/default /etc/nginx/sites-available/default.bak

ADD ./nginx/sites-available/default /etc/nginx/sites-available/
ADD ./deb/conf/ /var/repositories/
ADD ./deb/packages/ /var/repositories/
ADD ./yum /var/repositories/
ADD ./build.sh /build.sh

CMD ["/build.sh"]
