FROM nvidia/cuda:10.1-base

WORKDIR /

COPY omaticaya /usr/local/bin

CMD ["omaticaya"]