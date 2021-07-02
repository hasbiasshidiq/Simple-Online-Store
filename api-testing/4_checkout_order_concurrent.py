from threading import Thread
import requests

concurrent = 2

url = "http://localhost:8888/checkout-order"

payload = {
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

def doWork():
    PostOrder(url)

def PostOrder(url_order):
    try:
        # make a request
        response = requests.post(url_order, json=payload)
        print(response)
        print(response.text)
    except Exception as e:
        print(e)

if __name__ == "__main__":
    for i in range(concurrent):
        t = Thread(target=doWork)
        t.start()
