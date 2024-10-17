
EUR_TO_MDL = 19.50
MDL_TO_EUR = 1 / EUR_TO_MDL

def convert_price(product, to_currency='EUR'):
    price_value = product.price_now

    if to_currency == 'EUR':  # Assuming category mentions currency
        price_value *= MDL_TO_EUR
    elif to_currency == 'MDL' :
        price_value *= EUR_TO_MDL

    return round(price_value, 2)



