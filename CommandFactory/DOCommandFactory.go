package CommandFactory

import "errors"
import "pjre/Commands"
import "strconv"

type DOCommandFactory struct {
	commands map[uint8]DOCommand
}

func NewDO() *DOCommandFactory {
	commands := map[uint8]DOCommand {
		3: Commands.NewDOTextCommand(),
		4: Commands.NewDOColorCommand(),
	}
	return &DOCommandFactory{commands}
}

func (cmdFactory *DOCommandFactory) Get(key uint8) (DOCommand, error) {
	if val, ok := cmdFactory.commands[key]; ok {
		return val, nil
	}
	errInfo := "Key [" + strconv.Itoa(int(key)) + "] does not exist. - DOC"
	return nil, errors.New(errInfo)
}
