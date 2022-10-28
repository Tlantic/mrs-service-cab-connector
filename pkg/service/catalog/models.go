package catalog

type ClientStockPriceResponse struct {
	ServiceResponse ClientStockPriceResponseVariable `json:"response"`
}

type ClientStockPriceResponseVariable struct {
	ProductDetails *StockPriceProductDetail `json:"produto"`
}

type StockPriceProductDetail struct {
	Codigo              string `json:"Codigo"`
	CodigoBarras        string `json:"CodigoBarras"`
	Ativo               string `json:"Ativo"`
	CodigoDepartamento  string `json:"CodigoDepartamento"`
	CodigoNCM           string `json:"CodigoNCM"`
	CodigoTributario    string `json:"CodigoTributario"`
	CodigoUnidade       string `json:"CodigoUnidade"`
	Complemento         string `json:"Complemento"`
	DataOferta          string `json:"DataOferta"`
	Departamento        string `json:"Departamento"`
	Descricao           string `json:"Descricao"`
	Estoque1            string `json:"Estoque1"`
	Estoque2            string `json:"Estoque2"`
	Estoque3            string `json:"Estoque3"`
	Estoque4            string `json:"Estoque4"`
	Estoque5            string `json:"Estoque5"`
	FatorPr3            string `json:"FatorPr3"`
	FatorPr4            string `json:"FatorPr4"`
	FatorPr5            string `json:"FatorPr5"`
	Grupo               string `json:"Grupo"`
	Marca               string `json:"Marca"`
	Oferta              string `json:"Oferta"`
	Preco               string `json:"Preco"`
	PrecoCartaz         string `json:"PrecoCartaz"`
	PrecoEtiqueta       string `json:"PrecoEtiqueta"`
	PrecoLista          string `json:"PrecoLista"`
	PrecoNormal         string `json:"PrecoNormal"`
	PrecoPDV            string `json:"PrecoPDV"`
	PrecoPDV3           string `json:"PrecoPDV3"`
	PrecoPDV4           string `json:"PrecoPDV4"`
	PrecoPDV5           string `json:"PrecoPDV5"`
	PrecoVenda3         string `json:"PrecoVenda3"`
	PrecoVenda4         string `json:"PrecoVenda4"`
	PrecoVenda5         string `json:"PrecoVenda5"`
	QuantidadeEmbalagem string `json:"QuantidadeEmbalagem"`
	SKU                 string `json:"SKU"`
	TipoEmbalagem       string `json:"TipoEmbalagem"`
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
