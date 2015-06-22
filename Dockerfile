FROM scratch
MAINTAINER mail.twerner@gmail.com

EXPOSE 4000

ADD web /main
CMD ["/main"]

