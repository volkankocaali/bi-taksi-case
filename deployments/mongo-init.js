db = db.getSiblingDB("bitaksi");

db.createCollection("users");

db.users.insertMany([
    {
        "_id": ObjectId("671ee0fc464bde4c5638daa1"),
        "password": "$2a$10$pu5.zNv3efrxAH1z66CTdOWFZNKmQqyU31E4u6781moDsXlUOW5cO",
        "service": "matching-service",
        "username": "volkan.kocaali"
    },
    {
        "_id": ObjectId("671ee18dec10f03e4015e3ac"),
        "password": "$2a$10$gMD1IFW.o/HJGzycuy/pcODNU3WTFYQnhu5yNhB3hPQS6M7RaY/2.",
        "service": "driver-service",
        "username": "john.doe"
    }
]);
