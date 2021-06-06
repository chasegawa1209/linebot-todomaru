package model

import "time"

// RequestModel スケジューラから受け取るリクエスト
type RequestModel struct {
    RequestData *RequestData `json:"data"`
}

// スケジューラから受け取るリクエスト内容
type RequestData struct {
    Action string    `json:"action"`
    Date   time.Time `json:"date"`
}
