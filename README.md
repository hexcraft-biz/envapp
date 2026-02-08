# envapp

## usage
### ENV variable:
```
APP_ENV=development|debug|stage|production
APP_HOST=my.example.app
APP_PATH=/path
APP_PORT=8080
TIMEZONE=
TRUST_PROXY=localhost
VISIBILITY=internal|external
```

### secrets JSON file:
```json
{
  "env": "development|debug|stage|production",
  "host": "my.example.app",
  "path": "/path",
  "port": "8080",
  "timezone": "",
  "trustProxy": "localhost",
  "visibility": "internal|external"
}
```
