import requests
import json

url = "http://127.0.0.1:8888/checkout-order"

payload = {
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

if __name__ == "__main__":
    response = requests.post(url, json = payload)
    # print(json.loads(response.text))
    # print(response.content)
    try:
        parsed = json.loads(response.text)
        json_object = json.dumps(parsed, indent=4, sort_keys=True)
        print(json_object)

        with open('output_checkout_order.txt', 'w') as outf:
            outf.write(json_object)
    except:
        print(response.text)
        with open('output_checkout_order.txt', 'w') as outf:
            outf.write(response.text)