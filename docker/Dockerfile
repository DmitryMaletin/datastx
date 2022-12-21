ARG ANTLR_VERSION="4.11.1"
FROM adoptopenjdk/openjdk11 AS builder

WORKDIR /usr/local/lib

# ARG ANTLR_VERSION="4.11.1"
ARG MAVEN_OPTS="-Xmx1G"

RUN apt-get -y update
RUN apt-get -y install curl
RUN curl -O https://www.antlr.org/download/antlr-4.11.1-complete.jar


FROM adoptopenjdk/openjdk11

ARG user=appuser
ARG group=appuser
ARG uid=1000
ARG gid=1000

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --no-create-home \
    --uid "${uid}" \
    "${user}"

# ARG ANTLR_VERSION="4.11.1"

COPY --from=builder /usr/local/lib/antlr-4.11.1-complete.jar /usr/local/lib/antlr-4.11.1-complete.jar



ENTRYPOINT ["java", "-jar", "/usr/local/lib/antlr-4.11.1-complete.jar"]

