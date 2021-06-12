package usecase

import (
	"time"

	"github.com/chasegawa1209/linebot-todomaru/lib/domain/repository"
)

// NotifyChangeTowelUsecaseInterface ユースケースのインターフェース
type NotifyChangeTowelUsecaseInterface interface {
    Exec() *NotifyChangeTowelUsecaseResult
}

// NotifyChangeTowelUsecase ユースケースの実装
type NotifyChangeTowelUsecase struct {
    repository repository.NotifyChangeTowelRepositoryInterface
}

// NewNotifyChangeTowelUsecase コンストラクタ
func NewNotifyChangeTowelUsecase(repository repository.NotifyChangeTowelRepositoryInterface) *NotifyChangeTowelUsecase {
    return &NotifyChangeTowelUsecase{
        repository: repository,
    }
}

// NotifyChangeTowelUsecaseResult usecase result
type NotifyChangeTowelUsecaseResult struct {
    Err            error
    ProcessingTime float64
}

func (u *NotifyChangeTowelUsecase) Exec() *NotifyChangeTowelUsecaseResult {
    result := &NotifyChangeTowelUsecaseResult{}
    now := time.Now()
    defer func() {
        result.ProcessingTime = time.Since(now).Seconds()
    }()

    if now.Weekday() == time.Saturday {
        // メッセージ作成
        msg := createMessage()

        // LineBotで送信
        if err := u.repository.PostMessage(msg); err != nil {
            result.Err = err
            return result
        }
    }

    return result
}

func createMessage() string {
    message := "今日は土曜日！\nタオルを変えるんだまる！"
    return message
}
