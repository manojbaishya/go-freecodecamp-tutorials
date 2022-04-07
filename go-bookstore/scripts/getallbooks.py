import requests as http
import default

url = default.get_url(endpoint="/book")

response = http.get(url)

default.pprint_response(response)
