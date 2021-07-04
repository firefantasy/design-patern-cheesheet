package vistor

import "log"

type IArea interface {
	Accept(IFlight)
}

type IFlight interface {
	Visit(IArea)
}

type Shanghai struct {
}

func NewShanghai () *Shanghai {
	return &Shanghai{}
}

func (s *Shanghai) Accept(i IFlight) {
	i.Visit(s)
}

type GuangZhou struct {
}

func NewGuangZhou () *GuangZhou {
	return &GuangZhou{}
}

func (s *GuangZhou) Accept(i IFlight) {
	i.Visit(s)
}

type flighChina struct {
}

func NewFlighChina () flighChina {
	return flighChina{}
}

func (v flighChina) Visit(i IArea) {
	switch (i).(type) {
	case *Shanghai:
		log.Println("travel to China to Visit Shanghai")
	case *GuangZhou:
		log.Println("travel to China to Visit GuangZhou")
	}
	
}

type flighJapan struct {}

func NewFlighJapan () flighJapan{
	return flighJapan{}
}

type Tokyo struct {
}

func NewTokyo () *Tokyo {
	return &Tokyo{} 
}

func (s *Tokyo) Accept(i IFlight) {
	i.Visit(s)
}

func (v flighJapan) Visit(i IArea) {
	switch (i).(type) {
	case *Tokyo:
		log.Println("travel to Japan to Visit Tokyo")
	}
}

type TravelPlan struct{
	areas []IArea
}

func NewTravelPlan() *TravelPlan{
	return &TravelPlan{}
}

func (t *TravelPlan) AddArea(i IArea) {
	t.areas = append(t.areas, i)
}

func (t *TravelPlan) DoTravel(i IFlight) {
	for _, a := range t.areas {
		a.Accept(i)
	}
	t.areas = nil
}