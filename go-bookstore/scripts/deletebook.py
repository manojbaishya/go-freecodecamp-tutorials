import requests as http
import default

id = 2
url = default.get_url(endpoint=f"/book/{id}")

response = http.delete(url)

default.pprint_response(response)
