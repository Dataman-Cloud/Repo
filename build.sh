#!/bin/bash

set -e

reprepro -b /var/repositories includedeb datamancloud /var/repositories/debs/packages/omega-agen*
createrepo -p -d -o /var/repositories/yum/centos/7/x86_64/ /var/repositories/yum/centos/7/x86_64/

# sudo service nginx restart

/usr/sbin/nginx -g "daemon off;"
