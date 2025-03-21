


Display request format and output
```
curl -X GET "http://localhost:8080/display?device_id=1&start=1742515530&end=1742515758"
```
```
[{"Timestamp":1742515530,"Value":89},{"Timestamp":1742515532,"Value":81},{"Timestamp":1742515534,"Value":0},{"Timestamp":1742515536,"Value":25},{"Timestamp":1742515538,"Value":72},{"Timestamp":1742515540,"Value":1},{"Timestamp":1742515542,"Value":46},{"Timestamp":1742515745,"Value":13},{"Timestamp":1742515747,"Value":23},{"Timestamp":1742515749,"Value":90},{"Timestamp":1742515751,"Value":65},{"Timestamp":1742515753,"Value":87},{"Timestamp":1742515755,"Value":87},{"Timestamp":1742515757,"Value":42}]
```


Server makes a call to sensor every 3 seconds 
Sensor returns timestamp, device id, temp, device type . Server logs shown here - 

```
 go run .
```
```
2025/03/20 19:11:48 started
Sensor data: &{1 52 1742515910 sensor}
Sensor data: &{1 8 1742515912 sensor}
[{1742515530 89} {1742515532 81} {1742515534 0} {1742515536 25} {1742515538 72} {1742515540 1} {1742515542 46} {1742515745 13} {1742515747 23} {1742515749 90} {1742515751 65}]
Sensor data: &{1 4 1742515914 sensor}
Sensor data: &{1 5 1742515916 sensor}
Sensor data: &{1 84 1742515918 sensor}
[{1742515530 89} {1742515532 81} {1742515534 0} {1742515536 25} {1742515538 72} {1742515540 1} {1742515542 46} {1742515745 13} {1742515747 23} {1742515749 90} {1742515751 65} {1742515753 87} {1742515755 87} {1742515757 42}]
Sensor data: &{1 35 1742515920 sensor}
Sensor data: &{1 51 1742515922 sensor}

``
