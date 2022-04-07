import requests as http
import default

url = default.get_url(endpoint="/book")

headers = {"Content-Type": "application/json"}
body = """{
    "Name": "A Tale of Two Cities",
    "Author": "Charles Dickens",
    "Publication": "Harper Collins"
}
"""

response = http.post(url, data=body, headers=headers)

default.pprint_response(response)
