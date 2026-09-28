package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	sliceInfo := strings.Split(data, ",")
	if len(sliceInfo) != 3 {
		return 0, "", 0, fmt.Errorf("Неправильный формат данных")
	}

	steps, err := strconv.Atoi(sliceInfo[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps < 0 {
		return 0, "", 0, fmt.Errorf("Отрицательное количество шагов")
	}

	walkTime, err := time.ParseDuration(sliceInfo[2])
	if err != nil {
		return 0, "", 0, err
	}
	if walkTime < 0 {
		return 0, "", 0, fmt.Errorf("Отрицательное время")
	}

	typeAct := sliceInfo[1]

	return steps, typeAct, walkTime, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration < 0 {
		return 0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeAct, walkTime, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var dist, speed, calories float64

	switch typeAct {
	case "Ходьба":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, walkTime)
		calories, err = RunningSpentCalories(steps, weight, height, walkTime)
		if err != nil {
			return "", err
		}
	case "Бег":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, walkTime)
		calories, err = WalkingSpentCalories(steps, weight, height, walkTime)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.Дистанция: %.2f км.Скорость: %.2f км/чСожгли калорий: %.2f", typeAct, walkTime.Hours(), dist, speed, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0, fmt.Errorf("Отрицательное число шагов")
	}
	if weight < 20 || weight > 300 {
		return 0, fmt.Errorf("Неверный вес")
	}
	if height < 0.5 || height > 3 {
		return 0, fmt.Errorf("Неверный рост")
	}
	if duration < 0 {
		return 0, fmt.Errorf("Отрицательное время")
	}

	speed := meanSpeed(steps, height, duration)

	calories := weight * speed * duration.Minutes() / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0, fmt.Errorf("Отрицательное число шагов")
	}
	if weight < 20 || weight > 300 {
		return 0, fmt.Errorf("Неверный вес")
	}
	if height < 0.5 || height > 3 {
		return 0, fmt.Errorf("Неверный рост")
	}
	if duration < 0 {
		return 0, fmt.Errorf("Отрицательное время")
	}

	speed := meanSpeed(steps, height, duration)

	calories := weight * speed * duration.Minutes() / minInH

	return calories * walkingCaloriesCoefficient, nil
}
