FROM library/golang:1.16.2

# Godep for vendoring

#ARUN  ["/bin/bash", "-c", "git clone https://gitee.com/zhuhongjie01/godep.git && go install"]

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/home/src/myblog
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./myblog)
ADD . $APP_DIR

# Compile the binary and statically link
#RUN cd $APP_DIR && CGO_ENABLED=0 godep go build -ldflags '-d -w -s'

EXPOSE 18080
