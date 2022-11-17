# Study Event-Driven Architecture with Golang

- Go 1.15
- Machinery v1

<br/>

## Run Server
### With a local binary
#### Build
```bash
# 1. run redis
$ docker run -p 6379:6379 redis 

# 2. compile the project
$ make all
```

#### For MacOS
```bash
$ ./bin/study-websocket-go
```
#### For Windows
```bash
$ ./bin/study-event-go.exe
```
#### For VSCode
- Click the button 'Launch'

<br/>

## Send a task
- You need to make data of Garden and Lily in the DB

### With cURL
```bash
$ curl --location --request POST '0.0.0.0:4567/api/v1/assault-lily/gardens/1/alarm' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cave_location": "shibuyaku",
    "huges": [
        {
            "huge_class": "gigant",
            "huge_type": "special"
        },
        {
            "huge_class": "small",
            "huge_type": "none"
        }
    ],
    "total_huge_count": 12,
    "alert_level": 3,
    "legion_member_count": 4
}'
```

### Result on the console
```bash
[ANNOUNCEMENT] LEGION IS CREATED!!!
[ANNOUNCEMENT] THE MEMBER IS ...
[ANNOUNCEMENT] 0) sachie jeanne fukuyama
[ANNOUNCEMENT] 1) yuria francisca kuroki
[ANNOUNCEMENT] 2) seren sophia amamiya
[ANNOUNCEMENT] 3) raimu lucia kishimoto
[ANNOUNCEMENT] GO TO shibuyaku NOW IF YOU ARE A MEMBER OF THE LEGION!
```
