FROM golang
ADD . ./
RUN go build -o main
ENV PORT 3000
EXPOSE 3000
CMD ["./main"]

