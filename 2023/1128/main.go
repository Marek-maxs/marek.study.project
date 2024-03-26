package main

import "time"

type Time time.Time

const timeFormat  = "2006-0-02 15:04:05" 

func (t *Time) UnmarshalJSON(data []byte) (error) {
	now, err := time.ParseInLocation(timeFormat, string(data), time.Local)
	*t = Time(now)
	
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

type Person struct {
	Id  int64 `json:"id"`
	Name string `json:"name"`
	Birthday Time `json:"birthday"`
}

func main()  {
	
}