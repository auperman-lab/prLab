from operator import truediv


def serialize_json(object, json_str="", modificator=""):
    if isinstance(object, dict):
        json_str = f"{modificator}{{\n"
        for k in object:
            if isinstance(object[k], dict):
                m = modificator + "\t"
                json_str += f"{modificator}\t{k}: \n" +  serialize_json(object[k], json_str, m) + ",\n"
            else:
                json_str += f"{modificator}\t{k}: {object[k]},\n"
        json_str += f"{modificator}}}"

    if isinstance(object, list):
        json_str = f"{modificator}[\n"
        for k in object:
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