FROM maven:3.6.3-openjdk-14-slim AS build

COPY settings.xml /usr/share/maven/conf/

COPY pom.xml pom.xml
COPY dis-api/pom.xml dis-api/pom.xml
COPY dis-model/pom.xml dis-model/pom.xml
COPY dis-base/pom.xml dis-base/pom.xml
COPY dis-database/pom.xml dis-database/pom.xml

RUN mvn dependency:go-offline package -B

COPY dis-api/src dis-api/src
COPY dis-model/src dis-model/src
COPY dis-base/src dis-base/src
COPY dis-database/src dis-database/src

RUN mvn install

FROM openjdk:14-ea-jdk-alpine
USER root

RUN mkdir service

COPY --from=build /dis-base/target/ /service/
COPY drop_data.json /service/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

ENV JAVA_TOOL_OPTIONS -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:5005

EXPOSE 5005

CMD /wait && java --enable-preview -jar /service/dis-base-1.0-SNAPSHOT.jar -Xdebug