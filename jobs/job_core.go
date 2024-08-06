package jobs

var jobRates = map[string]float64{
	"hotel_shift":          25.0,
	"pet_sitting":          17.0,
	"cat_visit":            25.0,
	"overnight_hotel":      80.0,
	"overnight_petsitting": 120.0,
	"cat_at_sitter_home":   25.0,
	"dog_at_sitter_home":   75.0,
}

func GetSupportedJobTypes() []string {
	supportedJobTypes := []string{
		"Hotel Shift",
		"Pet Sitting",
		"Cat Visit",
		"Overnight Hotel Shift",
		"Overnight Pet Sitting",
		"Cat at Sitter Home",
		"Dog at Sitter Home",
		"Uber / Expense",
	}

	return supportedJobTypes
}
