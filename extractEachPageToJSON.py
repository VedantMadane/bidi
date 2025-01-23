import fitz  # PyMuPDF
import json

def pdf_to_json(pdf_path, json_path):
    # Open the PDF file
    pdf_document = fitz.open(pdf_path)
    pdf_data = []

    # Iterate through each page
    for page_num in range(len(pdf_document)):
        page = pdf_document.load_page(page_num)
        page_text = page.get_text("text")
        pdf_data.append({
            "page": page_num + 1,
            "text": page_text
        })

    # Write the data to a JSON file
    with open(json_path, 'w', encoding='utf-8') as json_file:
        json.dump(pdf_data, json_file, ensure_ascii=False, indent=4)

# Example usage
pdf_to_json("""https://www.sadagopan.org/ebook/pdf/Raghava%20Yadaveeyam.pdf""", "output.json")