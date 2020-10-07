# papersvc

How to running apps

1. Clone this project first
2. Running migration file in folder migration
3. Enter Clone Folder using command line
4. Write in command line "make run"
5. Register to endpoint POST   /v1/register 
6. Login to POST   /v1/login 
7. Put token to header authorization

List Of Endpoint

1. POST   /v1/register
   Content-Type : application/json
   Payload : 
   {
    "username":string,
    "password":string
   }
   
2. POST   /v1/login
   Content-Type : application/json
   Payload : 
   {
    "username":string,
    "password":string
   }
   
3. POST   /v1/trx
   Authorization : token
   Content-Type : application/json
   Payload : 
   {
    "account_id":integer
   }
   
4. GET    /v1/trx
   Authorization : token
   Query Param : 
   - account_id
   - id
   - sort
   - page
   - limit
   
5. GET    /v1/trx/:id  
   Authorization : token
   
6. PUT    /v1/trx/:id  
   Authorization : token
   Content-Type : application/json
   Payload : 
   {
    "account_id":integer
   }
7. DELETE /v1/trx/:id   
   Authorization : token
   
8. POST   /v1/account 
   Authorization : token
   Content-Type : application/json
   Payload : 
   {
    "name":string
   }
9. GET    /v1/account    
   Authorization : token
   
10. GET    /v1/account/:id 
    Authorization : token
    Query Param : 
    - id
    - name
    - sort
    - page
    - limit
11. PUT    /v1/account/:id  
    Authorization : token
    Content-Type : application/json
    Payload : 
    {
      "name":string
    }
12. DELETE /v1/account/:id   
    Authorization : token
