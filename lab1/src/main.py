from lab1.src.consts import base_url
from lab1.src.parser import parse_all_categories, parse_products
from lab1.src.scraper import http_scraper
import json
import pika
import time


def send_products_to_queue(products, queue_name='products'):
    # Connect to RabbitMQ
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    channel = connection.channel()

    # Declare the queue
    channel.queue_declare(queue=queue_name)

    # Publish each product as a separate message
    for product in products:
        product_dict = product.to_dict()  # Convert product to a dictionary
        product_json = json.dumps(product_dict)  # Serialize to JSON
        channel.basic_publish(exchange='', routing_key=queue_name, body=product_json)
        print(f" [x] Sent product: {product.name}")
        time.sleep(1)

    connection.close()



def main():
    milk_path = "/ro/catalog/produse_lactate?page=1"
    milk_url = "https://" + base_url + milk_path
    # categories_path = "/ro/catalog"
    # categories_url = "https://"+base_url+ categories_path
    #
    html = http_scraper(milk_url)

    products = parse_products(html)
    send_products_to_queue(products)



    pass







if __name__ == "__main__":
    main()