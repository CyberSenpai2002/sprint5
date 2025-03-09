package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 2 {
		return errors.New("invalid data")
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}
	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return err
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("invalid duration")
	}
	distance := spentenergy.Distance(ds.Steps)
	walkCall := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила: %.2f км.\nВы Сожгли: %.2f ккал.",
		ds.Steps, distance, walkCall)
	return result, nil
}
