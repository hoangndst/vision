FROM golang:1.23 AS build
COPY / /src
WORKDIR /src
RUN make build-local-linux

FROM ubuntu:22.04
# Install Dependencies
RUN apt-get update -y && apt-get install python3 python3-pip -y

# VISION
ENV PATH="/root/go/bin:${PATH}"
ENV VISION_HOME="$HOME/.vision"
ENV LANG=en_US.utf8

COPY --from=build /src/_build/bundles/vision-linux/bin/vision /usr/local/bin/vision
CMD ["/usr/local/bin/vision", "server"]
