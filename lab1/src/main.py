
import requests
from requests.auth import HTTPBasicAuth

from lab1.src.consts import base_url
from lab1.src.model.products_price_range import ProductsPriceRange
from lab1.src.parser import parse_products
from lab1.src.scraper import http_scraper, tcp_scraper, send_custom_post_request
from lab1.src.serialize import serialize_json, serialize_xml, serialize_sql
from lab1.src.value_converter import convert_price


def main():
    milk_path = "/ro/catalog/produse_lactate?page=1"
    milk_url = "https://" + base_url + milk_path

    # html = html_scraper(milk_url)

    # html = tcp_scraper(milk_path)
    #
    #
    # products = parse_products(html)
    #
    # for product in products:
    #     product.price_eur  = convert_price(product)
    #
    # cheap_products = ProductsPriceRange(products, 0, 20)
    # print(cheap_products)



    # --------------------------------------------------------------------------------------

    data = [
        {
            "name": "Yagotinsky Branza",
            "price_old": "34.50",
            "price_new": "24.90",
            "discount": "-27%",
            "category": "Dairy Products",
            "sub_category": "Cottage Cheese",
            "link": "/ro/catalog/branza_proaspata/yagotinskiy_branza_de_casa_9__350g"
        },
        {
            "name": "Yagotinsky Branza",
            "price_old": "34.50",
            "price_new": "24.90",
            "discount": "-27%",
            "category": "Dairy Products",
            "sub_category": "Cottage Cheese",
            "link": "/ro/catalog/branza_proaspata/yagotinskiy_branza_de_casa_9__350g"
        }
    ]
    #
    # print(serialize_sql("data",  data))
    #



# --------------------------------------------------------------------------------------
    username = "302"
    password = "503"

    headers, response_body = send_custom_post_request("/upload", data, username, password)

    print(headers)
    print(response_body)





if __name__ == "__main__":
    main()