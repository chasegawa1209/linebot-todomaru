package interactor

import (
	"github.com/chasegawa1209/linebot-todomaru/lib/handler"
	"github.com/chasegawa1209/linebot-todomaru/lib/domain/repository"
	"github.com/chasegawa1209/linebot-todomaru/lib/infra/linestore"
	"github.com/chasegawa1209/linebot-todomaru/lib/usecase"

	"go.uber.org/zap"
)

// Interactor インタラクタ
type Interactor struct {
    logger          *zap.Logger
    lineAccessToken string
    lineSecret      string
    lineRoomID      string
}

// NewInteractor コンストラクタ
func NewInteractor(logger *zap.Logger, lineAccessToken string, lineSecret string, lineRoomID string) *Interactor {
    return &Interactor{
        logger:          logger,
        lineAccessToken: lineAccessToken,
        lineSecret:      lineSecret,
        lineRoomID:      lineRoomID,
    }
}

// NewNotifyChangeTowelHandler ハンドラー
func (i *Interactor) NewNotifyChangeTowelHandler() handler.NotifyChangeTowelHandlerInterface {
    return handler.NewNotifyChangeTowelHandler(
        i.NewNotifyChangeTowelUsecase(),
    )
}

// NewNotifyChangeTowelUsecase ユースケース
func (i *Interactor) NewNotifyChangeTowelUsecase() usecase.NotifyChangeTowelUsecaseInterface {
    return usecase.NewNotifyChangeTowelUsecase(
        i.NewNotifyChangeTowelRepository(),
    )
}

// NewNotifyChangeTowelRepository リポジトリ
func (i *Interactor) NewNotifyChangeTowelRepository() repository.NotifyChangeTowelRepositoryInterface {
    return repository.NewNotifyChangeTowelRepository(
        i.logger,
        i.NewLineStore(),
    )
}

// NewLineStore LINE接続
func (i *Interactor) NewLineStore() linestore.LineStoreInterface {
    linestore, err := linestore.NewLineStore(
        i.logger,
        i.lineAccessToken,
        i.lineSecret,
        i.lineRoomID,
    )
    if err != nil {
        panic(err)
    }
    return linestore
}
