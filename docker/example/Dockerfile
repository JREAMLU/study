FROM hub.longzhu.com/library/alpine:latest

ENV CONSUL_HTTP_ADDR=10.200.150.3:8500
ENV MICRO_BROKER="kafka"
ENV MICRO_BROKER_ADDRESS="10.200.150.3:9092"
COPY log.xml log-test.xml log-prod.xml /
COPY ./relationship-web /

CMD ["/relationship-web","--server=grpc","--client=grpc","--transport=grpc", "--server_address=0.0.0.0:40016"]
