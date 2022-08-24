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


## CRUD

**_id**

Every document **must** have a unique *_id* field & should have a value.

By default, _id has a value of ```ObjectId()```. But you can use custom value for "_id". It can have any object, and doesn't need to be ```ObjectId()```. Always ensure the value to **_id** must be unique. 

Note: While inserting a document, adding ```_id``` is not mandatory, because MongoDB can automatically generate a unique key for you before insertion. 

More about Object ID: https://www.mongodb.com/docs/manual/reference/method/ObjectId/#objectid

### Create

-Make sure the documents you're about to create/insert must have a unique _id.

#### Single Document

```js
db.collection_name.insert()
```

##### Example

```js
db.inspections.insert({
      "_id" : ObjectId("56d61033a378eccde8a8354f"),
      "id" : "10021-2015-ENFO",
      "certificate_number" : 9278806,
      "business_name" : "ATLIXCO DELI GROCERY INC.",
      "date" : "Feb 20 2015",
      "result" : "No Violation Issued",
      "sector" : "Cigarette Retail Dealer - 127",
      "address" : {
              "city" : "RIDGEWOOD",
              "zip" : 11385,
              "street" : "MENAHAN ST",
              "number" : 1712
         }
  })
```

#### Multiple Document

```js
db.collection_name.insert([])
```

##### Example

```js

db.inspections.insert([ { "test": 1 }, { "test": 2 }, { "test": 3 } ])
```

```ordered```

MongoDB allows you to insert many documents in a single go, but the problem arises when a developer specifices the _id which have same values, hence, not making it unique. 

```js
db.inspections.insert([{ "_id": 1, "test": 1 },{ "_id": 1, "test": 2 },
                       { "_id": 3, "test": 3 }])
```

In tha above code, 2 documents have same value for ```_id``` field. MongoDB has a rule to iterate and insert documents in the given order form. So first it iterates through index 0 and inserts that document. In the next iteration, it finds the duplicate key error and hence it stops the iteration. So even though the next document has a unique ```_id```, it does not gets inserted becaused the iteration has been terminated. 

To fix this; we use ```{"ordered": false}``` in ```.insert()``` which adds all the documents with unique ids and then throws the error for non-unique documents.

```js
db.inspections.insert([{ "_id": 1, "test": 1 },{ "_id": 1, "test": 2 },
                       { "_id": 3, "test": 3 }],{ "ordered": false })
```

---

### Update

#### Update One

```js
db.collection_name.updateOne(query, {update_operator: document})
```

##### Examples

```js

db.zips.updateOne({ "zip": "12534" }, { "$set": { "pop": 17630 } })
```

```js

db.grades.updateOne({ "student_id": 250, "class_id": 339 },
                    { "$push": { "scores": { "type": "extra credit",
                                             "score": 100 }
                                }
                     })
```


#### Update Many

```js
db.collection_name.updateOne(query, {update_operator: document})
```

##### Example

```js

db.zips.updateMany({ "city": "HUDSON" }, { "$inc": { "pop": 10 } })
```


#### Update Operators

![image](https://user-images.githubusercontent.com/28825619/186328480-399aabb9-4aeb-4af2-9cbf-b2cdb30032a8.png)

