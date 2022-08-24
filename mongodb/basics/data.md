# Data

MongoDB uses BSON - Binary JavaScript Object Notation

JSON - JavaScript Object Notation 

https://www.mongodb.com/json-and-bson

## Importing & Exporting Data

### JSON 

```mongoimport```

Let's you import data from JSON file.

```javascript
mongoimport --uri="mongodb+srv://<your username>:<your password>@<your cluster>.mongodb.net/sample_supplies" --drop sales.json
```

---

```mongoexport```

Exports data in JSON

```javascript
mongoexport --uri="mongodb+srv://<your username>:<your password>@<your cluster>.mongodb.net/sample_supplies" --collection=sales --out=sales.json
```

---

### BSON:

```mongorestore```

Let's you import data from BSON file.

```javascript
mongorestore --uri "mongodb+srv://<your username>:<your password>@<your cluster>.mongodb.net/sample_supplies"  --drop dump
```

---

```mongodump``` 

Let's you export data in BSON form 

```javascript
mongodump --uri "mongodb+srv://<your username>:<your password>@<your cluster>.mongodb.net/sample_supplies"
```

---


## Commands

### Establishing Connection to Cluster

```js
mongo "mongodb+srv://<username>:<password>@<cluster>.mongodb.net/admin"
```

### List All Databases

```js
show dbs
```

### Use a Database

```js
use <db_name>
```

### List All Collections

```js
show collections
```

### Query Data 

To query data, we use ```.find()``` method.

```js
db.collection_name.find()
```

Inside ```()``` of ```find()``` method, we write JSON to query our data.

```js
db.zips.find({"state": "NY"})
db.zips.find({"state": "NY", "city": "ALBANY"})
db.zips.find({"state": "NY", "city": "ALBANY"}).pretty()
```

```.pretty()``` returns data in formatted way.


#### Count Documents

We use ```.count()``` method to count number of documents.

```js
db.zips.find({"state": "NY"}).count()
```

#### ```it```

Whenever we use ```.find()```, MongoDB displays a specific number of data only. To see more data, we use ```it``` command.
```it``` stands for iterate. This lets you to iterate your cursor to next sets of data.

**Cursor**
```.find()``` command returns an object called cursor. Cursor is a pointer to a result set of a query.
A ```pointer``` is also an object returned by ```cursor``` which directs the address of the memory location.




