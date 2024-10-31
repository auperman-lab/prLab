def serialize_json(data, json_str="", modificator=""):
    if isinstance(data, dict):
        json_str = f"{modificator}{{\n"
        for k in data:
            if isinstance(data[k], dict):
                m = modificator + "\t"
                json_str += f"{modificator}\t{k}: \n" + serialize_json(data[k], json_str, m) + ",\n"
            else:
                json_str += f"{modificator}\t{k}: {data[k]},\n"
        json_str += f"{modificator}}}"

    if isinstance(data, list):
        json_str = f"{modificator}[\n"
        for k in data:
            if isinstance(k, list):
                m = modificator + "\t"
                json_str += serialize_json(k, json_str, m) + ",\n"
            json_str += serialize_json(k, json_str) + ",\n"
        json_str += f"{modificator}]"

    return json_str


def serialize_xml(tag, data):

    if isinstance(data, dict):
        xml_str = f"<{tag}>\n"
        for key, value in data.items():
            xml_str += serialize_xml(key, value)
        xml_str += f"</{tag}>\n"
        return xml_str

    elif isinstance(data, list):
        xml_str = ""
        for item in data:
            xml_str += serialize_xml(tag, item)
        return xml_str

    else:
        return f"<{tag}>{data}</{tag}>\n"


def serialize_sql(table_name, data):
    # Base case: if data is not a dictionary or list of dictionaries, return an empty string
    if not isinstance(data, (dict, list)):
        return ""

    # If data is a single dictionary, wrap it in a list for consistent processing
    if isinstance(data, dict):
        data = [data]

    # Recursively process each dictionary to build SQL statements
    sql_statements = []
    for record in data:
        # Extract columns and values
        columns = []
        values = []
        for key, value in record.items():
            columns.append(key)

            # Handle nested dictionaries or lists recursively
            if isinstance(value, dict):
                nested_sql = serialize_sql(key, value)
                values.append(f"({nested_sql})")
            elif isinstance(value, list):
                nested_sql = ", ".join(serialize_sql(key, item) for item in value)
                values.append(f"({nested_sql})")
            else:
                # Quote strings, leave numbers as they are
                if isinstance(value, str):
                    values.append(f"'{value}'")
                else:
                    values.append(str(value))

        # Construct the SQL statement
        columns_str = ", ".join(columns)
        values_str = ", ".join(values)
        sql_statement = f"INSERT INTO {table_name} ({columns_str}) VALUES ({values_str});"
        sql_statements.append(sql_statement)

    # Join all SQL statements
    return "\n".join(sql_statements)
