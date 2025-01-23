def process_json(input_json):
    # Convert JSON to dictionary
    data = json.loads(input_json)
    
    # Track the last speaker
    last_speaker = None
    
    # Iterate through keys in order
    for key in sorted(data.keys()):
        # Check if key has no alphabetic suffix
        if key.isnumeric():
            last_speaker = data[key]['text']['ur']
        
        # If key has alphabetic suffix and no speaker
        if not key.isnumeric() and 'speaker' not in data[key]:
            data[key]['speaker'] = last_speaker
    
    return json.dumps(data, ensure_ascii=False, indent=2)
