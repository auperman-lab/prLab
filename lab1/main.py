
from bs4 import BeautifulSoup
from urllib.request import urlopen
from prettytable import PrettyTable

base_url = "https://linella.md"
milk_url = f"{base_url}/ro/catalog/produse_lactate?page=1"


def get_soup(url):
    page = urlopen(url)  # task1
    html = page.read().decode("utf-8")
    return BeautifulSoup(html, "html.parser")


soup = get_soup(milk_url)
table = PrettyTable()




product = soup.find_all("div", {"class": ["products-catalog-content__item", "products-catalog-content__item_marked"]})

product_data = dict()
for p in product:

    name_div = p.find("div",  {"class": "products-catalog-content__body"})
    p_name = name_div.find("a", {"class": "products-catalog-content__name"}).string #task2

    p_link = name_div.find("a", {"class": "products-catalog-content__name"})
    p_link = p_link["href"]

    product_soup = get_soup(base_url+p_link)
    product_breadcrumbs = [i.get_text(strip=True) for i in
                           product_soup.find("ul", {"class": "breadcrumbs"}).find_all("li")]

    product_category =  product_breadcrumbs[2]
    product_subcategory =  product_breadcrumbs[3]


    if name_div.find("span", {"class": ["price-products-catalog-content__static"]}): #task4
        price_old = name_div.find("span", {"class": ["price-products-catalog-content__static"]}).get_text(strip=True) #task4 every string is stripped
        price_new = ""
        discount = ""
    else:
        price_old = name_div.find("span", {"class": "price-products-catalog-content__old"}).get_text(strip=True) #task2
        price_new = name_div.find("span", {"class": "price-products-catalog-content__new"}).get_text(strip=True) #task3
        discount = name_div.find("div", {"class": "price-products-catalog-content__discount"}).get_text(strip=True) #task3



    product_data[p_name] = [ product_category, product_subcategory, price_old, price_new, discount, base_url + p_link]





table.field_names = ["category", "subcategory", "name", "old_price", "new_price", "discount", "link"]
for k, v in product_data.items():
    table.add_row([k, v[0], v[1], v[2], v[3], v[4], v[5]])

print(table)


