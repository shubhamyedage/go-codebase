FROM debian:latest

RUN mkdir /app
WORKDIR /app
ADD consignment-service /app/consignment-service
RUN ls
RUN pwd

CMD ["./consignment-service"]
