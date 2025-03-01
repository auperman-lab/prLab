from bs4 import BeautifulSoup
from lab1.src.model.product import Product
from lab1.src.scraper import http_scraper, tcp_scraper


def parse_products(html_content):
    soup = BeautifulSoup(html_content, 'html.parser')
    products = []
    for category in soup.find_all('div', {"class":["k-product-section"]})[2:-2]:
        category_name = category.find('h2', {"class":["k-product-section__headline"]}).get_text(strip=True)


        for item in category.find_all('div', {"class": ["k-product-grid__item"]}):

            p_name = item.find("div", {"class": "k-product-tile__title"}).string
            p_subname = item.find("div", {"class": "k-product-tile__subtitle"})
            p_subname = p_subname.string if p_subname.string else ""
            p_name = p_name + " " + p_subname


            # todo -find link for this
            # p_link = name_div.find("a", {"class": "products-catalog-content__name"})
            # p_link = p_link["href"]


            price_div = item.find("div", {"class": ["k-product-tile__pricetag"]})

            if price_div.find("div", {"class": "k-price-tag__old-price"}):
                price_old = price_div.find("span", {"class": "k-price-tag__old-price-line-through"}).get_text(strip=True)
            else:
                price_old = "0"
            price_now = price_div.find("div", {"class": "k-price-tag__price"}).get_text(strip=True)
            discount = price_div.find("div", {"class": "k-price-tag__discount"}).get_text(strip=True)

            product = Product(
                name=p_name,
                link="",
                price_old= price_old,
                price_now=price_now,
                discount=discount,
                category=category_name,
                sub_category=category_name
            )

            print(product)
            products.append(product)



    return products

