package CommandFactory

import "errors"
import "pjre/Commands"
import "strconv"

type DCommandFactory struct {
	commands map[uint8]DCommand
}

func NewD() *DCommandFactory {
	commands := map[uint8]DCommand {
		3: Commands.NewDTextCommand(),
		4: Commands.NewDColorCommand(),
		5: Commands.NewOptionCommand(),
		6: Commands.NewNarratorCommand(),
		7: Commands.NewSpriteCommand(),
	}
	return &DCommandFactory{commands}
}

func (cmdFactory *DCommandFactory) Get(key uint8) (DCommand, error) {
	if val, ok := cmdFactory.commands[key]; ok {
		return val, nil
	}
	errInfo := "Key [" + strconv.Itoa(int(key)) + "] does not exist. - DC"
	return nil, errors.New(errInfo)
}
