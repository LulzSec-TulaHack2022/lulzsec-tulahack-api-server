# lulzsec-tulahack-api-server

Endpoints:
* /getcatalog - returns plants catalog

* /getcurrentweather - returns random weather parameters: *temperature*["low"|"medium"|"high"|"any"], *humidity*["low"|"medium"|"high"|"any"], *illumination*[integer: 0->100]

* /addflower - creates new plant to exact user. JSON Format:
```javascript
{
  "catalog_id": string,
  "owner_id": string,
  "name": string
}
```

* /getalluserflowers - returns all users plants. Method: *GET* with parameter *owner_id*. Example:
```javascript
https://lulzsec-tulahack-api-server.herokuapp.com/getalluserflowers?owner_id=<user id>
```

* /getuserflower - returns exact plant of exact user. Method: *GET* with parameter *flower_id*. Example:
```javascript
https://lulzsec-tulahack-api-server.herokuapp.com/getuserflower?flower_id=<flower id>
```

* /deleteflower - removes specified plant. Method: *GET* with parameter *flower_id*. Example:
```javascript
https://lulzsec-tulahack-api-server.herokuapp.com/deleteflower?flower_id=<flower id>
```

* /dead - modifies plants field *alive* to *false*. It means that the plant is dead. Method: *GET* with parameter *flower_id*. Example:
```javascript
https://lulzsec-tulahack-api-server.herokuapp.com/dead?flower_id=<flower id>
```

___
Every plant has unique id generated by XID module.
