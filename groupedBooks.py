
import os
import json
from collections import defaultdict

def group_by_book(directory):
    # Create output directory for grouped books
    output_dir = os.path.join(directory, "grouped_chapters1")
    os.makedirs(output_dir, exist_ok=True)
    # Dictionary to store books grouped by book number
    book_groups = defaultdict(dict)
    # Dictionary to store chapters grouped by book number
    chapter_groups = defaultdict(dict)

    # Iterate through MBh JSON files
    for filename in os.listdir(directory):
        if filename.startswith("MBh") and filename.endswith(".json"):
            filepath = os.path.join(directory, filename)
            
            # Read JSON file
            with open(filepath, "r", encoding="utf-8") as f:
                file_data = json.load(f)
            
            # Group entries by book
            for key, entry in file_data.items():
                book_number = entry.get("book")
                for ke, entry in file_data.items():
                    if entry.get("book") == book_number:
                        entry["chapter"] = ke
                chapter_number = entry.get("chapter")
                book_groups[book_number][key] = entry
                chapter_groups[chapter_number][key] = entry
    # Write grouped books to separate JSON files
    for book_number, book_entries in book_groups.items():
        chapter_numbers = sorted(set(entry["chapter"] for entry in book_entries.values()))
        chapter_entries = {chapter_number: chapter_groups[chapter_number] for chapter_number in chapter_numbers}
        for chapter_number, chapter_entries in chapter_entries.items():
            output_filepath = os.path.join(output_dir, f"Book_{book_number:02d}_Chapter_{chapter_number}.json")
        
        with open(output_filepath, "w", encoding="utf-8") as f:
            json.dump(book_entries, f, ensure_ascii=False, indent=2)

    print(f"Grouped {len(book_groups)} unique books")

# Execute the function
group_by_book(r"D:\\Projects\\oper\\decomposed")
