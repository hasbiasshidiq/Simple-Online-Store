# HOW TO USE
```
# build store api app image
$ make build-appImage

# run store app docker instance and postgresql docker instance
$ make run

# stop running docker
$ make stop
```

## Overview ERD DESIGN
This is database schemes of our simple online store
![alt text](https://github.com/hasbiasshidiq/Simple-Online-Store/blob/main/images/ERD.png)


## API Testing

Our api testing scripts is written in python3. So if your working environment is not compatible and you are looking for neat implementation, you can build your python testing environment with docker as explained below :

```
$ cd api-testing && build -t testing-api .
```

Don't forget to make sure if application and postgresql container running properly. 

```
$ docker ps
```

Still on api-testing directory, you can run api testing by following these recommended steps. You can also change parameter value or payload defined in python scripts

### 1. View Inventory
return list of inventory description, you can filter it by sellerID and Category. In this scenario we set parameter value as described below

---
* HTTP GET request "http://127.0.0.1:8888/view-inventory"
* Params Value:

    seller_id=barokah_store
    category=herbal
---


```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 1_view_inventory.py
```
or
```
$ python3 1_view_inventory.py
```

### 2. Update Inventory
Update inventory Quantity by sellerID and productID

---
```
**Description**

* HTTP PUT "http://127.0.0.1:8888/update-inventory"
* Payload
{
    "seller_id": "barokah_store",
    "product_id": 3,
    "product_name": "Flexamove",
    "quantity": 5
}
```
---

```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 2_update_inventory.py
```
or
```
$ python3 2_update_inventory.py
```

### 3. Checkout Order
Customer checkout some products from a particular seller

- HTTP POST
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
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 3_checkout_order.py
```
or
```
$ python3 3_checkout_order.py
```

### 4. Checkout Order Concurrent
Make two concurrent order request with 4 items of flexamove(product name), but before that we need to make a request as in step 2 so Flexamove quantity in inventory is set to 5 
- HTTP POST
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
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 4_checkout_order_concurrent.py
```
or
```
$ python3 4_checkout_order_concurrent.py
```

the output should be like this:
```
<Response [202]>
Success
<Response [202]>
Success
```

then check Flexamove quantity in inventory by doing step 1. Isn't it -3?

### 5. Checkout Order Concurrent with Row Locking
Make two concurrent order request with 4 items of flexamove(product name), but before that we need to make a request as in step 2 so Flexamove quantity in inventory is set to 5 

---
- HTTP POST
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
---

```
$ docker run --rm -it --net=host -v `pwd`/:/code -w /code testing-api:latest python3 5_checkout_order_concurrent_withLock.py
```
or
```
$ python3 5_checkout_order_concurrent_withLock.py
```

the output should be like this:
```
<Response [202]>
Success
<Response [202]>
{"failed_order_items":[{"order_item_id":0,"order_id":415,"product_id":3,"product_name":"Flexamove","quantity":4,"total_price":10000}]}
```

then check Flexamove quantity in inventory by doing step 1. Isn't it 1?