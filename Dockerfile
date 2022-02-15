FROM golang

RUN mkdir /interview-accountapi
WORKDIR /interview-accountapi
COPY . .
CMD ["go", "test", "-v"]
