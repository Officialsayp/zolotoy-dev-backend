package memory

type AvailabilityChecker struct {
}

func (a *AvailabilityChecker) IsAvailable(productID string, _ int64) (bool, error) {
	if productID == "unavailable" {
		return false, nil
	}
	return true, nil
}
