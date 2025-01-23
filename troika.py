#Write a program to extract text from https://bombay.indology.info/mahabharata/text/ASCII/MBh01.txt and https://bombay.indology.info/mahabharata/text/UR/MBh01.txt and append it to the MBh01.json file
import requests
import json
def extract_text_from_url(url):
    response = requests.get(url)
    return response.text
def append_text_to_json(json_file, text):
    with open(json_file, 'r', encoding='utf-8') as file:
        data = json.load(file)
        data['ascii'] = text
        with open(json_file, 'w', encoding='utf-8') as file:
            json.dump(data, file, ensure_ascii=False, indent=4
                      # text = extract_text_from_url("https://bombay.indology.info/mahabharata/text/ASCII/MBh01.txt")
                      # append_text_to_json("D:/Projects/oper/MBh/MBh01.json", text)
                      ) 
for i in range(1, 25):
    text = extract_text_from_url(f"https://bombay.indology.info/mahabharata/text/ASCII/MBh{i:02d}.txt")