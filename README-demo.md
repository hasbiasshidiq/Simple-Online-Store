# HOW TO RUN
Clone this repository
```
$ git clone https://github.com/hasbiasshidiq/Simple-Online-Store
```
move to working directory
```
$ cd Simple-Online-Store
```

Build, run, and stop application
```
# build store api app image
$ make build-appImage

# run store app docker instance and postgresql docker instance
$ make run

# stop running docker
$ make stop
```

## Overview ERD DESIGN
Database schemes of simple online store
![alt text](https://github.com/hasbiasshidiq/Simple-Online-Store/blob/main/images/ERD.png)


## API Testing

This api testing scripts are all written in python3. So if your working environment is not compatible and you are looking for proper implementation, you can build your python testing environment with docker as explained below :

```
$ cd api-testing && build -t testing-api .
```

Don't forget to make sure if the application and postgresql container running properly (as a result of executing `make run` command in line 7). 

```
$ docker ps
```

Still on api-testing directory, you can run api testing by following these recommended steps. You can also change parameter value or payload defined in python scripts

### 1. View Inventory
return list of inventory description, you can filter these by sellerID and Category. In this scenario, parameter values are described below

**Description**
```
- HTTP GET request "http://127.0.0.1:8888/view-inventory"
- Params Value:

    seller_id=barokah_store
    category=herbal
```

**Command**
```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 1_view_inventory.py
```

**Output**
```
{
    "inventory": [
        {
            "category": "herbal",
            "price": "120000",
            "product_id": "3",
            "product_name": "Flexamove",
            "quantity": "5",
            "seller_id": "barokah_store",
            "seller_name": "Barokah Store"
        },
        {
            "category": "herbal",
            "price": "200000",
            "product_id": "4",
            "product_name": "Antapro Barata",
            "quantity": "75",
            "seller_id": "barokah_store",
            "seller_name": "Barokah Store"
        }
    ]
}
```

### 2. Update Inventory
Update inventory Quantity by sellerID and productID

**Description**
```
- HTTP PUT "http://127.0.0.1:8888/update-inventory"
- Payload
{
    "seller_id": "barokah_store",
    "product_id": 3,
    "product_name": "Flexamove",
    "quantity": 5
}
```

**Command**
```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 2_update_inventory.py
```

**Output**
```
Success
```

### 3. Checkout Order
Customer checkout some products from a particular seller

**Description**
```
- HTTP POST "http://127.0.0.1:8888/checkout-order"
- Payload
{
    "seller_id":"barokah_store",
    "customer_id": "bambang_pamungkas",
    "order_status": "checkout",
    "order_items": [
        {
            "product_id":3,
            "product_name": "Flexamove",
            "quantity": 3,
            "total_price": 10000 
        },
        {
            "product_id":4,
            "product_name": "Antapro Barata",
            "quantity": 100,
            "total_price": 75000 
        }   
    ]
}
```

**Command**
```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 3_checkout_order.py
```

**Output**
```
{
    "failed_order_items": [
        {
            "order_id": 1,
            "order_item_id": 0,
            "product_id": 4,
            "product_name": "Antapro Barata",
            "quantity": 100,
            "total_price": 75000
        }
    ]
}
```

### 4. Checkout Order Concurrent
Make two concurrent order request with 4 items of flexamove(product name), but before that we need to make a request as in step 2 so Flexamove quantity in inventory is set to 5 

**Description**
```
- HTTP POST "http://localhost:8888/checkout-order"
- 2X Payload 
{
    "seller_id":"barokah_store",
    "customer_id": "bambang_pamungkas",
    "order_status": "checkout",
    "order_items": [
        {
            "product_id":3,
            "product_name": "Flexamove",
            "quantity": 4,
            "total_price": 10000 
        }   
    ]
}
```

**Command**
```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 2_update_inventory.py
```
and then
```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 4_checkout_order_concurrent.py
```

**Output**
```
<Response [202]>
Success
<Response [202]>
Success
```

then check Flexamove quantity in inventory by doing step 1. Isn't it to be -3?

### 5. Checkout Order Concurrent with Row Locking
Make two concurrent order request with 4 items of flexamove(product name), but before that we need to make a request as in step 2 so Flexamove quantity in inventory is set to 5 

**Description**
```
- HTTP POST "http://localhost:8888/checkout-order-withLock"
- 2X Payload 
{
    "seller_id":"barokah_store",
    "customer_id": "bambang_pamungkas",
    "order_status": "checkout",
    "order_items": [
        {
            "product_id":3,
            "product_name": "Flexamove",
            "quantity": 4,
            "total_price": 10000 
        }   
    ]
}
```

**Command**
```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 2_update_inventory.py
```
and then
```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 5_checkout_order_concurrent_withLock.py
```

**Output**
```
<Response [202]>
Success
<Response [202]>
{"failed_order_items":[{"order_item_id":0,"order_id":415,"product_id":3,"product_name":"Flexamove","quantity":4,"total_price":10000}]}
```

then check Flexamove quantity in inventory by doing step 1. Isn't it to be 1?