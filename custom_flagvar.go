package main

import (
	"errors"
	"fmt"
	"strings"
)

type Job string

func (j *Job) String() string {
	return fmt.Sprint(*j)
}

func (j *Job) Set(v string) error {
	if len(*j) > 0 {
		return errors.New("job flag already set")
	}
	*j = Job(strings.Join([]string{"job:", v}, " "))

	return nil
}
