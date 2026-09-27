package memory

type AvailabilityChecker struct {
}

func (a *AvailabilityChecker) IsAvailable(product string) (bool, error) {
	if product == "unavailable" {
		return false, nil
	}
	return true, nil
}
