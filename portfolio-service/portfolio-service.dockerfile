FROM alpine:latest

RUN mkdir /app

COPY portfolioApp /app

CMD ["/app/portfolioApp"]