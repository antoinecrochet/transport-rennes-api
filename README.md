# Transport Rennes Backend

Back-end part of Alexa application for public transport at Rennes developped in Go.

# Applications

## tr-lambda

AWS lambda for Alexa application.
The project uses :
* AWS Lambda SDK (https://github.com/aws/aws-lambda-go)
* alexa-sdk-go (https://github.com/dasjott/alexa-sdk-go)

### Build

In PowerShell, run the following:

```sh
$env:GOOS = "linux"
$env:CGO_ENABLED = "0"
$env:GOARCH = "amd64"
go build -o main .\cmd\tr-lambda\main.go
~\Go\Bin\build-lambda-zip.exe -output main.zip main
```

Have a look at https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html for more information.

## tr-server

Restful server giving endpoint to get the upcomping buses according to:
* the bus name (C1, C2...)
* the bus stop (Metz Volney, République...)
* the final destination (Chantepie, La Poterie...)

Only the bus stop is mandatory.

### API

#### Upcoming bus

Request `HTTP GET /upcomingbus`

* Input example
```json
{
   "busline": "C1",
   "stop": "Metz Volney",
   "destination": "Chantepie"
}
```

* Output
```json
{
   "message": "Prochain bus dans 29 min, le suivant dans 37 min"
}
```

### Configuration
Generate a config.json file next to the executable using the template (config.json.dist):

```json
{
   "base_url": "https://data.explore.star.fr"
}
```

To generate your api key, follow the instructions here https://help.opendatasoft.com/apis/ods-search-v1/#finding-and-generating-api-keys using the opendatasoft of Star Rennes https://data.explore.star.fr
