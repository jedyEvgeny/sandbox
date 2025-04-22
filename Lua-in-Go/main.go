//Практика интеграции Луа в Голанг
// Описание процесса в публикации https://dzen.ru/a/aAIvrGF_NSuF9PSP

package main

import (
	"fmt"
	"log"

	lua "github.com/yuin/gopher-lua"
)

type Person struct {
	Name     string
	Age      int
	Job      string
	Address  Address
	IsSenior bool
}

type Address struct {
	Street string
	City   string
}

const scriptLua = `
age = person:findAge()
print("Age:", age)

if age > 18 then
  person:setSeniorStatus()
end

city = person:findAdress():findCity()
print("City:", city)
`

func main() {
	person := newPerson()

	envLua := lua.NewState()
	defer envLua.Close()

	addressLua := envLua.NewTable()
	personLua := envLua.NewTable()

	envLua.SetField(personLua, "Adress", addressLua)
	envLua.SetGlobal("person", personLua)

	envLua.SetField(personLua, "findAdress", envLua.NewFunction(func(L *lua.LState) int {
		address := person.findAdress()
		envLua.SetField(addressLua, "Street", lua.LString(address.Street))
		envLua.SetField(addressLua, "City", lua.LString(address.City))
		L.Push(addressLua)

		return 1
	}))

	envLua.SetField(addressLua, "findCity", envLua.NewFunction(func(L *lua.LState) int {
		address := person.Address.findCity()
		L.Push(lua.LString(address))

		return 1
	}))

	envLua.SetField(personLua, "setSeniorStatus", envLua.NewFunction(func(L *lua.LState) int {
		person.setSeniorStatus()

		return 0
	}))

	envLua.SetField(personLua, "findAge", envLua.NewFunction(func(L *lua.LState) int {
		age := person.findAge()
		L.Push(lua.LNumber(age))

		return 1
	}))

	if err := envLua.DoString(scriptLua); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", *person)
}

func newPerson() *Person {
	return &Person{
		Name: "John Doe",
		Age:  25,
		Job:  "N/A",
		Address: Address{
			Street: "Baker",
			City:   "London",
		},
	}
}

func (p *Person) findAdress() *Address {
	if p != nil {
		return &p.Address
	}

	return nil
}

func (a *Address) findCity() string {
	if a != nil {
		return a.City
	}

	return ""
}

func (p *Person) setSeniorStatus() {
	if p != nil {
		p.IsSenior = true
	}
}

func (p *Person) findAge() int {
	if p != nil {
		return p.Age
	}

	return 0
}
