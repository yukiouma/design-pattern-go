package behavioralcommand

import (
	"behavioral-command/internal/biz"
	"testing"
)

func TestCommand(t *testing.T) {
	car := biz.NewCar()
	app := biz.NewApp()
	turnLeft := biz.NewTurnLeftCommand(car, app)
	turnRight := biz.NewTurnRightCommand(car, app)
	undo := biz.NewUndoCommand(car, app)

	turnLeft.Execute()
	turnRight.Execute()
	undo.Execute()
	if car.Direction() != biz.Left {
		t.Errorf("expect left, got %s", biz.DirectionString[car.Direction()])
		return
	}

	turnRight.Execute()
	turnRight.Execute()
	turnRight.Execute()
	if car.Direction() != biz.Down {
		t.Errorf("expect down, got %s", biz.DirectionString[car.Direction()])
		return
	}

	turnLeft.Execute()
	turnLeft.Execute()
	if car.Direction() != biz.Up {
		t.Errorf("expect up, got %s", biz.DirectionString[car.Direction()])
		return
	}

	undo.Execute()
	undo.Execute()
	if car.Direction() != biz.Down {
		t.Errorf("expect down, got %s", biz.DirectionString[car.Direction()])
		return
	}
}
