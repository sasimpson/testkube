FROM hypriot/rpi-golang
RUN mkdir /go
COPY ./service-linux-arm /go/
RUN chmod +x /go/service-linux-arm
CMD ["/go/service-linux-i386"]
