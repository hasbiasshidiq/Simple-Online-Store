import requests
import json

url = "http://127.0.0.1:8888/update-inventory"

payload = {
    "seller_id": "barokah_store",
    "product_id": 3,
    "product_name": "Flexamove",
    "quantity": 5
}

if __name__ == "__main__":
    response = requests.put(url, json = payload)
    # print(json.loads(response.text))
    # print(response.content)
    try:
        parsed = json.loads(response.text)
        json_object = json.dumps(parsed, indent=4, sort_keys=True)
        print(json_object)

        with open('output_update_inventory.txt', 'w') as outf:
            outf.write(json_object)
    except:
        print(response.text)
        with open('output_update_inventory.txt', 'w') as outf:
            outf.write(response.text)