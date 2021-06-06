package handler

import (
	"go.uber.org/zap"
	"fmt"
	"github.com/chasegawa1209/linebot-todomaru/lib/usecase"
	"net/http"
)

// NotifyChangeTowelHandlerInterface ハンドラーのインターフェース
type NotifyChangeTowelHandlerInterface interface {
    Handler(w http.ResponseWriter, r *http.Request)
    LineHandler(w http.ResponseWriter, r *http.Request)
}

// NotifyChangeTowelHandler ハンドラーの実装
type NotifyChangeTowelHandler struct {
    logger  *zap.Logger
    usecase usecase.NotifyChangeTowelUsecaseInterface
}

// NewNotifyChangeTowelHandler コンストラクタ
func NewNotifyChangeTowelHandler(logger *zap.Logger, usecase usecase.NotifyChangeTowelUsecaseInterface) *NotifyChangeTowelHandler {
    return &NotifyChangeTowelHandler{
        logger:  logger,
        usecase: usecase,
    }
}

func (h *NotifyChangeTowelHandler) Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!\n")
    w.WriteHeader(http.StatusOK)
}

func (h *NotifyChangeTowelHandler) LineHandler(w http.ResponseWriter, r *http.Request) {
    result := h.usecase.Exec()
    h.logger.Sugar().Infof("ProcessingTime: %f[s]", result.ProcessingTime)
    if result.Err != nil {
        h.logger.Sugar().Fatal("failed to NotifyChangeTowelBatch: %s", result.Err.Error())
    }
    h.logger.Sugar().Infof("success to NotifyChangeTowelBatch")
    w.WriteHeader(http.StatusOK)
}
