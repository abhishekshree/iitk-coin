# iitk-coin

Database Used: `SQLite`

The database currently has two tables, they are:

- User

```sql
CREATE TABLE User (
    rollno TEXT,
    name TEXT,
    password TEXT,
    coins REAL,
    Admin BOOLEAN DEFAULT 0,
    PRIMARY KEY(rollno)
);
```

- Transactions

```sql
CREATE TABLE Transactions (
    id INTEGER PRIMARY KEY,
    from_roll TEXT,
    to_roll TEXT,
    type TEXT,
    timestamp TEXT,
    amount_before_tax REAL,
    tax REAL
);
```

## Details of the endpoints:

### Signup

```
url : /signup
method : POST

Request Body: {
    "Roll" : "",
    "Name" : "",
    "Password" : ""
}

Response : {
            "success": true/false
}

```

### Login

```
url : /login
method : POST

Request Body: {
    "Roll" : "",
    "Password" : ""
}

Response : {
            "token": JWT Token
            "status": true/false
}
NOTE: Token is returned only after a successful login. Also it expires in 3 days.
```

### Secretpage

```
url : /secretpage
method : GET

Response : "This is a very secret string."
NOTE: Can access only after a successful JWT verification and if the user Exists.
```

---

### Get Coins

```
url : /getCoins
method : GET

Request Body: {
    "rollno" : "",
}

Response : {
            "rollno": <Roll Number>,
            "coins": <Coins currently held by user>,
}
```

### Award Coins

```
url : /awardCoins
method : POST (JWT Required)
Request Body: {
    "rollno" : "",
    "amount": <float>
}

Response : {
            "message": "Coins Awarded."
}
Note: Check from JWT if the amount coming from user X is actually after when user X logs in and is an admin.
```

### Transfer Coins

```
url : /transferCoins
method : POST (JWT Required)

Request Body: {
    "from" : "",
    "to": "190028",
    "amount": <float>
}

Response : {
            "message": "Coins Transferred.",
            "amount":  <Amount Transferred after tax deduction>,
}
Note: Check from JWT if the amount coming from user X is actually after when user X logs in.
```
---

## Sidenote:

Also defined some functions to give or take admin privileges in the db package.

```go

func MakeAdmin(rollno string) bool
func RemoveAdmin(rollno string) bool
func IsAdmin(rollno string) bool

```