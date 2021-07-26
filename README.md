# IITK Coin

[![](https://img.shields.io/docker/cloud/build/abhishekshree/iitk-coin?style=flat-square)](https://hub.docker.com/r/abhishekshree/iitk-coin)
[![Netlify](https://img.shields.io/netlify/3fccc76a-3ea3-4141-9cf2-e152c5f65be8?style=flat-square)](https://iitk-coin-docs.netlify.app/)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/abhishekshree/iitk-coin?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/abhishekshree/iitk-coin?style=flat-square)

This repository contains the backend for IITK Coin, which is a centralized pseudo-coin system in golang (fiber) for use in IITK Campus.

## <u>Index</u>

- [Directory Structure](#directory-structure)
- [Database](#database-sqlite)
- [Details of the endpoints](#details-of-the-endpoints)

---

## <u>Directory Structure</u>
<br />

```
├── config
│   └── config.go
├── db
│   ├── admin.go
│   ├── db.go
│   ├── records.go
│   ├── redeem.go
│   └── transactions.go
├── Dockerfile
├── go.mod
├── go.sum
├── iitk-coin
├── main.go
├── middleware
│   ├── hash.go
│   └── jwt.go
├── README.md
├── routes
│   ├── redeem.go
│   ├── routes.go
│   └── transactions.go
└── Users.db

4 directories, 19 files
```

---

## <u>Database: `SQLite`</u>
The database currently has three tables, they are:
<br />

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

---

## <u>Details of the endpoints:</u>
Can also be viewed [here.](https://iitk-coin-docs.netlify.app/)
<br />

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

### Sidenote:

Also defined some functions to give or take admin privileges in the db package.

```go

func MakeAdmin(rollno string) bool
func RemoveAdmin(rollno string) bool
func IsAdmin(rollno string) bool

```

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
