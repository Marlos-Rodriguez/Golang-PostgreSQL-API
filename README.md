# Golang-PostgreSQL-API
A Golang REST API with PostgreSQL of To-Do

# Install
Copy the repo and create a ```.env``` with the variables ```PORT``` and ```DATABASE_URL```
```
PORT=8080
DATABASE_URL=Your URL of PostgreSQL
```

# Paths
 **Get All To-Do's**:   
```Get``` ```api/todos```

Output

```json
{
    "Success": true,
    "Data": [
        {
            "ID": 1,
            "Description": "New to-do"
        }
    ],
    "Errors": []
}
```

**Create new To-Do:**

```POST``` ```api/todos```

body: 
```json
{
	"description":  "to-do"
}
```

Output

```
{
    "Success": true,
    "Data": [
        {
            "ID": 1,
            "Description": ""
        }
    ],
    "Errors": []
}
```

 **Get one To-Do**:   
```Get``` ```api/todos/{TodoID}```

Output

```json
{
    "Success": true,
    "Data": [
        {
            "ID": 1,
            "Description": "New to-do"
        }
    ],
    "Errors": []
}
```

 **Delete one To-Do**:   
```DELETE``` ```api/todos/{TodoID}```

Output

```json
{
    "Success": true,
    "Data": [
        {
            "ID": 1,
            "Description": ""
        }
    ],
    "Errors": []
}
```

 **Update one To-Do**:   
```PUT``` ```api/todos/{TodoID}```

body: 
```json
{
	"description":  "update to-do"
}
```

Output

```json
{
    "Success": true,
    "Data": [
        {
            "ID": 1,
            "Description": "update to-do"
        }
    ],
    "Errors": []
}
```
