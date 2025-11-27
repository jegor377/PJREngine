package GameEngine

import "github.com/veandco/go-sdl2/sdl"

func (engine *GameEngine) checkEventsForButtons(e *sdl.MouseButtonEvent) {
	for _, v := range engine.buttonMgr.GetButtons() {
		if v.HasBeenClicked(e.X, e.Y, e.Button, e.State) { // e.X, e.Y, e.Button, e.State
			dialogOption, err := engine.findDialogOption(v.GetSource())
			if err != nil {
				panic(err)
			}
			if engine.appState, err = dialogOption.DoJob(engine.appState); err != nil {
				panic(err)
			}
			dialog, err := engine.findDialog(v.GetTarget())
			if err != nil {
				panic(err)
			}
			if engine.appState, err = dialog.DoJob(engine.appState); err != nil {
				panic(err)
			}
			engine.appState.IsEnding = dialog.IsEnding()
			engine.appState.SetActualDialogId(v.GetTarget())
		}
	}
}
