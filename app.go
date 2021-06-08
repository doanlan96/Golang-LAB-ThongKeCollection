package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
"name":"Dee Leng",
"email":"dleng0@cocolog-nifty.com",
"job":"developer",
"gender":"Female",
"city":"London",
"salary":9662,
"birthdate":"2007-09-30" */
type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("person.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var people []Person

	json.Unmarshal(byteValue, &people)

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(&people[i])
		}
	*/

	//Câu 1
	fmt.Println("--------------")
	fmt.Println("Group people by city: ")
	peopleByCity := GroupPeopleByCity(people)
	for key, value := range peopleByCity {
		fmt.Println(key)
		for _, person := range value {
			fmt.Println("  ", (&person).Name)
		}
	}

	//Câu 2
	fmt.Println("--------------")
	fmt.Println("Group people by job: ")
	peopleByJob := GroupPeopleByJob(people)
	for k, v := range peopleByJob {
		fmt.Println(k, "-", v)
	}

	//Câu 3
	fmt.Println("----------------")
	fmt.Println("Top five jobs by number: ")
	Top5JobsByNumber(people)

	//Câu 4
	fmt.Println("----------------")
	fmt.Println("Top five cities by number: ")
	Top5CitiesByNumber(people)

	//Câu 5
	fmt.Println("----------------")
	fmt.Println("Top job in each city: ")
	topJobInCity := TopJobByNumerInEachCity(people)
	for key, value := range topJobInCity {
		fmt.Println(key)
		for k, v := range value {
			fmt.Println(" ", k, v)
		}
	}

	//Câu 6
	fmt.Println("----------------")
	fmt.Println("Average salary by job: ")
	salaryByJob := AverageSalaryByJob(people)
	for k, v := range salaryByJob {
		fmt.Println(k, "-", v)
	}

	//Câu 7
	fmt.Println("---------------")
	fmt.Println("Top Five Cities Has Top Average Salary: ")
	FiveCitiesHasTopAverageSalary(people)

	//Câu 8
	fmt.Println("---------------")
	fmt.Println("Top Five Cities Has Top Average Salary For Developer: ")
	FiveCitiesHasTopAverageSalaryForDeveloper(people)

	//Câu 9
	fmt.Println("----------------")
	fmt.Println("Average age by job: ")
	ageByJob := AverageAgePerJob(people)
	for k, v := range ageByJob {
		fmt.Println(k, "-", v)
	}

	//Câu 10
	fmt.Println("----------------")
	fmt.Println("Average age by city: ")
	ageByCity := AverageAgePerCity(people)
	for k, v := range ageByCity {
		fmt.Println(k, "-", v)
	}
}
