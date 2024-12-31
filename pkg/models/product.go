package models

type Product struct {
	Title     string
	Image_Url string
	Url       string
	Price     int
}

type Product_schema struct {
	Context         string
	Type            string
	Name            string
	Image           any
	AggregateRating any
	Offers          Offer
}

type AggregateRating struct {
	RatingValue any
	ReviewCount any
}

type Offer struct {
	Type          string
	Price         any
	PriceCurrency any
	Availability  string
	URL           string
}
