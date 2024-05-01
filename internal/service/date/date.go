package date

import (
	"errors"
	"go_final_project/internal/util"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	notFoundRule = "not found repeat rule"
	badRule      = "bad repeat rule"
)

var rules = map[string]func(now time.Time, date, repeat string) (string, error){
	"^d\\s\\d{1,3}$":            dayRule,
	"^y$":                       yearRule,
	"^w\\s[1-7]?(,[1-7]){0,6}$": weekRule,
	"^m\\s-?\\d+(,-?\\d+){0,30}(\\s\\d+(,\\d+){0,11})?$": monthRule,
}

func months(months []string) ([12]bool, error) {
	ans := [12]bool{}
	for _, month := range months {
		monthIndex, _ := strconv.Atoi(month)
		if monthIndex > 12 || monthIndex < 1 {
			return ans, errors.New(badRule)
		}
		ans[monthIndex-1] = true
	}
	return ans, nil
}

func inSet(now time.Time, daysSet map[int]bool) bool {
	if daysSet[now.Day()] {
		return true
	}
	year, month, _ := now.Date()
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return daysSet[-1] && firstDayOfMonth.AddDate(0, 1, -1).Equal(now) ||
		daysSet[-2] && firstDayOfMonth.AddDate(0, 1, -2).Equal(now)
}

func monthRule(now time.Time, date, repeat string) (string, error) {
	splitedRepeat := strings.Split(repeat, " ")
	monthDays := strings.Split(splitedRepeat[1], ",")
	includeDaysSet := map[int]bool{}
	for _, monthDay := range monthDays {
		dayIndex, _ := strconv.Atoi(monthDay)
		if dayIndex < -2 || dayIndex > 31 {
			return "", errors.New(badRule)
		}
		includeDaysSet[dayIndex] = true
	}

	var includeMonths [12]bool
	if len(splitedRepeat) > 2 {
		var err error
		includeMonths, err = months(strings.Split(splitedRepeat[2], ","))
		if err != nil {
			return "", err
		}
	} else {
		includeMonths = [12]bool{}
		for i := range includeMonths {
			includeMonths[i] = true
		}
	}

	curDate, err := time.Parse(util.DateFormat, date)
	if err != nil {
		return "", err
	}

	if now.After(curDate) {
		curDate = now
	}

	curDate = curDate.AddDate(0, 0, 1)
	for !includeMonths[curDate.Month()-1] || !inSet(curDate, includeDaysSet) {
		curDate = curDate.AddDate(0, 0, 1)
	}
	return curDate.Format(util.DateFormat), nil
}

func weekRule(now time.Time, date, repeat string) (string, error) {
	weekDays := strings.Split(strings.Split(repeat, " ")[1], ",")
	weekDaysSet := map[int]bool{}
	for _, weekDay := range weekDays {
		dayIndex, _ := strconv.Atoi(weekDay)
		weekDaysSet[dayIndex%7] = true
	}

	curDate, err := time.Parse(util.DateFormat, date)
	if err != nil {
		return "", err
	}

	if now.After(curDate) {
		curDate = now
	}

	curDate = curDate.AddDate(0, 0, 1)

	for !weekDaysSet[int(curDate.Weekday())] {
		curDate = curDate.AddDate(0, 0, 1)
	}

	return curDate.Format(util.DateFormat), nil
}

func dayRule(now time.Time, date, repeat string) (string, error) {
	items := strings.Split(repeat, " ")

	days, err := strconv.Atoi(items[1])
	if err != nil {
		return "", err
	}

	if days > 400 || days < 1 {
		return "", errors.New(badRule)
	}

	curDate, err := time.Parse(util.DateFormat, date)
	if err != nil {
		return "", err
	}

	for next := true; next; next = !curDate.After(now) {
		curDate = curDate.AddDate(0, 0, days)
	}

	return curDate.Format(util.DateFormat), nil
}

func yearRule(now time.Time, date, _ string) (string, error) {

	curDate, err := time.Parse(util.DateFormat, date)
	if err != nil {
		return "", err
	}

	for next := true; next; next = !curDate.After(now) {
		curDate = curDate.AddDate(1, 0, 0)
	}

	return curDate.Format(util.DateFormat), nil
}

func NextDate(now time.Time, date string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", nil
	}
	for r, f := range rules {
		re := regexp.MustCompile(r)
		if re.MatchString(repeat) {
			result, err := f(now, date, repeat)
			if err != nil {
				return "", err
			}
			return result, nil
		}
	}
	return "", errors.New(notFoundRule)
}
