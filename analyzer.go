package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}

func GroupPeopleByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job]++
	}
	return result
}

func Top5JobsByNumber(p []Person) {
	jobCount := GroupPeopleByJob(p)
	result := SortMapFromHighToLowByValue(jobCount)
	for i := 0; i < 5; i++ {
		fmt.Printf("%s - %d\n", result[i].Key, result[i].Value)
	}
}

func CountPeopleByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City]++
	}
	return result
}

type kv struct {
	Key   string
	Value int
}

func SortMapFromHighToLowByValue(m map[string]int) (ss []kv) {
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	return ss
}

func Top5CitiesByNumber(p []Person) {
	cityCount := CountPeopleByCity(p)
	result := SortMapFromHighToLowByValue(cityCount)
	for i := 0; i < 5; i++ {
		fmt.Printf("%s - %d\n", result[i].Key, result[i].Value)
	}
}

func GroupJobByCity(p []Person) (result map[string][]string) {
	result = make(map[string][]string)
	for _, person := range p {
		result[person.City] = append(result[person.City], person.Job)
	}
	return result
}

func CountNumberEachJob(input []string) (result map[string]int) {
	result = make(map[string]int)
	for _, job := range input {
		result[job]++
	}
	return result
}

func TopJobByNumerInEachCity(p []Person) (result map[string]map[string]int) {
	result = make(map[string]map[string]int)
	jobByCity := GroupJobByCity(p)
	for key, value := range jobByCity {
		resultValue := CountNumberEachJob(value)
		result[key] = resultValue
	}
	return result
}

func TotalSalaryInEachJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job] += person.Salary
	}
	return result
}

func AverageSalaryByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	groupJob := GroupPeopleByJob(p)
	salaryEachJob := TotalSalaryInEachJob(p)
	for key, value := range groupJob {
		for k, v := range salaryEachJob {
			if key == k {
				result[key] = v / value
			}
		}
	}
	return result
}

func TotalSalaryInEachCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City] += person.Salary
	}
	return result
}

func FiveCitiesHasTopAverageSalary(p []Person) {
	resultMap := make(map[string]int)
	var result []kv
	peopeByCity := CountPeopleByCity(p)
	salaryEachCity := TotalSalaryInEachCity(p)
	for key, value := range peopeByCity {
		for k, v := range salaryEachCity {
			if key == k {
				resultMap[key] = v / value
			}
		}
	}
	result = SortMapFromHighToLowByValue(resultMap)
	for i := 0; i < 5; i++ {
		fmt.Printf("%s - %d\n", result[i].Key, result[i].Value)
	}
}

func CountDeveloperByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	job := "developer"
	for _, person := range p {
		if person.Job == job {
			result[person.City]++
		}
	}
	return result
}

func SalaryDeveloperByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	job := "developer"
	for _, person := range p {
		if person.Job == job {
			result[person.City] += person.Salary
		}
	}
	return result
}

func FiveCitiesHasTopAverageSalaryForDeveloper(p []Person) {
	resultMap := make(map[string]int)
	var result []kv
	developerByCity := CountDeveloperByCity(p)
	salaryDeveloperEachCity := SalaryDeveloperByCity(p)
	for key, value := range developerByCity {
		for k, v := range salaryDeveloperEachCity {
			if key == k {
				resultMap[key] = v / value
			}
		}
	}
	result = SortMapFromHighToLowByValue(resultMap)
	for i := 0; i < 5; i++ {
		fmt.Printf("%s - %d\n", result[i].Key, result[i].Value)
	}
}

func CalculateAge(birthDate string) (result int) {
	now := time.Now()
	ny := now.Year()
	nm := int(now.Month())
	nd := now.Day()
	birthDateSplit := strings.Split(birthDate, "-")
	tmp := make([]int, 0)
	for _, value := range birthDateSplit {
		if x, err := strconv.Atoi(value); err == nil {
			tmp = append(tmp, x)
		} else {
			fmt.Println(err)
		}
	}
	var age int
	if tmp[1] > nm {
		age = ny - tmp[0]
	}
	if tmp[1] == nm {
		if tmp[2] >= nd {
			age = ny - tmp[0]
		} else {
			age = ny - tmp[0] - 1
		}
	} else if tmp[1] < nm {
		age = ny - tmp[0] - 1
	}
	return age
}

func TotalAgeInEachJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job] += CalculateAge(person.Birthday)
	}
	return result
}

func AverageAgePerJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	groupJob := GroupPeopleByJob(p)
	ageEachJob := TotalAgeInEachJob(p)
	for key, value := range groupJob {
		for k, v := range ageEachJob {
			if key == k {
				result[key] = v / value
			}
		}
	}
	return result
}

func TotalAgeInEachCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City] += CalculateAge(person.Birthday)
	}
	return result
}

func AverageAgePerCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	groupCity := CountPeopleByCity(p)
	ageEachCity := TotalAgeInEachCity(p)
	for key, value := range groupCity {
		for k, v := range ageEachCity {
			if key == k {
				result[key] = v / value
			}
		}
	}
	return result
}
