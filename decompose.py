import os
import json

def process_json_files(input_folder, output_folder):
    # Ensure output folder exists
    processed_data = {}
    os.makedirs(output_folder, exist_ok=True)
    
    # Iterate through all JSON files in the input folder
    for filename in os.listdir(input_folder):
        if filename.endswith('.json'):
            input_file_path = os.path.join(input_folder, filename)
            output_file_path = os.path.join(output_folder, filename)
            
            # Read input JSON file
            with open(input_file_path, 'r', encoding='utf-8') as file:
                data = json.load(file)
            
            # Process each key in the JSON
            for key in list(data.keys()):
                # Check if key is numeric and 8 digits long
                if 1==1:
                    # Decompose key into book, chapter, verse
                    processed_data[key]={
                        **data[key],
                        "book": int(key[0:2]),
                        "chapter": int(key[2:5]),
                        "verse": int(key[5:8])
                    }
            
            # Write processed data to output file
            with open(output_file_path, 'w', encoding='utf-8') as outfile:
                json.dump(processed_data, outfile, ensure_ascii=False, indent=2)

# Example usage
input_folder = "D:\\Projects\\oper\\troika"
output_folder = "D:\\Projects\\oper\decomposed"
process_json_files(input_folder, output_folder)
