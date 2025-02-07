# Transport Rennes API

This project is a web application built with Go to retrieve upcoming buses from public transport in Rennes.
It provides an opendatasoft client to retrieve data on public transport in Rennes (used in an alexa skill cf. transport-rennes-alexa repository).

# Application

## tr-server

Web application exposing an api to get the upcomping buses according to:
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
