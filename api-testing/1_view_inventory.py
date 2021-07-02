import requests
import json

url = "http://127.0.0.1:8888/view-inventory"

PARAMS = {
    'seller_id':'barokah_store',
    'category':'herbal'
    }

# # other example
# PARAMS_EX_1 = {
#     'seller_id':'all',
#     'category':'all'
#     }

# PARAMS_EX_2 = {
#     'seller_id':'barokah_store',
#     'category':'all'
#     }

if __name__ == "__main__":
    response = requests.get(url, params = PARAMS)
    # print(json.loads(response.text))
    # print(response.content)
    try:
        parsed = json.loads(response.text)
        json_object = json.dumps(parsed, indent=4, sort_keys=True)
        print(json_object)

        with open('output_view_inventory.txt', 'w') as outf:
            outf.write(json_object)
    except:
        print(response.text)
        with open('output_view_inventory.txt', 'w') as outf:
            outf.write(response.text)