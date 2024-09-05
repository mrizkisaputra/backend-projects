package helper

import "github.com/sirupsen/logrus"

const IN_PROGRESS = "in-progress"
const DONE = "done"
const TODO = "todo"

func LogIfError(log *logrus.Logger, err error) {
	if err != nil {
		log.Error(err)
		return
	}
}
