FROM scratch

ADD parse-changelog /usr/bin/parse-changelog
ADD validate-changelog /usr/bin/validate-changelog
ADD lint-changelog /usr/bin/lint-changelog

ENTRYPOINT ["/usr/bin/validate-changelog"]