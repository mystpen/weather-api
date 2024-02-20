# Weather API
<b>REST API service in GO that interacts with weather data</b>

## Usage
- <b>To update or create data:</b>
```
PUT /weather?location=London
```
- <b>To get current weather:</b>
```
GET /weather?location=London
```
```
{
    "Location": "London",
    "Temperature": 11,
    "Description": "overcast clouds"
}
```

## Run
- Clone repository:
```
git clone git@github.com:mystpen/weather-api.git
```
- Go to ```weather-api``` directory
- Create ```config.yaml``` file with [OpenWeatherMap](https://openweathermap.org/) api-token and mongodb database name:
```
token:
db_name: 
```
- To run project:
```
make run
```
```
go run ./cmd/app/ .
```