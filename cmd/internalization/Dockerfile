FROM jrei/systemd-centos:7

COPY Internalization.viper.yaml /etc/Internalization/Internalization.viper.yaml
COPY internalization /usr/local/bin/internalization

CMD ["internalization", "run"]
