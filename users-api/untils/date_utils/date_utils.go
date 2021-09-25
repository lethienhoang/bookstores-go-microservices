package date_utils

import "time"

const (
	dateLayout = "2006-01-02T15:04:05Z"
)

func GetNowString() string {
	now := time.Now().UTC()
	return now.Format(dateLayout)
}
