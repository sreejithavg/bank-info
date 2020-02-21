**Get all users after 15th october **
----
    Get all  users information and count falls after 15th october

  
* **URL**

       /user/filter

* **Method:**

       POST
  
* **URL Params**

       none

* **Data Params**

       {
       	"start_age":20,
       	"end_age":30,
       	"martial_status":"married"
       }

* **Success Response:**

  * **Code:** 200 OK 
    
       **Content:** 
    
        {
            "status_code": 200,
            "status_description": "OK",
            "message": "number of  users between specified age group and martial status",
            "data": 185
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
    OR
    * **Code:** 400 BAD REQUEST 
       
        **Content:** 
            
            {
                "status_code": 400,
                "status_description": "failed",
                "message": "Key: '' Error:Field validation for '' failed on the 'gtfield' tag",
                "data": null
            }