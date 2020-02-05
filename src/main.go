package main

import "fmt"

type Person struct {
  age int
}

// checks if there is a person who is exactly twice as old as another in the list
func getTwiceAsOld(people []Person) bool {
  for i:=0; i < len(people); i ++ { 
    if checksTwiceAgeExists(people[i].age / 2, people) {
      return true
    }
  }
  return false
}

func checksTwiceAgeExists(age int, people []Person) bool {
	for _, b := range people {
		if b.age == age {
			return true
		}
	}
	return false
}

//checks if there is a person who is at least twice as old as another person in the list 
func getAtLeastTwiceAsOld(people []Person) bool {
  for i:=0; i < len(people); i ++ { 
    if checksAtLeastTwiceAgeExists(people[i].age, people) {
      return true
    }
  }
  return false
}

func checksAtLeastTwiceAgeExists(age int, people []Person) bool {
	for _, b := range people {
		if b.age * 2 <= age {
			return true
		}
	}
	return false
}

func main() {
	people := []Person{
		{
			age: 5,
		},
		{
			age: 10,
		},
		{
			age: 16,
		},
		
	}
	fmt.Println(getTwiceAsOld(people))
	fmt.Println(getAtLeastTwiceAsOld(people))
}