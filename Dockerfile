FROM alpine

ADD server .

EXPOSE 55051

ENTRYPOINT [ "./server", "8080" ]