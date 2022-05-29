FROM golang:1.18 as build_courier

COPY . /courier

WORKDIR /courier/services/courier
RUN CGO_ENABLED=0 go build -o courier


FROM alpine
COPY --from=build_courier /courier/services/courier/courier /courier/courier
WORKDIR /courier
CMD ["./courier"]