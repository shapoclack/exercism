package airportrobot

// Write your code here.

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}
type Italian struct{}
type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func (i Italian) LanguageName() string {
	return "Italian"
}

func (p Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}
func (i Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}
func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
