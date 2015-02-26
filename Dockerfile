FROM scratch
MAINTAINER mail.twerner@gmail.com

EXPOSE 4000

ADD bin/web /main
CMD ["/main"]

