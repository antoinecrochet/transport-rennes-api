# Transport Rennes API

This project is a web application built with Go to retrieve upcoming buses from public transport in Rennes.
It provides an opendatasoft client to retrieve data on public transport in Rennes (used in an alexa skill cf. transport-rennes-alexa repository).

# Application

## tr-server

Web application exposing an api to search for the upcomping buses according to:
* the bus name (C1, C2...)
* the bus stop (Metz Volney, République...)
* the final destination (Chantepie, La Poterie...)

Only the bus stop is mandatory.

### API

#### Search upcoming bus

POST `/search/upcomingbus`

Request body:
```json
{
   "busline": "C1",
   "stop": "Metz Volney",
   "destination": "Chantepie"
}
```

Response body:
```json
{
    "message": "Prochain bus dans 5 min, le suivant dans 6 min",
    "totalCount": 14,
    "hits": [
        {
            "busline": "C4",
            "stop": "Beaulieu Chimie",
            "departure": "2025-03-11T18:31:01Z",
            "destination": "Grand Quartier"
        },
        {
            "busline": "C4",
            "stop": "Beaulieu Chimie",
            "departure": "2025-03-11T18:31:27Z",
            "destination": "ZA Saint-Sulpice"
        },
        {
            "busline": "C4",
            "stop": "Beaulieu Chimie",
            "departure": "2025-03-11T18:39:27Z",
            "destination": "ZA Saint-Sulpice"
        },
        ...
    ]
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
