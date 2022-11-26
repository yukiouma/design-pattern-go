package biz

import "errors"

var (
	ErrSetSalary   = errors.New("error: you can not change salary for member from other department")
	ErrNewDirector = errors.New("error: you can not create new director using group leader from other department")
)
