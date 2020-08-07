package ramerica

// Record holds one campsite record// Record is a specific Campsite Record
type Record struct {
	Summary         interface{} `json:"summary"`
	MatchingFilters bool        `json:"matchingFilters"`
	NamingLabel     string      `json:"namingLabel"`
	Available       bool        `json:"available"`
	ReservableType  interface{} `json:"reservableType"`
	Type            string      `json:"type"`
	LogoSrc         interface{} `json:"logoSrc"`
	Proximity       float64     `json:"proximity"`
	Name            string      `json:"name"`
	TargetItem      bool        `json:"targetItem"`
	NamingID        string      `json:"namingId"`
	Details         struct {
		LineOfBusinesses []string `json:"lineOfBusinesses"`
		LegacyRequired   bool     `json:"legacyRequired"`
		OfferingSupport  []struct {
			Business string `json:"business"`
			Type     string `json:"type"`
			URL      string `json:"url"`
		} `json:"offeringSupport"`
		Rating       interface{} `json:"rating"`
		Description  interface{} `json:"description"`
		Availability struct {
			Available      bool   `json:"available"`
			ReservableType string `json:"reservableType"`
		} `json:"availability"`
		NonClientFacility bool `json:"nonClientFacility"`
		FavoriteSupport   bool `json:"favoriteSupport"`
		Cta               struct {
			Caption string      `json:"caption"`
			Type    interface{} `json:"type"`
			URL     interface{} `json:"url"`
		} `json:"cta"`
		ContrCode   string `json:"contrCode"`
		Affiliation struct {
			FacilityID int    `json:"facilityID"`
			Name       string `json:"name"`
			Agencies   []struct {
				Code string      `json:"code"`
				Name string      `json:"name"`
				Logo interface{} `json:"logo"`
				ID   int         `json:"id"`
				URL  interface{} `json:"url"`
			} `json:"agencies"`
			URL interface{} `json:"url"`
		} `json:"affiliation"`
		ImageURL               string      `json:"imageURL"`
		ID                     int         `json:"id"`
		VerifiableAvailability bool        `json:"verifiableAvailability"`
		FlexResult             interface{} `json:"flexResult"`
		FacilityTypes          []string    `json:"facilityTypes"`
		ActvAdvInd             interface{} `json:"actvAdvInd"`
		CrossOverCustomLanding bool        `json:"crossOverCustomLanding"`
		RatingSupport          bool        `json:"ratingSupport"`
		WebAddress             interface{} `json:"webAddress"`
		DayUseOnly             bool        `json:"dayUseOnly"`
		Coordinates            struct {
			LatitudeDec        interface{} `json:"latitudeDec"`
			LongitudeDegMinSec interface{} `json:"longitudeDegMinSec"`
			LongitudeDec       interface{} `json:"longitudeDec"`
			Latitude           float64     `json:"latitude"`
			LatitudeDegMinSec  interface{} `json:"latitudeDegMinSec"`
			Longitude          float64     `json:"longitude"`
		} `json:"coordinates"`
		BaseURL     string   `json:"baseURL"`
		StateCodes  []string `json:"stateCodes"`
		Name        string   `json:"name"`
		DayUse      bool     `json:"dayUse"`
		HiddenOnMap bool     `json:"hiddenOnMap"`
		Favorite    bool     `json:"favorite"`
	} `json:"details"`
	ID int `json:"id"`
}

type SearchPage struct {
	Name string
}
