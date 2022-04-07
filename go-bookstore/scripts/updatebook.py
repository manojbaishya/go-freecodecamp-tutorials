import requests as http
import default

id = 2
url = default.get_url(endpoint=f"/book/{id}")

headers = {"Content-Type": "application/json"}
body = """{
    "Name": "A Tale of Two Cities",
    "Author": "Charles Dickens",
    "Publication": "Penguin UK"
}
"""

response = http.put(url, data=body, headers=headers)

default.pprint_response(response)
