FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/user-api
WORKDIR /go/src/user-api

# build models
RUN go get ./database
RUN go get ./
RUN go build

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080