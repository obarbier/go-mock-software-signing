using policy as our acl similar to vault implementation

Users -(has)-> policies -(provide)-> (access)

Policy Definition
```
  [
        "user/*" : {
            capabilities : [ "create", "read","update", "delete"]
        },

        "key/*"{
            capabilities : ["read"]
        }
  ]


create  --> POST
read    --> GET
update  --> PUT
delete  --> delete
```


some concerns:
    - how to do reconciliation on path variable i.e user/* as opposed to /user/:userid/list with different capabilities for example

how to store policy in database:
- reference https://bitworks.software/en/2017-10-20-storing-trees-in-rdbms.html
user/2, user/3, user/5,  user/4, key/1

PATH_ID             CHILD_ID                PATH                IS_END            IS_ROOT
1                   2,7                     null                false             true
2                   3,4,5,6                 user                false             false
3                   null                    2                   true              false
4                   null                    3                   true              false
5                   null                    5                   true              false
6                   null                    4                   true              false
7                   8                       key                 false             false
8                   null                    1                   true              false

Option2: using postgres LTREE

User_id --> PATH_ID