// Package convert provides utilities to convert various document formats (PDF, DOCX, HTML) into plain text format.
// It exposes functions to process and extract textual content from these document types.
//
// Overview
//
// The `convert` package is designed to convert a variety of document formats into plain text.
// It supports the following formats:
//   - PDF: Extracts text from PDF files using the `github.com/ledongthuc/pdf` library.
//   - DOCX: Converts DOCX files into plain text using the `github.com/fumiama/go-docx` library.
//   - HTML: Strips HTML tags and extracts textual content using the `jaytaylor.com/html2text` package.
//
// Exported Functions
//
// Convert: Converts all supported document files from the input directory to plain text files based on the configuration settings.
//
// Example:
//    > err := convert.Convert(config)
//    > if err != nil {
//    >     log.Fatalf("Conversion failed: %v", err)
//    > }

package convert

