package main

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
)

func menghasilkan(arg1 string) error {
	if Check() != arg1 {
		return fmt.Errorf("It Should Be %s", Check())
	}
	return nil
}

func sayaMemilikiNilaiSisi(arg1, arg2, arg3 int) error {
	a = arg1
	b = arg2
	c = arg3
	return nil

}

func sayaMendeteksiSegitiga() error {
	if a <= 0 || b <= 0 || c <= 0 {
		return fmt.Errorf("This is not triangle")
	}
	return nil
}

func InitializeTestSuite(sc *godog.TestSuiteContext) {
	sc.BeforeSuite(func() {
		a = 7
		b = 7
		c = 7
	})
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		a = 7
		b = 7
		c = 7 // clean the state before every scenario

		return ctx, nil
	})
	sc.Step(`^Menghasilkan "([^"]*)"$`, menghasilkan)
	sc.Step(`^Saya memiliki nilai sisi (\d+), (\d+), (\d+)$`, sayaMemilikiNilaiSisi)
	sc.Step(`^Saya mendeteksi segitiga$`, sayaMendeteksiSegitiga)
}
