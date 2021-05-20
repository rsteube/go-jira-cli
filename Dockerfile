FROM golang:alpine

RUN apk add \
        --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/testing \
        bat \
        elvish

ADD . /go-jira-cli
RUN cd /go-jira-cli/cmd/gj \
 && go install

RUN mkdir -p /root/.config/gj /root/.elvish \
 && echo 'issues.apache.org/jira: {}' >> /root/.config/gj/hosts.yaml \
 && echo 'host: issues.apache.org/jira' >> /root/.config/gj/default.yaml \
 && echo 'eval (gj _carapace|slurp)' >> /root/.elvish/rc.elv

CMD [ "elvish" ]
