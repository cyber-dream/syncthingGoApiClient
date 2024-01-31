package main

import (
	"SyncthingGoApiClient/syncapi"
	"context"
)

func main() {
	a, err := syncapi.NewApiClient("MPQoqNpjAKggVAr4spFCmuspy99UkHWd", "http://127.0.0.1:8384")
	if err != nil {
		panic(err.Error())
	}
	resBr, err := a.GetSystemBrowse(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resConn, err := a.GetSystemConnections(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resDebug, err := a.GetSystemDebug(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resERr, err := a.GetSystemError(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resLog, err := a.GetSystemLog(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resPaths, err := a.GetSystemPaths(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resPing, err := a.GetSystemPing(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resStat, err := a.GetSystemStatus(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resUpgr, err := a.GetSystemUpgrade(context.Background())
	if err != nil {
		panic(err.Error())
	}
	resVer, err := a.GetSystemVersion(context.Background())
	if err != nil {
		panic(err.Error())
	}
	_ = resBr
	_ = resConn
	_ = resDebug
	_ = resERr
	_ = resLog
	_ = resPaths
	_ = resPing
	_ = resStat
	_ = resUpgr
	_ = resVer

	config, err := a.GetConfig(context.Background())
	if err != nil {
		panic(err.Error())
	}
	restReq, err := a.GetConfigRestartRequired(context.Background())
	if err != nil {
		panic(err.Error())
	}
	ConfFolders, err := a.GetConfigFolders(context.Background())
	if err != nil {
		panic(err.Error())
	}
	ConfDevices, err := a.GetConfigDevices(context.Background())
	if err != nil {
		panic(err.Error())
	}

	_ = config
	_ = restReq
	_ = ConfFolders
	_ = ConfDevices
	println("")
}
