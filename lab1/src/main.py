from lab1.src.consts import base_url
from lab1.src.model.products_price_range import ProductsPriceRange
from lab1.src.parser import parse_products
from lab1.src.scraper import http_scraper, tcp_scraper
from lab1.src.serialize import serialize_json, serialize_xml
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

    products_info = [
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
            "name": "Original Sheep Cheese",
            "price_old": "56.70/0.3kg",
            "price_new": "",
            "discount": "",
            "category": "Dairy Products",
            "sub_category": "Sheep Cheese",
            "link": "/ro/catalog/branza_proaspata/original_sheep_braza_sarata_de_oi__kg"
        },
        {
            "name": {"name00": {"n": 1, "n1": 1, "n2": 1}, "name1": "a", "name2": "a", "name3": "a"},
            "price_old": "49.90",
            "price_new": "39.90",
            "discount": "-20%",
            "category": "Dairy Products",
            "sub_category": "Mozzarella",
            "link": "/ro/catalog/branza_proaspata/mozzarella_block_500g"
        }
    ]


    print(serialize_xml( "product", products_info))




if __name__ == "__main__":
    main()