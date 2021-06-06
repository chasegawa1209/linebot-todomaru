package main

import (
	"fmt"
	"os"

	"github.com/chasegawa1209/linebot-todomaru/app/notify_change_towel/interactor"
	"github.com/chasegawa1209/linebot-todomaru/lib/infra/logging"
)

func main() {
    LINE_ACCESS_TOKEN := os.Getenv("LINE_ACCESS_TOKEN")
    LINE_SECRET := os.Getenv("LINE_SECRET")
    LINE_ROOM_ID := os.Getenv("LINE_ROOM_ID")

    fmt.Println("Hello, world")

    // logger
    isDebug := true
    logger := logging.NewZapLogger(isDebug)

    i := interactor.NewInteractor(logger, LINE_ACCESS_TOKEN, LINE_SECRET, LINE_ROOM_ID)
    result := i.NewNotifyChangeTowelUsecase().Exec()
    logger.Sugar().Infof("ProcessingTime: %f[s]", result.ProcessingTime)
    if result.Err != nil {
        logger.Sugar().Fatal("failed to NotifyChangeTowelBatch: %s", result.Err.Error())
    }
    logger.Sugar().Infof("success to NotifyChangeTowelBatch")
}
