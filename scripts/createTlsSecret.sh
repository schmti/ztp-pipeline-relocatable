#!/bin/bash

set -e

mkdir -p ./certs
if [ x${TLS_CERT_FILE} = x ] ; then
  export TLS_CERT_FILE=./certs/tls.crt
  export TLS_KEY_FILE=./certs/tls.key

  echo Autogenerating TLS certificates, set TLS_CERT_FILE and TLS_KEY_FILE environment variables if you want otherwise 
  openssl req -subj '/C=US' -new -newkey rsa:2048 -sha256 -days 365 -nodes -x509 -keyout certs/tls.key -out certs/tls.crt
fi

oc delete secret kubeframe-ui-certs -n kubeframe-ui
oc create secret tls kubeframe-ui-certs -n kubeframe-ui \
  --cert=${TLS_CERT_FILE} \
  --key=${TLS_KEY_FILE}

