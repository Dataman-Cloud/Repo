#!/bin/bash

set -e

reprepro -b /var/repositories includedeb datamacloud /var/repositories/deb/packages/omega-agen*
createrepo -p -d -o /var/repositories/yum/centos/7/x86_64/ /var/repositories/yum/centos/7/x86_64/

service nginx restart


