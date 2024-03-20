FROM maven:3.8.4-jdk-11 as builder

COPY src /app/src
COPY pom.xml /app

WORKDIR /app

RUN mvn clean package

FROM openjdk:11-jre-slim

ENV APP_VERSION="1.0-SNAPSHOT"

LABEL org.opencontainers.image.title="VMware Secrets Manager Java SDK" \
      org.opencontainers.image.description="Java SDK for VMware Secrets Manager" \
      org.opencontainers.image.version=$APP_VERSION \
      org.opencontainers.image.authors="VSecM Maintainers <maintainers@vsecm.com>" \
      org.opencontainers.image.vendor="VMware, Inc." \
      org.opencontainers.image.licenses="BSD-2-Clause" \
      org.opencontainers.image.url="https://vsecm.com/" \
      org.opencontainers.image.documentation="https://vsecm.com/docs" \
      org.opencontainers.image.source="https://github.com/vmware-tanzu/secrets-manager" \
      org.opencontainers.image.revision="v$APP_VERSION"

WORKDIR /app

COPY --from=builder /app/target/sdk-java-$APP_VERSION.jar /app/app.jar

ENTRYPOINT ["java", "-jar", "app.jar"]
