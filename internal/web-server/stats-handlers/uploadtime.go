package statshandlers

import "time"

type UploadTime struct {
	count int
	avg   time.Duration
}

func (u *UploadTime) Record(d time.Duration) {
	u.count++
	u.avg = calcAvgUploadTime(u.avg, d, u.count)
}

func (u *UploadTime) HasData() bool {
	return u.count > 0
}

func (u *UploadTime) Avg() string {
	if u.count == 0 {
		return "0s"
	}
	return u.avg.String()
}

func calcAvgUploadTime(current, next time.Duration, count int) time.Duration {
	return current + (next-current)/time.Duration(count)
}
