package jedlik

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() Car {
    if c.battery < c.batteryDrain {
        return *c
    }
	c.distance = c.distance + c.speed
	c.battery = c.battery - c.batteryDrain
    
	return *c
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
	steps := (trackDistance + c.speed - 1) / c.speed
	battery := c.battery
	for i := 0; i < steps; i++ {
		battery -= c.batteryDrain
		if battery < 0 {
			return false
		}
	}
	return true
}
