import os
import json
import re

def extract_chapter_number(key):
    """Extract chapter number from the key format."""
    match = re.match(r'(\d{2})001(\d{4})', key)
    return match.group(0) if match else key

def process_mbh_files():
    source_dir = r'D:\Projects\oper\MBh'
    output_dir = os.path.join(source_dir, 'processed_chapters')
    os.makedirs(output_dir, exist_ok=True)

    # Process each MBh JSON file
    for book_num in range(1, 19):
        filename = f'MBh{book_num:02d}.json'
        filepath = os.path.join(source_dir, filename)
        
        # Skip if file doesn't exist
        if not os.path.exists(filepath):
            print(f"File {filename} not found. Skipping.")
            continue

        # Read JSON file
        with open(filepath, 'r', encoding='utf-8') as f:
            data = json.load(f)

        # Group chapters by chapter number
        chapter_groups = {}
        for key, value in data.items():
            chapter_num = extract_chapter_number(key)
            if chapter_num not in chapter_groups:
                chapter_groups[chapter_num] = {}
            chapter_groups[chapter_num][key] = value

        # Create book-specific directory
        book_dir = os.path.join(output_dir, f'Book_{book_num:02d}')
        os.makedirs(book_dir, exist_ok=True)

        # Write chapter files
        for chapter_num, chapter_data in chapter_groups.items():
            chapter_file = os.path.join(book_dir, f'{chapter_num}.json')
            with open(chapter_file, 'w', encoding='utf-8') as f:
                json.dump(chapter_data, f, ensure_ascii=False, indent=2)

        print(f"Processed {filename}: {len(chapter_groups)} chapters")

if __name__ == '__main__':
    process_mbh_files()
