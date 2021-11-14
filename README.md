# alura-go

learning go with alura

## dependencies

- [go](https://golang.org/doc/install)
- [docker](https://docs.docker.com/engine/install/)

## local development

Start database

`docker-compose up -d`

Install dependencies

`go mod tidy`

Run go script

`go run main.go`

## alura repo

This project was based on alura go course, here is [the source repo](https://github.com/alura-cursos/web_com_golang)

## Next steps

until this commit we followed, with some changes, the instructor code, following I want to do some other enhancements:

- [x] get db string from env vars
- [x] deployed [on heroku](https://alura-go.herokuapp.com/)
- [ ] clean/refactor code
