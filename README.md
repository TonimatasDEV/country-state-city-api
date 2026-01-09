[![Hits](https://hits.sh/github.com/tonimatasdev/country-state-city-api.svg?view=today-total&label=Views)](https://hits.sh/github.com/tonimatasdev/country-state-city-api/)

# Country State City API

With this API you can get countries, states and cities around the world.

## Routes
- https://country-state-city.net/country/names
- https://country-state-city.net/country/nativenames
- https://country-state-city.net/country/iso2
- https://country-state-city.net/country/iso3
- https://country-state-city.net/state/:country/names
- https://country-state-city.net/city/:country/:state/names

## Build
```shell
go build -o dist/CountryStateCityApi ./cmd
```