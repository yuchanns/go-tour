package timer

import "time"

var location *time.Location

func init() {
	location, _ = time.LoadLocation("Asia/Shanghai")
}

func GetNowTime() time.Time {
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
