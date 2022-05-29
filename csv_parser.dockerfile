FROM golang:1.18 as build_csv_parser

COPY . /csv_parser

WORKDIR /csv_parser/services/csv_parser
RUN CGO_ENABLED=0 go build -o csv_parser


FROM alpine
COPY --from=build_csv_parser /csv_parser/services/csv_parser/csv_parser /csv_parser/csv_parser
WORKDIR /csv_parser
CMD ["./csv_parser"]