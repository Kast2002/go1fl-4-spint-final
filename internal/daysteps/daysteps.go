package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	stepsAndWalk := strings.Split(data, ",")
	if len(stepsAndWalk) != 2 {
		return 0, 0, fmt.Errorf("Неправильный формат данных")
	}

	steps, err := strconv.Atoi(stepsAndWalk[0])
	if err != nil {
		return 0, 0, err
	}
	if steps < 0 {
		return 0, 0, fmt.Errorf("Отрицательное количество шагов")
	}

	walkTime, err := time.ParseDuration(stepsAndWalk[1])
	if err != nil {
		return 0, 0, err
	}
	if walkTime < 0 {
		return 0, 0, fmt.Errorf("Отрицательное время")
	}

	return steps, walkTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, walkTime, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// Не очень понимаю, зачем эта проверка, если она реализована в parsePackage
	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * stepLength / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkTime)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.", steps, distance, calories)
}
