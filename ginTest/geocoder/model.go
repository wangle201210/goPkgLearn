package geocoder

type Hospital struct {
	Id       int32   `json:"id"`
	Province string  `json:"province"`
	City     string  `json:"city"`
	Country  string  `json:"country"`
	Address  string  `json:"address"`
	Lng      float32 `json:"lng"`
	Lat      float32 `json:"lat"`
}

type Geocoder struct {
	Message string `json:"message"`
	Result  struct {
		AdInfo struct {
			Adcode string `json:"adcode"`
		} `json:"ad_info"`
		AddressComponents struct {
			City         string `json:"city"`
			District     string `json:"district"`
			Province     string `json:"province"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_components"`
		Deviation int64 `json:"deviation"`
		Level     int64 `json:"level"`
		Location  struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Reliability int64   `json:"reliability"`
		Similarity  float64 `json:"similarity"`
		Title       string  `json:"title"`
	} `json:"result"`
	Status int64 `json:"status"`
}

func (Hospital) TableName() string {
	return "hospital"
}
