package utils

import "time"

func ParseTime(str string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, str)
	if err != nil {
		layout := "2006-01-02T15:04:05-07:00"
		t, _ := time.Parse(layout, str)

		return t
	}

	return t
}

func Max(data map[string]int) float64 {
	var max int
	for _, v := range data {
		if v >= max {
			max = v
		}
	}

	return float64(max)
}

func Min(data map[string]int) float64 {
	min := int(^uint(0) >> 1)
	for _, v := range data {
		if v <= min {
			min = v
		}
	}

	return float64(min)
}

func Average(data map[string]int) float64 {
	var sum, counter int
	for _, v := range data {
		sum += v
		counter++
	}

	return float64(sum / counter)
}

func Median(data map[string]int) float64 {
	var arr []int
	for _, v := range data {
		arr = append(arr, v)
	}

	counter := len(arr)

	if counter+1%2 == 0 {
		a := arr[(counter / 2)]
		b := arr[(counter/2)+1]
		return float64((a + b) / 2)
	} else {
		return float64(arr[counter/2])
	}
}
