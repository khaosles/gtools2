package g

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/*
   @File: timestamputc.go
   @Author: khaosles
   @Time: 2023/6/16 14:25
   @Desc: json格式化返回 YYYY-mm-dd HH:MM:SS 格式的时间字段
*/

// 时间格式
type TimestampUTC time.Time

func (ts *TimestampUTC) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(`"`+timeLayout+`"`, string(data))
	if err != nil {
		return err
	}
	*ts = TimestampUTC(t)
	return nil
}

func (ts TimestampUTC) MarshalJSON() ([]byte, error) {
	t := time.Time(ts)
	formatted := fmt.Sprintf(`"%s"`, t.Format(timeLayout))
	return []byte(formatted), nil
}

func (ts TimestampUTC) Value() (driver.Value, error) {
	return time.Time(ts), nil
}

func (ts *TimestampUTC) Scan(value interface{}) error {
	if value == nil {
		*ts = TimestampUTC(time.Time{})
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*ts = TimestampUTC(t)
		return nil
	}
	return fmt.Errorf("failed to scan CustomTime value")
}
