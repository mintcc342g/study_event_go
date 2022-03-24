# Study Event-Driven Architecture with Golang

- Go 1.15
- Machinery V1

<br/>

## Run Server

- for MacOS and Windows

```bash
# 1. run redis
$ docker run -p 6379:6379 redis 

# 2. compile the project
$ make all


# 3-1. run the server for MacOS
$ ./bin/study-websocket-go

# 3-2. run the server for Windows
$ ./bin/study-event-go.exe

# 3-3. You can use the debug mod of VSCode to click the button 'Launch'
```

<br/>

## Event Test

- You need to make data of Garden and Lily
- cURL example
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

- result on console
```bash
[ANNOUNCEMENT] LEGION IS CREATED!!!
[ANNOUNCEMENT] THE MEMBER IS ...
[ANNOUNCEMENT] 0) sachie jeanne fukuyama
[ANNOUNCEMENT] 1) yuria francisca kuroki
[ANNOUNCEMENT] 2) seren sophia amamiya
[ANNOUNCEMENT] 3) raimu lucia kishimoto
[ANNOUNCEMENT] GO TO shibuyaku NOW IF YOU ARE A MEMBER OF THE LEGION!
```