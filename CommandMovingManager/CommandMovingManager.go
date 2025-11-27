package CommandMovingManager

import "pjre/JumpingHandlers"
import "strconv"
import "errors"

type CommandMovingFactory struct {
	commands map[uint8]JumpingHandler
}

func New() *CommandMovingFactory {
	commands := map[uint8]JumpingHandler {
		0: JumpingHandlers.NewBgColorJumpingHandler(),
		1: JumpingHandlers.NewAppNameJumpingHandler(),
		2: JumpingHandlers.NewWYDTJumpingHandler(),
		3: JumpingHandlers.NewTextJumpingHandler(),
		4: JumpingHandlers.NewColorJumpingHandler(),
		5: JumpingHandlers.NewOptionJumpingHandler(),
		6: JumpingHandlers.NewNarratorJumpingHandler(),
		7: JumpingHandlers.NewSpriteJumpingHandler(),
		8: JumpingHandlers.NewBgSongJumpingHandler(),
	}
	return &CommandMovingFactory{commands}
}

func (jmpFactory *CommandMovingFactory) Get(key uint8) (JumpingHandler, error) {
	if val, ok := jmpFactory.commands[key]; ok {
		return val, nil
	}
	errInfo := "Key [" + strconv.Itoa(int(key)) + "] does not exist. - CMM"
	return nil, errors.New(errInfo)
}
