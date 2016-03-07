FROM ubuntu:14.04
MAINTAINER upccup

RUN apt-get update
RUN apt-get install -yq reprepro createrepo nginx
RUN mkdir -p /var/repositories/
RUN mv /etc/nginx/sites-available/default /etc/nginx/sites-available/default.bak

ADD /etc/nginx/sites-available/ ./nginx/sites-available/default
ADD /var/repositories/ ./deb/conf/
ADD /var/repositories/ ./deb/packages/
ADD /var/repositories/ ./yum


CMD ["/build.sh"]
