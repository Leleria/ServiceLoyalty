package Models

type PromoCode struct {
	Id               int32
	Name             string
	TypeDiscount     int32
	ValueDiscount    int32
	DateStartActive  string
	DateFinishActive string
	MaxCountUses     int32
}
