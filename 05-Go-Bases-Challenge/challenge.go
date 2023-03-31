package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type TicketInformation struct {
	ID          string
	Name        string
	Email       string
	Destination string
	Time        string
	Price       string
}

const (
	MADRUGADA = "madrugada"
	MANANA    = "manana"
	TARDE     = "tarde"
	NOCHE     = "noche"
)

var tickets []TicketInformation

func ReadFileAndSaveByObject(filename string) error {
	file, err := os.Open("tickets.csv")

	defer file.Close()
	if err != nil {
		return errors.New("File not exist")
	}

	text, errRead := io.ReadAll(file)
	if errRead != nil {
		return errors.New("Cant read file")
	}

	splited_text_by_line := strings.Split(string(text), "\n")

	for _, value := range splited_text_by_line {
		splited_value := strings.Split(value, ",")
		ticket := TicketInformation{
			ID:          splited_value[0],
			Name:        splited_value[1],
			Email:       splited_value[2],
			Destination: splited_value[3],
			Time:        splited_value[4],
			Price:       splited_value[5],
		}

		tickets = append(tickets, ticket)
	}

	return nil
}

func GetTotalTickets(destination string) (int, error) {

	destination_count := 0
	for _, value := range tickets {
		if value.Destination == destination {
			destination_count++
		}
	}
	return destination_count, nil

}

func GetCountByPeriod(time string) (int, error) {

	min_time := 0
	max_time := 0
	switch time {
	case MADRUGADA:
		max_time = 6
	case MANANA:
		min_time = 7
		max_time = 12
	case TARDE:
		min_time = 13
		max_time = 19
	case NOCHE:
		min_time = 20
		max_time = 23
	default:
		return 0, errors.New("Time not exist, choose some of this: (madrugada, manana, tarde, noche)")
	}

	people_count := 0
	for _, value := range tickets {

		ticket_time_splited := strings.Split(value.Time, ":")
		value, err_conv := strconv.Atoi(ticket_time_splited[0])

		if err_conv != nil {
			return 0, errors.New("Error during conversion")
		}

		if value >= min_time && value <= max_time {
			people_count++
		}
	}

	return people_count, nil
}

func AverageDestination(destination string, total int) (float64, error) {
	// Is not specified the total function
	// The average can be calculated with the GetTotalTickets function

	people_count, err := GetTotalTickets(destination)
	if err != nil {
		return 0.0, err
	}
	return float64(people_count / total), nil
}

func main() {

	err_by_save := ReadFileAndSaveByObject("tickets.csv")
	if err_by_save != nil {
		fmt.Println("Error: ", err_by_save)
	}

	// Requirement 1
	a, err := GetTotalTickets("Mexico")
	if err != nil {
		fmt.Println("Requirement 1: ", err)
	}
	fmt.Println("Requirement 1 Result: People Count: ", a)

	// Requirement 2
	b, err_b := GetCountByPeriod("madrugada")
	if err_b != nil {
		fmt.Println("Requirement 2: ", err_b)
	}
	fmt.Println("Requirement 2 Result: People Count: ", b)

	//Requirement 3
	c, err_c := AverageDestination("mexico", 10)
	if err_c != nil {
		fmt.Println("Requirement 3: ", err_c)
	}
	fmt.Println("Requirement 3 Result: Average: ", c)

	//Tests
}
