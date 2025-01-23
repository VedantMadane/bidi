import requests
import json
import os
from bs4 import BeautifulSoup

def web_scrape_sacred_texts(book_num, chapter_num):
    # Construct URL
    url = f"https://sacred-texts.com/hin/m{book_num:02d}/m{book_num:02d}{chapter_num:03d}.htm"
    
    try:
        # Fetch webpage
        response = requests.get(url)
        response.raise_for_status()
        
        # Parse HTML
        soup = BeautifulSoup(response.text, 'html.parser')
        
        # Extract text (adjust based on actual page structure)
        text_content = soup.get_text()
        
        return text_content
    
    except requests.RequestException as e:
        print(f"Error scraping {url}: {e}")
        return None

def update_json_with_translation(input_folder, output_folder):
    # Ensure output folder exists
    os.makedirs(output_folder, exist_ok=True)
    
    # Iterate through books
    for book_num in range(1, 19):
        input_file = os.path.join(input_folder, f"MBh{book_num:02d}.json")
        output_file = os.path.join(output_folder, f"MBh{book_num:02d}.json")
        
        # Read input JSON
        with open(input_file, 'r', encoding='utf-8') as f:
            data = json.load(f)
        
        # Track processed keys to avoid duplicates
        processed_keys = set()
        
        # Iterate through entries
        for key, entry in data.items():
            # Check if key matches 8-digit format
            if len(key) == 8 and key.startswith(f"{book_num:02d}"):
                # Extract chapter from key
                chapter = int(key[2:5])
                
                # Scrape translation if not processed before
                if key not in processed_keys:
                    translation = web_scrape_sacred_texts(book_num, chapter)
                    
                    if translation:
                        # Add translation to entry
                        entry['translation'] = {
                            'en': translation
                        }
                        
                        # Mark as processed
                        processed_keys.add(key)
        
        # Write updated JSON
        with open(output_file, 'w', encoding='utf-8') as f:
            json.dump(data, f, ensure_ascii=False, indent=2)

# Main execution
input_folder = "D:\Projects\oper\\vaktāsaha"
output_folder = "D:\Projects\oper\\translations"
update_json_with_translation(input_folder, output_folder)
