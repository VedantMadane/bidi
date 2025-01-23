import json
from pathlib import Path
from collections import defaultdict

def process_mbh_files():
    # Initialize speaker dictionary to store accumulated text
    speakers = defaultdict(lambda: {"ud": "", "ur": "", "ascii": ""})
    
    # Counter for renaming keys
    counter = 1
    key_mapping = {}  # To store original key to new number mapping
    
    # Process files MBh01.json through MBh18.json
    for book_num in range(1, 19):
        file_path = Path(f"D:/Projects/oper/vaktāsaha/MBh{book_num:02d}.json")
        
        if not file_path.exists():
            print(f"Warning: File {file_path} not found")
            continue
            
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                data = json.load(f)
                
            # Process each verse entry
            for verse_key, verse_data in data.items():
                # Skip if no vaktā (speaker) information
                if "vaktā" not in verse_data:
                    continue
                    
                speaker = verse_data["vaktā"]
                # Convert speaker dict to tuple for hashable key
                speaker_key = tuple(sorted([
                    (k, v.strip()) for k, v in speaker.items()
                ]))
                
                # If this is a new speaker, assign a new number
                if speaker_key not in key_mapping:
                    key_mapping[speaker_key] = str(counter)
                    counter += 1
                
                # Get the text content if it exists
                if "text" in verse_data:
                    text = verse_data["text"]
                    # Append text for each format
                    speakers[key_mapping[speaker_key]]["ud"] += text.get("ud", "")
                    speakers[key_mapping[speaker_key]]["ur"] += text.get("ur", "")
                    speakers[key_mapping[speaker_key]]["ascii"] += text.get("ascii", "")
                    
        except json.JSONDecodeError as e:
            print(f"Error reading {file_path}: {e}")
        except Exception as e:
            print(f"Unexpected error processing {file_path}: {e}")

    # Create reverse mapping for output
    speaker_info = {}
    for speaker_key, number in key_mapping.items():
        speaker_dict = dict(speaker_key)
        speaker_info[number] = {
            "speaker": speaker_dict,
            "text": speakers[number]
        }

    # Write the output file
    output_path = Path("D:/Projects/oper/vaktāsaha/by_speaker.json")
    try:
        with open(output_path, 'w', encoding='utf-8') as f:
            json.dump(speaker_info, f, ensure_ascii=False, indent=2)
        print(f"Successfully created {output_path}")
        print(f"Total speakers processed: {len(speaker_info)}")
    except Exception as e:
        print(f"Error writing output file: {e}")

if __name__ == "__main__":
    process_mbh_files()
