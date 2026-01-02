package daysteps

import (
	"errors"
	"fmt"
	"log"
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
	// парсим строку
	separation := strings.Split(data, ",")
	if len(separation) != 2 {
		return 0, 0, errors.New("что-то пошло не так")
	}

	// преопбразуем шаги из строки в Int
	steps, err := strconv.Atoi(separation[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("колличество шагов должно быть больше 0")
	}

	// преобразуем время из строки в time.Duration
	duration, err := time.ParseDuration(separation[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("время должно быть больше 0")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	// парсим строку с помощью функции parsePackage()
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	// вычисляем дистанцию
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	// вычисляем колличество калорий
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Panicln(err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
	return result
}
