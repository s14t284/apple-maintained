package usecase

import (
	"fmt"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// WatchInteractor apple watchの情報をやりとりするためのgateway
type WatchInteractor struct {
	wr database.WatchRepositoryImpl
}

// NewWatchInteractor WatchInteractorを生成
func NewWatchInteractor(dbClient *infrastructure.SQLClient) *WatchInteractor {
	return &WatchInteractor{wr: database.WatchRepositoryImpl{SQLClient: dbClient}}
}

// FindWatchAll apple watchの情報を取得
func (interactor *WatchInteractor) FindWatchAll() (model.Watches, error) {
	watches, err := interactor.wr.FindWatchAll()
	return watches, err
}

// FindByURL 指定したURLを持つapple watchを取得
func (interactor *WatchInteractor) FindByURL(url string) (*model.Watch, error) {
	return interactor.wr.FindByURL(url)
}

// IsExist 指定したオブジェクトが存在するかを取得
func (interactor *WatchInteractor) IsExist(watch *model.Watch) (bool, uint, time.Time, error) {
	return interactor.wr.IsExist(watch)
}

// AddWatch apple watchの情報を追加
func (interactor *WatchInteractor) AddWatch(watch *model.Watch) (err error) {
	err = interactor.wr.AddWatch(watch)
	return
}

// UpdateWatch apple watchの情報の更新
func (interactor *WatchInteractor) UpdateWatch(watch *model.Watch) (err error) {
	if watch.ID <= 0 {
		return fmt.Errorf("cannot logical remove watch because invalid watch id: %d", watch.ID)
	}
	err = interactor.wr.UpdateWatch(watch)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *WatchInteractor) UpdateAllSoldTemporary() (err error) {
	err = interactor.wr.UpdateAllSoldTemporary()
	return
}

// RemoveWatch apple watchの情報を削除
func (interactor *WatchInteractor) RemoveWatch(id int64) (err error) {
	err = interactor.wr.RemoveWatch(id)
	return
}
