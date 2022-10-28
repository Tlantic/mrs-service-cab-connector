package catalog

import "encoding/json"

type ClientStockPriceResponse struct {
	ServiceResponse ClientStockPriceResponseVariable `json:"response"`
}

type ClientStockPriceResponseVariable struct {
	ProductDetails *StockPriceProductDetail `json:"produto"`
}

type StockPriceProductDetail struct {
	Codigo              json.Number `json:"Codigo"`
	CodigoBarras        string      `json:"CodigoBarras"`
	Ativo               string      `json:"Ativo"`
	CodigoDepartamento  json.Number `json:"CodigoDepartamento"`
	CodigoNCM           string      `json:"CodigoNCM"`
	CodigoTributario    string      `json:"CodigoTributario"`
	CodigoUnidade       string      `json:"CodigoUnidade"`
	Complemento         string      `json:"Complemento"`
	DataOferta          string      `json:"DataOferta"`
	Departamento        string      `json:"Departamento"`
	Descricao           string      `json:"Descricao"`
	Estoque1            json.Number `json:"Estoque1"`
	Estoque2            json.Number `json:"Estoque2"`
	Estoque3            json.Number `json:"Estoque3"`
	Estoque4            json.Number `json:"Estoque4"`
	Estoque5            json.Number `json:"Estoque5"`
	FatorPr3            json.Number `json:"FatorPr3"`
	FatorPr4            json.Number `json:"FatorPr4"`
	FatorPr5            json.Number `json:"FatorPr5"`
	Grupo               string      `json:"Grupo"`
	Marca               string      `json:"Marca"`
	Oferta              string      `json:"Oferta"`
	Preco               json.Number `json:"Preco"`
	PrecoCartaz         json.Number `json:"PrecoCartaz"`
	PrecoEtiqueta       json.Number `json:"PrecoEtiqueta"`
	PrecoLista          json.Number `json:"PrecoLista"`
	PrecoNormal         json.Number `json:"PrecoNormal"`
	PrecoPDV            json.Number `json:"PrecoPDV"`
	PrecoPDV3           json.Number `json:"PrecoPDV3"`
	PrecoPDV4           json.Number `json:"PrecoPDV4"`
	PrecoPDV5           json.Number `json:"PrecoPDV5"`
	PrecoVenda3         json.Number `json:"PrecoVenda3"`
	PrecoVenda4         json.Number `json:"PrecoVenda4"`
	PrecoVenda5         json.Number `json:"PrecoVenda5"`
	QuantidadeEmbalagem json.Number `json:"QuantidadeEmbalagem"`
	SKU                 json.Number `json:"SKU"`
	TipoEmbalagem       string      `json:"TipoEmbalagem"`
}

type ClientLabelResponseBytes struct {
	ServiceResponse ClientLabelResponseVariable `json:"response"`
}

type ClientLabelResponseVariable struct {
	ProductDetails LabelProductBytes `json:"produto"`
}

type LabelProductBytes []byte

func (b *LabelProductBytes) UnmarshalJSON(input []byte) error {
	*b = input
	return nil
}
