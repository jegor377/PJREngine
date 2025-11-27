package GameEngine

import "io/ioutil"
import "os"
import "pjre/ConvertingTools"

func (engine *GameEngine) saveState() {
	f, err := os.OpenFile("save.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if !engine.appState.IsEnding {
		convDialogId, err := ConvertingTools.ConvertUint32ToBytesLittleEndian(engine.appState.GetActualDialogId())
		if err != nil {
			panic(err)
		}
		f.Write(convDialogId)
	} else {
		convDialogId, err := ConvertingTools.ConvertUint32ToBytesLittleEndian(0)
		if err != nil {
			panic(err)
		}
		f.Write(convDialogId)
	}
	
	if err := f.Close(); err != nil {
		panic(err)
	}
}

func (engine *GameEngine) loadState() {
	if _, err := os.Stat("save.txt"); !os.IsNotExist(err) {
		file, err := ioutil.ReadFile("save.txt")
		if err != nil {
			panic(err)
		}
		dialogId, err := ConvertingTools.ConvertBytesToUint32LittleEndian(file)
		if err != nil {
			panic(err)
		}
		engine.appState.SetActualDialogId(dialogId)
	} else {
		engine.saveState()
	}
}
