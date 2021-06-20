# iitk-coin

Database Used: `SQLite`

Details of the endpoints:

## Signup

```
url : /signup
method : POST

Response : {
            "success": true/false
}

```

## Login

```
url : /login
method : POST

Response : {
            "token": JWT Token
            "status": true/false
}
NOTE: Token is returned only after a successful login. Also it expires in 3 days.
```

## Secretpage

```
url : /secretpage
method : GET

Response : "This is a very secret string."
NOTE: Can access only after a successful JWT verification and if the user Exists.
```

---

## Get Coins

```
url : /getCoins
method : GET

Response : {
            "rollno": <Roll Number>,
            "coins": <Coins currently held by user>,
}
```

## Award Coins

```
url : /awardCoins
method : POST

Response : {
            "message": "Coins Awarded."
}
```

## Transfer Coins

```
url : /transferCoins
method : POST

Response : {
            "message": "Coins Transferred.",
            "amount":  <Amount Transferred after tax deduction>,
}
```
