package main

import (
	"encoding/json"
	"fmt"
)

type Value struct {
	Kind string
	Name string
	Status string
}

type Data struct {
	Key string
	Value Value

}

type Result struct {
	Data Data
}

func main() {
	dataJson := `{
    "data": {
        "key": "planet10",
        "value": {
            "kind": "planet",
            "name": "Mercury",
            "status": "active"
        }
    },
    "datacontenttype": "application/json",
    "id": "b8fb3e63-ec90-4626-96ad-dd5114d2ea75",
    "pubsubname": "planetpubsub",
    "source": "python-publisher",
    "specversion": "1.0",
    "topic": "planets",
    "traceid": "00-41a25578caab2b80882e57805aaa56e0-286972373ca97326-01",
    "traceparent": "00-41a25578caab2b80882e57805aaa56e0-286972373ca97326-01",
    "tracestate": "",
    "type": "com.dapr.event.sent"
}`
	var result Result
	json.Unmarshal([]byte(dataJson), &result)
	fmt.Println(result.Data)
}