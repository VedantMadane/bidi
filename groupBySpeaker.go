import os
import json
from typing import Dict, List

def group_by_speaker(base_path: str) -> Dict[str, List[dict]]:
    # Dictionary to store grouped speakers
    speaker_groups = {}

    # Iterate through Mahabharata books
    for book_num in range(1, 19):
        filename = f"MBh{book_num:02d}.json"
        filepath = os.path.join(base_path, filename)

        # Read JSON file
        with open(filepath, 'r', encoding="utf-8") as f:
            book_data = json.load(f)

        # Group by speaker
        for key, entry in book_data.items():
            # Check if entry has a vaktā (speaker)
            if "vaktā" in entry:
                speaker = entry["vaktā"]["ur"]
                
                # Initialize speaker group if not exists
                if speaker not in speaker_groups:
                    speaker_groups[speaker] = []
                
                # Add entry to speaker's group
                speaker_groups[speaker].append(entry)

    return speaker_groups

# Main execution
base_path = 'D:\Projects\oper\\vaktāsaha'
grouped_speakers = group_by_speaker(base_path)

# Optional: Print summary
for speaker, entries in grouped_speakers.items():
    print(f"Speaker: {speaker}")
    print(f"Total verses: {len(entries)}")
