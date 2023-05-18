
FROM golang:1.19

# Copy the source code
WORKDIR /app
RUN mkdir -m 755 /app/results/
RUN mkdir -m 755 /app/orders/
COPY . .
# download depedency
RUN go mod tidy
# Build the binary files

# ENV API_EXTERNAL=http://margono.dvl.to/coba 
# ENV TELEGRAM_API_TOKEN=5839616364:AAEYhcMPL-AoUtcuNF7o6F5S76ye8uCoVow 
# ENV GROUPCHAT_ID=-1001704190576 
# ENV RESULTDIR=/app 
# ENV ORDERDIR=/app
# ENV HOST=0.0.0.0
# ENV PORT=4100
# # ENV AMQP_URL=amqp://guest:guest@localhost:5672/
# ENV APP_MODE=PROD

RUN go build -o binary

# ENTRYPOINT ["/binary"]
CMD ["./binary"]