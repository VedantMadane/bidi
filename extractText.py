# Write a python program to extract text from file:///D:/Downloads/dokumen.tips_maitrayani-samhita-of-yajurveda.pdf using tesseract-ocr
def extract_text_from_pdf(pdf_path):
    from pdf2image import convert_from_path
    from pytesseract import image_to_string

    # Convert PDF to images
    images = convert_from_path(pdf_path)

    # Extract text from each image
    text = ""
    for image in images:
        text += image_to_string(image)

    return text

extract_text_from_pdf(input("Enter the path to the PDF file: "))