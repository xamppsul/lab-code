package config

import "time"

var location *time.Location

func SetTimezone(timezone string) error {
	time_location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err.Error())
	}
	location = time_location
	return nil
}

func GetTimezone(t time.Time) time.Time {
	return t.In(location)
}
