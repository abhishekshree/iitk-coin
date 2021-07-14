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

- RedeemRequests

```sql
CREATE TABLE RedeemRequests (
    id INTEGER PRIMARY KEY,
    rollno TEXT,
    item TEXT,
    timestamp TEXT,
    status INTEGER DEFAULT 0
);

-- for status: 0 -> Pending, 1 -> Redeemed, 2 -> Declined
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

## Redeem Logic

### Get a list of Redeemable items and price

```
url : /getRedeemList
method : GET

Request Body: <None>

Response : <Redeemable items>
```

### Create a Redeem Request


```
url : /redeemRequest

method : POST

Request Body: {
    "item": "B"
}

Response : {
    "message": <Message>
}

Note: The Roll number is obtained from active JWT
```

### Accept a Redeem Request


```
url : /acceptRedeemRequest

method : POST

Request Body: {
    "id":6
}

Response : {
    "message": <Message>
}

Note: This route is ADMIN ONLY.
```

### Reject a Redeem Request


```
url : /rejectRedeemRequest

method : POST

Request Body: {
    "id":6
}

Response : {
    "message": <Message>
}

Note: This route is ADMIN ONLY. Also, a request which was accepted earlier can be rejected too, in that case the coins are lost (like in other coin based systems).
```

### Reject all the pending requests from a user (in case someone spams a lot of requests)

```
url : /rejectPendingRequests

method : POST

Request Body: {
    "roll":"190028"
}

Response : {
    "message": <Message>
}

Note: This route is ADMIN ONLY.
```