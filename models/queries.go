package models

var SQL_JOINER = `
	select user_id from users where gmail = $1
`

var SQL_INSERT = `
	insert into users(username, gmail, password)
		values ($1, $2, $3)
`

var SQL_LOGIN = `
	select user_id from users
	where (username = $1 or gmail = $1 ) and password = $2 and token is not null
`

var SQL_ADD_TOKEN = `
	update users set token = $3
	where (username = $1 or gmail = $1 ) and password = $2 
	returning
		token,
		username,
		gmail
`

var SQL_SELECT_USER = `
	select
		token,
		username,
		gmail
	from users
	where user_id = $1
`

var SQL_CHANGE_PASS = `
	update users set password = $3
	where password != $3 and (username = $1 or gmail = $1 ) and password = $2 and $2 != $3
	returning user_id 
`

var SQL_CHANGE_TOKEN = `
	update users set token = $3
	where (username = $1 or gmail = $1 ) and password = $2 
	returning
		token,
		username,
		gmail
`