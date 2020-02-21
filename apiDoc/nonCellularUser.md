**Get all  non cellular user count**
----
    Get all non-cellular user 

  
* **URL**

       /user/non_cellular

* **Method:**

       GET
  
* **URL Params**

       none

* **Data Params**

       None

* **Success Response:**

  * **Code:** 200 OK 
    
       **Content:** 
    
        {
                    "status_code": 200,
                    "status_description": "OK",
                    "message": "number of non-cellular user",
                    "data": 1625
        }
 
* **Error Response:**

  * **Code:** 500 INTERNAL SERVER ERROR 
   
    **Content:** 
        
        {
                            "status_code": 500,
                            "status_description": "Internal Server Error",
                            "message": "number of non-cellular user",
                            "data": null
        }


