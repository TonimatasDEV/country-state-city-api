[![Hits](https://hits.sh/github.com/tonimatasdev/country-state-city-api.svg?view=today-total&label=Views)](https://hits.sh/github.com/tonimatasdev/country-state-city-api/)

# Country State City API

With this API you can get countries, states and cities around the world.

## Routes
- https://country-state-city.net/countries
- https://country-state-city.net/states/:country
- https://country-state-city.net/cities/:country/:state

The params :country and :state are Iso2.

## Build
```shell
export GIN_MODE=release
go build -o dist/CountryStateCityApi ./cmd
```