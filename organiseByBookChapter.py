import os
import json
from typing import Dict, Any

def organize_json_files(input_folder, output_folder):
    # Create output directories
    os.makedirs(output_folder, exist_ok=True)

    # Process each JSON file in the input folder
    for filename in os.listdir(input_folder):
        if filename.endswith('.json'):
            input_path = os.path.join(input_folder, filename)
            
            # Read input JSON
            with open(input_path, 'r', encoding='utf-8') as f:
                data = json.load(f)
            
            # Reorganized data structure
            organized_data: Dict[str, Dict[str, Dict[str, Any]]] = {}

            # Process each entry
            for key, entry in data.items():
                # Extract book, chapter, verse from key
                book = key[0:2]
                chapter = key[2:5]
                verse = key[5:8]


                # Initialize nested dictionaries
                if book not in organized_data:
                    organized_data[book] = {}
                if chapter not in organized_data[book]:
                    organized_data[book][chapter] = {}
                
                # Add entry to organized structure
                organized_data[book][chapter][verse] = entry

            # Write organized data to new files
            for book, chapters in organized_data.items():
                book_folder = os.path.join(output_folder, f"Book{book:02d}")
                os.makedirs(book_folder, exist_ok=True)
                
                for chapter, verses in chapters.items():
                    chapter_file = os.path.join(book_folder, f"Chapter{chapter:03d}.json")
                    
                    with open(chapter_file, 'w', encoding='utf-8') as f:
                        json.dump(verses, f, ensure_ascii=False, indent=2)

# Main execution
input_folder = 'D:\Projects\oper\decomposed'
output_folder = 'D:\Projects\oper\\organized'

organize_json_files(input_folder, output_folder)
