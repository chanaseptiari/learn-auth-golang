package helper

import "log"

func check_error(err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
