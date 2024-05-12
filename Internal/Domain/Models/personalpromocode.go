package Models

type PersonalPromoCode struct {
	Id               int32
	IdClient         int32
	IdGroup          int32
	NamePromoCode    string
	TypeDiscount     int32
	ValueDiscount    int32
	DateStartActive  string
	DateFinishActive string
}
