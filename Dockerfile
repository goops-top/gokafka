FROM golang as golang-build
WORKDIR /
COPY . / 
RUN make default
FROM alpine
COPY --from=golang-build build/gokafka.linux /
CMD ["/gokafka.linux"]
