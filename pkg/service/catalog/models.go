package catalog

import (
	"github.com/Tlantic/go-util/v4/mrs"
	"github.com/Tlantic/mrs-service-superkoch-connector/proto"
)

type LabelModel int32

const (
	_ = iota
	LabelModelSingle
	LabelModelWholesale
	LabelModelWholesaleDiscount
	LabelModelSingleDiscount
	LabelModelBuyTake
	LabelModelDMCardDiscount
	LabelModelBuyWin
	LabelModelInstallment

	// LabelModelWholesaleVSRetailPlusDMCard
	//
	// This label will be used on products that have better wholesale
	// price than retail + DMCard promotional price
	LabelModelWholesaleVSRetailPlusDMCard

	LabelModelDepreciation
	LabelModelHundredGrams
	LabelModelBuyXWinYPercentageDiscount

	LabelModel13
	LabelModel14
	LabelModel15
)

type CabLabel struct {
	Sku           string `json:"sku"`
	Price         string `json:"price"`
	SalesQuantity string `json:"sales_quantity"`
	EAN           string `json:"ean"`
	ItemDesc      string `json:"item_desc"`
}

var labelModelName = map[int32]string{
	LabelModelSingle:            "1",
	LabelModelWholesale:         "2",
	LabelModelWholesaleDiscount: "3",
	LabelModelSingleDiscount:    "4",
	// LabelModelSingleDiscount:               "41",
	LabelModelBuyTake: "5",
	// LabelModelBuyTake:                      "51",
	LabelModelDMCardDiscount: "6",
	// LabelModelDMCardDiscount:               "61",
	LabelModelBuyWin: "7",
	// LabelModelBuyWin:                       "71",
	LabelModelInstallment: "8",
	// LabelModelInstallment:                  "81",
	LabelModelWholesaleVSRetailPlusDMCard: "9",
	// LabelModelWholesaleVSRetailPlusDMCard:  "91",
	LabelModelDepreciation:               "10",
	LabelModelHundredGrams:               "11",
	LabelModelBuyXWinYPercentageDiscount: "12",
	// LabelModelBuyXWinYPercentageDiscount:   "121",
	LabelModel13: "13",
	// LabelModel13:   "131",
	LabelModel14: "14",
	// LabelModel14:   "141",
	LabelModel15: "15",
	// LabelModel15:   "151",
}

var labelModelValue = map[string]int32{
	"1":   LabelModelSingle,
	"2":   LabelModelWholesale,
	"3":   LabelModelWholesaleDiscount,
	"4":   LabelModelSingleDiscount,
	"41":  LabelModelSingleDiscount,
	"5":   LabelModelBuyTake,
	"51":  LabelModelBuyTake,
	"6":   LabelModelDMCardDiscount,
	"61":  LabelModelDMCardDiscount,
	"7":   LabelModelBuyWin,
	"71":  LabelModelBuyWin,
	"8":   LabelModelInstallment,
	"81":  LabelModelInstallment,
	"9":   LabelModelWholesaleVSRetailPlusDMCard,
	"91":  LabelModelWholesaleVSRetailPlusDMCard,
	"10":  LabelModelDepreciation,
	"11":  LabelModelHundredGrams,
	"12":  LabelModelBuyXWinYPercentageDiscount,
	"121": LabelModelBuyXWinYPercentageDiscount,
	"13":  LabelModel13,
	"131": LabelModel13,
	"14":  LabelModel14,
	"141": LabelModel14,
	"15":  LabelModel15,
	"151": LabelModel15,
}

func (l LabelModel) String() string {
	return mrs.EnumName(labelModelName, int32(l))
}

func ToLabelModel(modelStr string) LabelModel {
	return LabelModel(mrs.EnumValue(labelModelValue, modelStr))
}

// ************************************ //

type ClientRequest struct {
	Store string `json:"lojaId"`
	Sku   string `json:"sku"`
}

type ClientStockPriceResponse struct {
	ServiceResponse ClientStockPriceResponseVariable `json:"response"`
}

type ClientStockPriceResponseVariable struct {
	ProductDetails *StockPriceProductDetail `json:"produto"`
}

type ClientLabelResponseBytes struct {
	ServiceResponse ClientLabelResponseVariable `json:"response"`
}

type ClientLabelResponseVariable struct {
	ProductDetails LabelProductBytes `json:"produto"`
}

type RequestProductDetail struct {
	Sku string `json:"sku"`
}

type StockPriceProductDetail struct {
	Codigo              string `json:":Codigo"`
	CodigoBarras        string `json:":CodigoBarras"`
	Ativo               string `json:":Ativo"`
	CodigoDepartamento  string `json:":CodigoDepartamento"`
	CodigoNCM           string `json:":CodigoNCM"`
	CodigoTributario    string `json:":CodigoTributario"`
	CodigoUnidade       string `json:":CodigoUnidade"`
	Complemento         string `json:":Complemento"`
	DataOferta          string `json:":DataOferta"`
	Departamento        string `json:":Departamento"`
	Descricao           string `json:":Descricao"`
	Estoque1            string `json:":Estoque1"`
	Estoque2            string `json:":Estoque2"`
	Estoque3            string `json:":Estoque3"`
	Estoque4            string `json:":Estoque4"`
	Estoque5            string `json:":Estoque5"`
	FatorPr3            string `json:":FatorPr3"`
	FatorPr4            string `json:":FatorPr4"`
	FatorPr5            string `json:":FatorPr5"`
	Grupo               string `json:":Grupo"`
	Marca               string `json:":Marca"`
	Oferta              string `json:":Oferta"`
	Preco               string `json:":Preco"`
	PrecoCartaz         string `json:":PrecoCartaz"`
	PrecoEtiqueta       string `json:":PrecoEtiqueta"`
	PrecoLista          string `json:":PrecoLista"`
	PrecoNormal         string `json:":PrecoNormal"`
	PrecoPDV            string `json:":PrecoPDV"`
	PrecoPDV3           string `json:":PrecoPDV3"`
	PrecoPDV4           string `json:":PrecoPDV4"`
	PrecoPDV5           string `json:":PrecoPDV5"`
	PrecoVenda3         string `json:":PrecoVenda3"`
	PrecoVenda4         string `json:":PrecoVenda4"`
	PrecoVenda5         string `json:":PrecoVenda5"`
	QuantidadeEmbalagem string `json:":QuantidadeEmbalagem"`
	SKU                 string `json:":SKU"`
	TipoEmbalagem       string `json:":TipoEmbalagem"`
}

type LabelProductBytes []byte

func (b *LabelProductBytes) UnmarshalJSON(input []byte) error {
	*b = input
	return nil
}

type LabelProductSimple struct {
	DynamicModel string `json:"dinamica"`
}

// ************************************ //

type StoresStockInfo struct {
	StoreID         string `json:"store_id,omitempty"`
	StoreName       string `json:"store_name,omitempty"`
	StockAfterSales string `json:"stock_after_sales,omitempty"`
}

func stockOtherStoresToProto(m []*StoresStockInfo) []*proto.StockResponse {
	if m == nil {
		return nil
	}

	list := make([]*proto.StockResponse, 0, len(m))

	for _, v := range m {
		list = append(list, &proto.StockResponse{
			StoreId:         v.StoreID,
			StoreName:       v.StoreName,
			StockAfterSales: v.StockAfterSales,
		})
	}
	return list
}
