package domain

type State struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Iso2      string `json:"iso2"`
	Iso31662  string `json:"iso3166_2"`
	Native    string `json:"native"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Type      string `json:"type"`
	Timezone  string `json:"timezone"`
	Cities    []City `json:"cities"`
}
