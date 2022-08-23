// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/weplanx/collector/app"
	"github.com/weplanx/collector/common"
)

// Injectors from wire.go:

func NewApp() (*app.App, error) {
	values, err := LoadValues()
	if err != nil {
		return nil, err
	}
	logger, err := UseZap()
	if err != nil {
		return nil, err
	}
	conn, err := UseNats(values)
	if err != nil {
		return nil, err
	}
	jetStreamContext, err := UseJetStream(conn)
	if err != nil {
		return nil, err
	}
	objectStore, err := UseStore(values, jetStreamContext)
	if err != nil {
		return nil, err
	}
	client := UseInflux(values)
	inject := &common.Inject{
		Values: values,
		Log:    logger,
		Js:     jetStreamContext,
		Store:  objectStore,
		Influx: client,
	}
	appApp := app.Initialize(inject)
	return appApp, nil
}