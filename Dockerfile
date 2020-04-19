FROM nvidia/cuda:10.1-base

WORKDIR /

COPY Omaticaya /usr/local/bin

CMD ["Omaticaya"]