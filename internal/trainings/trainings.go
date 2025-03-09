package trainings

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 3 {
		return errors.New("invalid data")
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}
	t.Steps = steps
	typeTrain := dataSlice[1]
	switch typeTrain {
	case "Ходьба":
		t.TrainingType = typeTrain
	case "Бег":
		t.TrainingType = typeTrain
	default:
		return errors.New("Неизвестный тип тренировки")
	}
	duration, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return err
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)
	if t.Duration <= 0 {
		return "", errors.New("invalid duration")
	}
	mSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)
	var result string
	switch t.TrainingType {
	case "Ходьба":
		walkCall := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		result = fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			t.TrainingType, t.Duration.Hours(), distance, mSpeed, walkCall)
	case "Бег":
		runCall := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
		result = fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			t.TrainingType, t.Duration.Hours(), distance, mSpeed, runCall)
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}
	return result, nil

}
