package vistor

import "testing"

func TestVistor(t *testing.T) {
	plan := NewTravelPlan()
	plan.AddArea(NewGuangZhou())
	plan.AddArea(NewShanghai())
	plan.DoTravel(NewFlighChina())
	plan.AddArea(NewTokyo())
	plan.DoTravel(NewFlighJapan())
}