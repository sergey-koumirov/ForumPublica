# run dev
server
`go run ./server/app.go`

run import
`go run ./sde/app.go --file ~/Downloads/sde-20180713-TRANQUILITY.zip`

# db & migrations

`create database publica;`
`create user 'publica'@'localhost' identified by 'publica';`
`grant all privileges on publica.* to 'publica'@'localhost';`
