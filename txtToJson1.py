import json
import requests
def convert_to_json(filename):
    result = {}
    current_key = None
    current_text = []

    with open(filename, 'r', encoding='utf-8') as file:
        for line in file:
            line = line.strip()
            if line.startswith(('1', '0')) and len(line.split()) > 0:
                # New key found
                if current_key:
                    result[current_key] = '\n'.join(current_text)
                current_key = line.split()[0]
                current_text = [line[len(current_key):].strip()]
            else:
                current_text.append(line)

        # Add last key's text
        if current_key:
            result[current_key] = '\n'.join(current_text)

    return result
# res = convert_to_json("D:/Projects/oper/MBh01.txt")

# Write res to a file
# with open("D:/Projects/oper/MBh01.json", 'w', encoding='utf-8') as json_file:
#     json.dump(res, json_file, ensure_ascii=False, indent=4)

# Write a function to iterate over https://bombay.indology.info/mahabharata/text/UD/MBh%02d.txt and convert each file to json

# Write a function that iterates over links and writes them to local

for i in range(18, 19):
    # Create file if not exists
    with open(f"D:/Projects/oper/MBh/MBh{i:02d}UR.txt", 'w', encoding='utf-8') as file:
        response = requests.get(f"https://bombay.indology.info/mahabharata/text/UR/MBh{i:02d}.txt")
        file.write(response.text)
for i in range(18, 19):
    res = convert_to_json(f"D:/Projects/oper/MBh/MBh{i:02d}UR.txt")
    with open(f"D:/Projects/oper/MBh/MBh{i:02d}UR.json", 'w', encoding='utf-8') as json_file:
        json.dump(res, json_file, ensure_ascii=False, indent=4)