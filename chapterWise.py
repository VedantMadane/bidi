import os
import json

def group_by_chapters(input_folder, output_base_folder):
    # Create base chapter-wise folder
    os.makedirs(output_base_folder, exist_ok=True)

    # Iterate through books
    for book_num in range(1, 19):
        # Create book-specific chapter folder
        book_chapter_folder = os.path.join(output_base_folder, f"Book{book_num:02d}")
        os.makedirs(book_chapter_folder, exist_ok=True)

        # Input file path
        input_file = os.path.join(input_folder, f"MBh{book_num:02d}.json")

        # Read input JSON
        with open(input_file, 'r', encoding='utf-8') as f:
            book_data = json.load(f)

        # Dictionary to store chapter-wise grouping
        chapter_groups = {}

        # Group by chapters
        for key, entry in book_data.items():
            # Check if key matches book number and has chapter information
            if len(key) == 8 and key.startswith(f"{book_num:02d}"):
                chapter = key[2:5]
                
                # Initialize chapter group if not exists
                if chapter not in chapter_groups:
                    chapter_groups[chapter] = {}
                
                # Add entry to chapter group
                chapter_groups[chapter][key] = entry

        # Write chapter-wise JSON files
        for chapter, chapter_entries in chapter_groups.items():
            output_file = os.path.join(book_chapter_folder, f"{chapter}.json")
            with open(output_file, 'w', encoding='utf-8') as f:
                json.dump(chapter_entries, f, ensure_ascii=False, indent=2)

# Main execution
input_folder = 'D:\Projects\oper\\vaktāsaha'
output_folder = 'D:\Projects\oper\chapterWise'
group_by_chapters(input_folder, output_folder)
