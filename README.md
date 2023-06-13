# mysubly

Rest API with Golang for Open Sourced Version of Subly

|RESOURCE   |HTTP METHOD |ROUTE                    |DESCRIPTION|
|-----------|------------|-------------------------|------------
|sub        |GET         |/api/subs                |Retrieve all sub  ✅
|sub        |POST        |/api/subs/create               |Creates a new sub ✅
|sub	    |GET         |/api/subs/:subid          |Retrieve a single sub ✅
|sub	    |PATCH       |/api/subs/:subid          |Update a sub ✅
|sub	    |DELETE      |/api/subs/:subid          | Delete a sub ✅
|signup        |POST        |/api/signup               |Creates an account ✅
|profile    |PATCH       |/api/users/userid/image  |Update a profile picture
|sub        |GET         |/api/subs/overmontly     |montly sub spend overall
|sub        |GET         |/api/reports             |reports charts
|searchSub  |GET         |/api/search              |
