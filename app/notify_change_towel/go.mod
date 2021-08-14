module github.com/chasegawa1209/linebot-todomaru/app/notify_change_towel

// +heroku goVersion go1.16
go 1.16

replace github.com/chasegawa1209/linebot-todomaru/lib => ./../../lib

require (
	github.com/chasegawa1209/linebot-todomaru/lib v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.16.0
)
