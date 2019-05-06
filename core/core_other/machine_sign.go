package core_other

import (
	"encoding/json"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lruio"
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/model"
	"lemon-robot-server/utils/workspace_util"
)

const machineSignSaveFile = "lr.msign"

var machineSignCache string

func GetMachineSign() (string, error) {
	if machineSignCache == "" {
		machineSign, err := initMachineSign()
		if err != nil {
			return "", err
		}
		machineSignCache = machineSign
	}
	return machineSignCache, nil
}

func initMachineSign() (string, error) {
	machineSign := readMachineSignFromLocal()
	if machineSign == "" {
		logger.Info("Machine Sign does not exist locally, start generating...")
		machineSign = lrustring.Uuid()
		saveErr := saveMachineSignToLocal(machineSign)
		if saveErr != nil {
			return "", saveErr
		}
		logger.Info("Machine Sign generation is completed: " + machineSign)
	}
	return machineSign, nil
}

func saveMachineSignToLocal(machineSign string) error {
	signObj := &model.LrMachineSign{
		Sign: machineSign,
	}
	jsonBytes, _ := json.Marshal(signObj)
	return lruio.ReplaceStrToFile(string(jsonBytes), workspace_util.GetWorkspacePath(machineSignSaveFile))
}

func readMachineSignFromLocal() string {
	signObj := model.LrMachineSign{}
	err := lruio.JsonToStruct(workspace_util.GetWorkspacePath(machineSignSaveFile), &signObj)
	if err != nil {
		return ""
	}
	return signObj.Sign
}
