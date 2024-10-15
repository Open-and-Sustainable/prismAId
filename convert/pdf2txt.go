package convert

import (
	
	pdf "github.com/ledongthuc/pdf"
)

func readPdf(path string) (string, error) {
	text := ""
	// Open the PDF file
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()
	if totalPage == 0 {
		return "", nil
	}

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			text += textsToString(row.Content) + "\n"
		}
	}
	return text, nil
}

// Convert a []Text to a single string by concatenating the Value fields
func textsToString(texts []pdf.Text) string {
    result := ""
    for _, text := range texts {
        result += text.S
    }
    return result
}