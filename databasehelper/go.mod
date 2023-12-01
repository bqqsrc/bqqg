module github.com/bqqsrc/bqqg/databasehelper

go 1.20

require (
	github.com/bqqsrc/bqqg/sqlfmt v0.0.0	
	github.com/bqqsrc/bqqg/database v0.0.0
	github.com/bqqsrc/bqqg/log v0.0.0	
)

replace (
	github.com/bqqsrc/bqqg/sqlfmt v0.0.0 => ../sqlfmt
	github.com/bqqsrc/bqqg/database v0.0.0 => ../database
	github.com/bqqsrc/bqqg/log v0.0.0	 => ../log
)