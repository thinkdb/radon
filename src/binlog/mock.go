/*
 * Radon
 *
 * Copyright (c) 2017 QingCloud.com.
 * Code is licensed under the GPLv3.
 *
 */

package binlog

import (
	"errors"
	"github.com/thinkdb/radon/src/xbase"
)

var (
	_ xbase.RotateFile = &mockRotateFile{}
)

type mockRotateFile struct {
}

func (mf *mockRotateFile) Write(b []byte) (int, error) {
	return 0, errors.New("mock.rfile.write.error")
}

func (mf *mockRotateFile) Sync() error {
	return nil
}

func (mf *mockRotateFile) Close() {
}

func (mf *mockRotateFile) Name() string {
	return ""
}

func (mf *mockRotateFile) GetOldLogInfos() ([]xbase.LogInfo, error) {
	return nil, errors.New("mock.rfile.GetOldLogInfos.error")
}

func (mf *mockRotateFile) GetNextLogInfo(logName string) (xbase.LogInfo, error) {
	return xbase.LogInfo{}, errors.New("mock.rfile.GetOldLogInfos.error")
}

func (mf *mockRotateFile) GetCurrLogInfo(ts int64) (xbase.LogInfo, error) {
	return xbase.LogInfo{}, errors.New("mock.rfile.GetCurrLogInfo.error")
}
