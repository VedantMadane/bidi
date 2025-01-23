import json

def create_raghava_yadaviyam_json():
    text_parts = text.split("||")
    json_obj = {}
    
    for i in range(1, 31):
        json_obj[str(i)] = {
            "anulom": text_parts[i-1].strip(),
            "pratilom": text_parts[i-1][::-1].strip()
        }
    
    return json_obj

# Get text from hk-ascii.txt
text = open("hk-ascii.txt", "r").read()

result = create_raghava_yadaviyam_json()
print(json.dumps(result, indent=2, ensure_ascii=False))
