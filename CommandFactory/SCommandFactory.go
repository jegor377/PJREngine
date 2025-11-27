package CommandFactory

import "errors"
import "pjre/Commands"
import "strconv"

type SCommandFactory struct {
	commands map[uint8]SCommand
}

func NewS() *SCommandFactory {
	commands := map[uint8]SCommand {
		0: Commands.NewBgColorCommand(),
		1: Commands.NewAppNameCommand(),
		2: Commands.NewWYDTCommand(),
		8: Commands.NewBgSongCommand(),
	}
	return &SCommandFactory{commands}
}

func (cmdFactory *SCommandFactory) Get(key uint8) (SCommand, error) {
	if val, ok := cmdFactory.commands[key]; ok {
		return val, nil
	}
	errInfo := "Key [" + strconv.Itoa(int(key)) + "] does not exist. - SC"
	return nil, errors.New(errInfo)
}
