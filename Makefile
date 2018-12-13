include .env
export

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/volatiletech/mig/...
mig-create:
	mig create $(title) -d mysql_migrations
mig-up:
	mig up mysql "$$MYSQL_USER:$$MYSQL_PASSWORD@/$$MYSQL_DATABASE" -d mysql_migrations
mig-upone:
	mig upone mysql "$$MYSQL_USER:$$MYSQL_PASSWORD@/$$MYSQL_DATABASE" -d mysql_migrations
mig-down:
	mig down mysql "$$MYSQL_USER:$$MYSQL_PASSWORD@/$$MYSQL_DATABASE" -d mysql_migrations
mig-downall:
	mig downall mysql "$$MYSQL_USER:$$MYSQL_PASSWORD@/$$MYSQL_DATABASE" -d mysql_migrations
docker-start:
	sudo docker-compose start
docker-stop:
	sudo docker-compose stop
dep:
	dep ensure
build:
	go build -o ./cassiopeia ./app/main.go
start-only:
	./cassiopeia
build-start: build start-only
