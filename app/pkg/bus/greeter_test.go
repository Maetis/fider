package bus_test

import (
	"context"

	"github.com/getfider/fider/app/pkg/bus"
)

type SayHelloCommand struct {
	Name   string
	Result string
}

var GreetingKey = "GreetingKey"

type GreeterService struct {
}

func (s GreeterService) Category() string {
	return "greeter"
}

func (s GreeterService) Enabled() bool {
	return true
}

func (s GreeterService) Init() {
	bus.AddHandler(s, Greet)
}

var getGreeting = func(ctx context.Context) string {
	return ctx.Value(GreetingKey).(string)
}

func Greet(ctx context.Context, cmd *SayHelloCommand) error {
	cmd.Result = getGreeting(ctx) + " " + cmd.Name
	return nil
}

type BetterGreeterService struct {
}

func (s BetterGreeterService) Category() string {
	return "greeter"
}

func (s BetterGreeterService) Enabled() bool {
	return true
}

func (s BetterGreeterService) Init() {
	bus.AddHandler(s, SayHello)
}

func SayHello(ctx context.Context, cmd *SayHelloCommand) error {
	cmd.Result = "Hello " + cmd.Name
	return nil
}
