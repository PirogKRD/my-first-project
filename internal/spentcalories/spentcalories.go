package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	// lenStep                    = 0.65 // средняя длина шага. не понимаю для чего тут эта константа, ведь для каждых вводных мы расчитываем их отдельно
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	// парсим строку
	seperation := strings.Split(data, ",")
	if len(seperation) != 3 {
		return 0, "", 0, errors.New("неверное количество элементов")
	}

	// преобразование шагов в int
	steps, err := strconv.Atoi(seperation[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("количество шагов должно быть больше 0")
	}

	// Вид активности
	activity := seperation[1]

	// преобразование времени
	duration, err := time.ParseDuration(seperation[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("время должно быть больше 0")
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// вычисляем длинну шага
	stepLen := height * stepLengthCoefficient

	// дистанция в метрах
	distanceMeters := float64(steps) * stepLen

	// дистанция в километрах
	distanceKm := distanceMeters / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distKm := distance(steps, height)

	// вычисляем среднюю скорость
	averageSpeed := distKm / duration.Hours()
	return averageSpeed
}

// строка с информацией о тренировке
func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var dist float64
	var speed float64
	var calories float64
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	switch activity {
	case "Бег":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		err := errors.New("неизвестный тип тренировки")
		return "", err
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, duration.Hours(), dist, speed, calories)
	return result, nil
}

// количество калорий, потраченных при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные данные для бега")
	}
	// средняя скорость
	averSpeed := meanSpeed(steps, height, duration)

	// время в минутах
	durationMin := duration.Minutes()
	calories := (weight * averSpeed * durationMin) / minInH
	return calories, nil
}

// количество калорий, потраченных при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные данные для хотьбы")
	}
	// средняя скорость
	averSpeed := meanSpeed(steps, height, duration)
	// время в минутах
	durationMin := duration.Minutes()
	calories := (weight * averSpeed * durationMin) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
