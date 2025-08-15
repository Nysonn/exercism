package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {

	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}

}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	//if the cars battery is less than the battery drain percentage so that the car cannot move
	if car.battery < car.batteryDrain {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain

	return car

}

// CanFinish determines if a car can finish a race track without running out of battery
func CanFinish(car Car, track Track) bool {
	// Calculate how many drives are needed to finish the track
	// We need to cover at least the track distance
	remainingDistance := track.distance - car.distance

	// If we're already at or past the finish line, we can finish
	if remainingDistance <= 0 {
		return true
	}

	// Calculate how many drives we need
	drivesNeeded := (remainingDistance + car.speed - 1) / car.speed // This is ceiling division

	// Calculate how much battery we need for those drives
	batteryNeeded := drivesNeeded * car.batteryDrain

	// Check if we have enough battery
	return car.battery >= batteryNeeded
}
