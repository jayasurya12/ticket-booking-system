package model

type BookingRequest struct {
	Token         string `json:"token"`
	TrainID       string `json:"trainId"`
	PassengerName string `json:"passengerName"`
}
