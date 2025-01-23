import os
import json

def extract_chapter_key(chapter):
    """Extract chapter key from the chapter identifier."""
    return chapter[2:8] if len(chapter) >= 8 else chapter

def process_mbh_files():
    source_dir = 'D:\Projects\oper\MBh'
    output_dir = os.path.join(source_dir, 'processed_chapters')
    os.makedirs(output_dir, exist_ok=True)

    # Process each MBh JSON file
    for book_num in range(1, 19):
        filename = f'MBh{book_num:02d}.json'
        filepath = os.path.join(source_dir, filename)
        
        try:
            with open(filepath, 'r', encoding='utf-8') as f:
                # Use json.load() to parse the entire file
                data = json.load(f)
                
                # Group chapters by their chapter key
                chapter_groups = {}
                for chapter in data:
                    chapter_key = extract_chapter_key(chapter.get('chapter', ''))
                    if chapter_key not in chapter_groups:
                        chapter_groups[chapter_key] = []
                    chapter_groups[chapter_key].append(chapter)
                
                # Create book-specific directory
                book_dir = os.path.join(output_dir, f'Book_{book_num:02d}')
                os.makedirs(book_dir, exist_ok=True)
                
                # Write chapter files
                for chapter_key, chapter_data in chapter_groups.items():
                    chapter_filepath = os.path.join(book_dir, f'{chapter_key}.json')
                    with open(chapter_filepath, 'w', encoding='utf-8') as out_file:
                        json.dump(chapter_data, out_file, indent=2, ensure_ascii=False)
                
                print(f"Processed {filename}")
        
        except Exception as e:
            print(f"Error processing {filename}: {e}")

if __name__ == '__main__':
    process_mbh_files()


