# run dev
`vgo run ./server/app.go`


# db & migrations

`create database publica;`
`create user 'publica'@'localhost' identified by 'publica';`
`grant all privileges on publica.* to 'publica'@'localhost';`
