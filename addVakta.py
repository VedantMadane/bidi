import os
def process_json_with_speaker(input_json, output_file_path):
    processed_data = {}
    
    # Convert JSON to dictionary
    os.makedirs(output_folder, exist_ok=True)
    
    # Iterate through all JSON files in the input folder
    for filename in os.listdir(input_folder):
        if filename.endswith('.json'):
            input_file_path = os.path.join(input_folder, filename)
            output_file_path = os.path.join(output_folder, filename)
            
            # Read input JSON file
            with open(input_file_path, 'r', encoding='utf-8') as file:
                data = json.load(file)
    # print(input_json)
                data_dict = dict(input_json)
    
    # Track the last numeric key's text
                last_speaker = None
    
    # Sort keys to maintain order
                sorted_keys = sorted(data_dict.keys())
    
                for i, key in enumerate(sorted_keys):
        # If key is numeric, update last_speaker
                    if key.isnumeric():
                        last_speaker = data_dict[key]['text']
        
        # If key has alphabetic suffix and last_speaker exists
                    if not key.isnumeric() and last_speaker:

                        processed_data[key]= data[key]['vaktā'] = last_speaker
    
                        # **data[key],
                
                # Write processed data to output file
                    with open(output_file_path, 'w', encoding='utf-8') as outfile:
                        json.dump(processed_data, outfile, ensure_ascii=False, indent=2)
                        return processed_data

# Example usage
input_json = {
    "04001001": {
        "text": {
            "ud": "जनमेजय उवाच\n",
            "ur": "janamejaya uvāca\n",
            "ascii": "janamejaya uvAca\n"
        }
    },
    "04001001a": {
        "text": {
            "ud": "कथं विराटनगरे मम पूर्वपितामहाः\n",
            "ur": "kathaṁ virāṭanagare mama pūrvapitāmahāḥ\n",
            "ascii": "kathaM virATanagare mama pUrvapitAmahAH\n"
        }
    }
}
import json


input_folder = "D:\\Projects\\oper\\decomposed"
output_folder = "D:\\Projects\\oper\\vaktāsaha"
result = process_json_with_speaker(input_folder, output_folder)
print(json.dumps(result, ensure_ascii=False, indent=2))