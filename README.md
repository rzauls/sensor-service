# Go + Sqlite HTTP api example

This project was written for educational purposes

## Build/Run instructions

### with Docker

Requirements:
 - working `docker` installation on any platform

1. Clone this repository `git clone https://github.com/rzauls/sensor-service`
2. Browse to repository `cd sensor-service`
3. Build docker image `docker build -t sensor-service .`
4. Run docker container `docker run -it -p 8080:8080 sensor-service`
5. Access API @ `http://localhost:8080`
    
### without Docker

Requirements:
 - working Go 1.14 installation
 - working `gcc` installation (for sqlite adapter compilation)

1. Clone this repository `git clone https://github.com/rzauls/sensor-service`
2. Browse to repository `cd sensor-service`
3. Run application `go run .`
5. Access API @ `http://localhost:8080`
 
 ## API description
 
 # Show all latest sensor measurements
 
 Get the details of all sensors in the database, and fetch the most recent measurement of each metric.
 
 **URL** : `/sensors/`
 
 **Method** : `GET`
 
 ## Success Response
 
 **Code** : `200 OK`
 
 **Content examples**
 
 ```json
 [
   {
     "sensor_id": 34,
     "name": "Hall",
     "serial_code": "1048827",
     "data": [
       {
         "metric_name": "Humidity",
         "rvalue": 50,
         "unit_name": "%",
         "precision": 0,
         "rtime": "2019-08-21 00:56:37"
       },
       {
         "metric_name": "Temperature",
         "rvalue": 25.05,
         "unit_name": "°C",
         "precision": 1,
         "rtime": "2019-08-21 00:26:38"
       }
     ]
   }
   ...
 ```

# Show all sensor min/max measurements on specific date
 
 Get the details of all sensors in the database, and fetch the minimal and maximal reading on a specific date.
 The date format should be `YYYY-MM-DD`, `YYYY-MM` or `YYYY`.

 
 **URL** : `/sensors/{date}`
 
 **Method** : `GET`

 ## Success Response
 
 **Code** : `200 OK`
 
 **Content examples**
 
 Query: `/sensors/2019-07-01`
 ```json
 [
   {
     "sensor_id": 22,
     "name": "Classroom",
     "serial_code": "1048836",
     "data": [
       {
         "metric_name": "Humidity",
         "unit_name": "%",
         "precision": 0,
         "rvalue_min": 43,
         "rvalue_max": 51,
         "date": "2019-07-01"
       },
       {
         "metric_name": "Temperature",
         "unit_name": "°C",
         "precision": 1,
         "rvalue_min": 22.549999237060547,
         "rvalue_max": 25.350000381469727,
         "date": "2019-07-01"
       }
     ]
   },
  ...
 ]
 ```


