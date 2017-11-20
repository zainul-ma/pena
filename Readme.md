# Pena

Pena is a package that can help circuit breaker to log the status. The status that can log by pena is :
 - Tripped
 - Failed
 - Closed

## Quick Example Usage

```
package main

import (
    ....

    "github.com/zainul-ma/pena"
    
    ....
)

var (
    ....

    circuitStatus pena.CircuitStatus
    
    ....
)

func init() {
    ....

    pena.SetDB("mongodb://localhost:27017", "your_db_service")
    circuitStatus.Dial("localhost:6379", "your_db_service")

    ....


    go circuitStatus.SetFail("host_destination_trip", "routing_destination_trip")

    go pena.WriteLog(pena.CircuitLog{
        App:       "your_app_name",
        ErrorCode: "your_error_code",
        Fail:      true,
        Host:      "host_destination_trip",
        Route:     "routing_destination_trip",
        Tripped:   false,
    })
}

```

Initiate circuit breaker status 

```
var circuitStatus pena.CircuitStatus
```

Set the DB or source in mongoDB & redis with credential

```
    pena.SetDB("mongodb://localhost:27017", "your_db_service")
    circuitStatus.Dial("localhost:6379", "your_db_service")
```

Write circuit breaker Log
```
pena.WriteLog(pena.CircuitLog{
    App:       "your_app_name",
    ErrorCode: "your_error_code",
    Fail:      true,
    Host:      "host_destination_trip",
    Route:     "routing_destination_trip",
    Tripped:   false,
})
```

Write circuit breaker :

- Status Fail
    ```
    circuitStatus.SetFail("host_destination_trip", "routing_destination_trip")
    ```
- Status Tripped
    ```
    circuitStatus.SetTripped("host_destination_trip", "routing_destination_trip")
    ```
- Status SetClosed
    ```
    circuitStatus.SetClosed("host_destination_trip", "routing_destination_trip")
    ```

Sample Result :
- Mongo
    ```
    {
        "_id" : ObjectId("5a127c748c0be847e5a68cfe"),
        "host" : "general_transaction_rule@127.0.0.1:58083",
        "app" : "txn",
        "route" : "tran_code/incoming",
        "fail" : true,
        "tripped" : false,
        "createdat" : ISODate("2017-11-20T13:55:48.879+07:00"),
        "error_code" : "E02009999"
    },
    ```

- Redis
    
    - Key:
        ```
        sav_txn:log:general_transaction_rule@127.0.0.1:58083:tran_code/incoming
        ```
    - Value :
        ```
        {"Closed":false,"Fail":false,"Tripped":true,"LastUpdate":"2017-11-20T15:35:09.237181258+07:00"}
        ```