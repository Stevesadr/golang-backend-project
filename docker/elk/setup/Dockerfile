ARG ELASTIC_VERSION

FROM elasticsearch:${ELASTIC_VERSION}

USER root

COPY entrypoint.sh /entrypoint.sh
COPY helpers.sh /helpers.sh

RUN set -eux; \
	chmod +x /entrypoint.sh; \
	chmod +x /helpers.sh; \
	mkdir /state; \
	chmod 0775 /state; \
	chown elasticsearch:root /state

USER elasticsearch:root

ENTRYPOINT ["/entrypoint.sh"]