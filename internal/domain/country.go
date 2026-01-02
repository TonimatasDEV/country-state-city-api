package domain

type Country struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Iso3           string       `json:"iso3"`
	Iso2           string       `json:"iso2"`
	NumericCode    string       `json:"numeric_code"`
	Phonecode      string       `json:"phonecode"`
	Capital        string       `json:"capital"`
	Currency       string       `json:"currency"`
	CurrencyName   string       `json:"currency_name"`
	CurrencySymbol string       `json:"currency_symbol"`
	Tld            string       `json:"tld"`
	Native         string       `json:"native"`
	Population     int          `json:"population"`
	Gdp            interface{}  `json:"gdp"`
	Region         string       `json:"region"`
	RegionID       int          `json:"region_id"`
	Subregion      string       `json:"subregion"`
	SubregionID    int          `json:"subregion_id"`
	Nationality    string       `json:"nationality"`
	Timezones      []Timezones  `json:"timezones"`
	Translations   Translations `json:"translations"`
	Latitude       string       `json:"latitude"`
	Longitude      string       `json:"longitude"`
	Emoji          string       `json:"emoji"`
	EmojiU         string       `json:"emojiU"`
	States         []States     `json:"states"`
}
