# Build the Go Binary.
FROM golang:1.14 as build_metrics
ENV CGO_ENABLED 0
ARG VCS_REF
ARG PACKAGE_NAME
ARG PACKAGE_PREFIX

# Create a location in the container for the source code. Using the
# default GOPATH location.
RUN mkdir -p /service

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
#COPY go.* /service/
#WORKDIR /service
#RUN go mod download

# Copy the source code into the container.
WORKDIR /service
COPY go.* ./
COPY cmd cmd
COPY internal internal
COPY vendor vendor

# Build the service binary. We are doing this last since this will be different
# every time we run through this process.
WORKDIR /service/cmd/${PACKAGE_PREFIX}${PACKAGE_NAME}
RUN go build -mod=vendor -ldflags "-X main.build=${VCS_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.11
ARG BUILD_DATE
ARG VCS_REF
ARG PACKAGE_NAME
ARG PACKAGE_PREFIX
COPY --from=build_metrics /service/cmd/${PACKAGE_PREFIX}${PACKAGE_NAME}/${PACKAGE_NAME} /app/main
WORKDIR /app
CMD /app/main

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="${PACKAGE_NAME}" \
      org.opencontainers.image.revision="${VCS_REF}"
