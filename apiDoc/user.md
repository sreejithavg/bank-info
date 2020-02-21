**Get all users after 15th october **
----
    Get all  user after 15th october

  
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
            "message": "users after 15th october",
            "data": {
                "records": 459,
                "bank_details": [
                    {
                        "age": 30,
                        "job": "unemployed",
                        "marital": "married",
                        "education": "primary",
                        "default": "no",
                        "balance": 1787,
                        "housing": "no",
                        "loan": "no",
                        "contact": "cellular",
                        "day": 19,
                        "month": "oct",
                        "duration": 79,
                        "campaign": 1,
                        "p_days": -1,
                        "previous": 0,
                        "p_outcome": "unknown",
                        "y": "no"
                    },
                    {
                        "age": 38,
                        "job": "management",
                        "marital": "divorced",
                        "education": "unknown",
                        "default": "no",
                        "balance": 0,
                        "housing": "yes",
                        "loan": "no",
                        "contact": "cellular",
                        "day": 18,
                        "month": "nov",
                        "duration": 96,
                        "campaign": 2,
                        "p_days": -1,
                        "previous": 0,
                        "p_outcome": "unknown",
                        "y": "no"
                    },
        			{
                        "age": 42,
                        "job": "management",
                        "marital": "divorced",
                        "education": "tertiary",
                        "default": "no",
                        "balance": 16,
                        "housing": "no",
                        "loan": "no",
                        "contact": "cellular",
                        "day": 19,
                        "month": "nov",
                        "duration": 140,
                        "campaign": 3,
                        "p_days": -1,
                        "previous": 0,
                        "p_outcome": "unknown",
                        "y": "no"
                    },
        			.
        			.
        			.
                    {
                        "age": 57,
                        "job": "self-employed",
                        "marital": "married",
                        "education": "secondary",
                        "default": "no",
                        "balance": 11494,
                        "housing": "no",
                        "loan": "no",
                        "contact": "cellular",
                        "day": 19,
                        "month": "nov",
                        "duration": 198,
                        "campaign": 1,
                        "p_days": -1,
                        "previous": 0,
                        "p_outcome": "unknown",
                        "y": "no"
                    },
                    {
                        "age": 54,
                        "job": "management",
                        "marital": "married",
                        "education": "primary",
                        "default": "no",
                        "balance": 1691,
                        "housing": "yes",
                        "loan": "no",
                        "contact": "cellular",
                        "day": 20,
                        "month": "nov",
                        "duration": 186,
                        "campaign": 1,
                        "p_days": -1,
                        "previous": 0,
                        "p_outcome": "unknown",
                        "y": "no"
                    },
                    {
                        "age": 41,
                        "job": "entrepreneur",
                        "marital": "married",
                        "education": "primary",
                        "default": "no",
                        "balance": 39,
                        "housing": "yes",
                        "loan": "no",
                        "contact": "cellular",
                        "day": 21,
                        "month": "nov",
                        "duration": 549,
                        "campaign": 2,
                        "p_days": -1,
                        "previous": 0,
                        "p_outcome": "unknown",
                        "y": "yes"
                    },
                ]
            }
        }
 
* **Error Response:**

  * **Code:** 500 INTERNAL SERVER ERROR 
   
    **Content:** 
        
        {
                            "status_code": 500,
                            "status_description": "Internal Server Error",
                            "message": "failed",
                            "data": null
        }