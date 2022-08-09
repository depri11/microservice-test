package entity

type Logistic struct {
	LogisticName    string
	Amount          int64
	DestinationName string
	OriginName      string
	Duration        string
}

type Logistics []Logistic
