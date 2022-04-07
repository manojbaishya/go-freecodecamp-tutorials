import json


def get_url(endpoint):
    base_url = ""
    with open("env.json") as env:
        filedata = json.load(env)
        base_url = filedata["base_url"]

    url = f"{base_url}{endpoint}"
    return url


def pprint_response(response):
    print("Response Body:")
    print(json.dumps(json.loads(response.text), indent=4))
