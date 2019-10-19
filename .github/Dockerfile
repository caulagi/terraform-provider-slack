FROM golang:1.13-alpine

ARG TERRAFORM_VERSION=0.12.12

RUN apk --update --no-cache add wget unzip
RUN wget "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip"
RUN unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip
RUN mv terraform /usr/bin