package ramerica

// SearchPage holds all the reponse for a search resul curl
type SearchPage2 struct {
	TotalRecords int `json:"totalRecords"`
	TotalPages   int `json:"totalPages"`
	StartIndex   int `json:"startIndex"`
	EndIndex     int `json:"endIndex"`
	Control      struct {
		CurrentPage int `json:"currentPage"`
		PageSize    int `json:"pageSize"`
	} `json:"control"`
	Records      []Record `json:"records"`
	TypesSummary struct {
		ContractTypeFEDERAL  int `json:"ContractType[FEDERAL]"`
		ContractTypeSTATE    int `json:"ContractType[STATE]"`
		ContractTypeREGIONAL int `json:"ContractType[REGIONAL]"`
		ContractTypeCOUNTY   int `json:"ContractType[COUNTY]"`
		ContractTypePRIVATE  int `json:"ContractType[PRIVATE]"`
		ContractTypeOTHER    int `json:"ContractType[OTHER]"`
	} `json:"typesSummary"`
	ActvAdvInfo interface{} `json:"actvAdvInfo"`
}

// Record holds one campsite record// Record is a specific Campsite Record
type Record struct {
	Type            string      `json:"type"`
	NamingID        string      `json:"namingId"`
	NamingLabel     string      `json:"namingLabel"`
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	TargetItem      bool        `json:"targetItem"`
	MatchingFilters bool        `json:"matchingFilters"`
	ReservableType  interface{} `json:"reservableType"`
	Available       bool        `json:"available"`
	Proximity       float64     `json:"proximity"`
	Summary         interface{} `json:"summary"`
	Details         struct {
		ID                     int      `json:"id"`
		ContrCode              string   `json:"contrCode"`
		StateCodes             []string `json:"stateCodes"`
		Name                   string   `json:"name"`
		BaseURL                string   `json:"baseURL"`
		ImageURL               string   `json:"imageURL"`
		VerifiableAvailability bool     `json:"verifiableAvailability"`
		Availability           struct {
			Available      bool   `json:"available"`
			ReservableType string `json:"reservableType"`
		} `json:"availability"`
		Affiliation struct {
			FacilityID int         `json:"facilityID"`
			Name       string      `json:"name"`
			URL        interface{} `json:"url"`
			Agencies   []struct {
				ID   int         `json:"id"`
				Code string      `json:"code"`
				Name string      `json:"name"`
				Logo interface{} `json:"logo"`
				URL  interface{} `json:"url"`
			} `json:"agencies"`
		} `json:"affiliation"`
		FacilityTypes []string    `json:"facilityTypes"`
		Description   interface{} `json:"description"`
		FlexResult    interface{} `json:"flexResult"`
		Cta           struct {
			Caption string      `json:"caption"`
			URL     interface{} `json:"url"`
			Type    interface{} `json:"type"`
		} `json:"cta"`
		DayUse          bool        `json:"dayUse"`
		DayUseOnly      bool        `json:"dayUseOnly"`
		RatingSupport   bool        `json:"ratingSupport"`
		Rating          interface{} `json:"rating"`
		ActvAdvInd      interface{} `json:"actvAdvInd"`
		FavoriteSupport bool        `json:"favoriteSupport"`
		Favorite        bool        `json:"favorite"`
		WebAddress      interface{} `json:"webAddress"`
		Coordinates     struct {
			Latitude           float64     `json:"latitude"`
			LatitudeDec        interface{} `json:"latitudeDec"`
			LatitudeDegMinSec  interface{} `json:"latitudeDegMinSec"`
			Longitude          float64     `json:"longitude"`
			LongitudeDec       interface{} `json:"longitudeDec"`
			LongitudeDegMinSec interface{} `json:"longitudeDegMinSec"`
		} `json:"coordinates"`
		HiddenOnMap            bool     `json:"hiddenOnMap"`
		CrossOverCustomLanding bool     `json:"crossOverCustomLanding"`
		LegacyRequired         bool     `json:"legacyRequired"`
		LineOfBusinesses       []string `json:"lineOfBusinesses"`
		OfferingSupport        []struct {
			Business string `json:"business"`
			Type     string `json:"type"`
			URL      string `json:"url"`
		} `json:"offeringSupport"`
		NonClientFacility bool `json:"nonClientFacility"`
	} `json:"details"`
	LogoSrc interface{} `json:"logoSrc"`
}
type SearchPage struct {
	TotalRecords int `json:"totalRecords"`
	TotalPages   int `json:"totalPages"`
	StartIndex   int `json:"startIndex"`
	EndIndex     int `json:"endIndex"`
	Control      struct {
		CurrentPage int `json:"currentPage"`
		PageSize    int `json:"pageSize"`
	} `json:"control"`
	Records      []Record `json:"records"`
	TypesSummary struct {
		ContractTypeFEDERAL  int `json:"ContractType[FEDERAL]"`
		ContractTypeSTATE    int `json:"ContractType[STATE]"`
		ContractTypeREGIONAL int `json:"ContractType[REGIONAL]"`
		ContractTypeCOUNTY   int `json:"ContractType[COUNTY]"`
		ContractTypePRIVATE  int `json:"ContractType[PRIVATE]"`
		ContractTypeOTHER    int `json:"ContractType[OTHER]"`
	} `json:"typesSummary"`
	ActvAdvInfo interface{} `json:"actvAdvInfo"`
}
