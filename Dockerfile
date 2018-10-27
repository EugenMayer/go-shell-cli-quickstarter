################################## backend #################################################
FROM eugenmayer/golang-builder AS go-build-env
# /go/src/ is mandatory due to the package format of go (lookup path )
# github.com due to our package name
# eugenmayer namespace
# go-antibash-boilerplate package name
# this needs to be changed to your likings / needs
ENV BUILDFOLDER=/go/src/github.com/eugenmayer/go-antibash-boilerplate
ENV DISTFOLDER=${BUILDFOLDER}/dist
RUN mkdir -p ${BUILDFOLDER}
WORKDIR ${BUILDFOLDER}
ADD . ${BUILDFOLDER}

# this simulates concourses input structure
ENV CI_BASE=/ci
RUN mkdir -p ${CI_BASE}/inputartifact
# rout make task
RUN make ci-build