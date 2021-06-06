package repository

import (
    "github.com/chasegawa1209/linebot-todomaru/lib/infra/linestore"

    "go.uber.org/zap"
)

// NotifyChangeTowelRepositoryInterface リポジトリのインターフェース
type NotifyChangeTowelRepositoryInterface interface {
    PostMessage(msg string) error
}

// NotifyChangeTowelRepository リポジトリの実装
type NotifyChangeTowelRepository struct {
    logger    *zap.Logger
    lineStore linestore.LineStoreInterface
}

// NewNotifyChangeTowelRepository コンストラクタ
func NewNotifyChangeTowelRepository(logger *zap.Logger, lineStore linestore.LineStoreInterface) *NotifyChangeTowelRepository {
    return &NotifyChangeTowelRepository{
        logger:    logger,
        lineStore: lineStore,
    }
}

func (r *NotifyChangeTowelRepository) PostMessage(msg string) error {
    if err := r.lineStore.Post(msg); err != nil {
        r.logger.Error(err.Error())
        return err
    }
    return nil
}
