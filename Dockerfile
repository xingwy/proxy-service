FROM xxxx
LABEL maintainer="xxx"
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 4040
EXPOSE 10086
ENTRYPOINT ["entrypoint.sh"]
