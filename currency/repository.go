package currency

import (
	"errors"
	"strconv"
)

type CurrencySymbol string

type Currency struct {
	Symbol        CurrencySymbol
	DecimalPlaces int
	Name          string
}

type CurrencyRepository struct {
	currencies []*Currency
	symbolMap  map[string]*Currency
}

func ProvideCurrencyRepository() *CurrencyRepository {
	currencies := []*Currency{}
	for _, d := range currencyData {
		var v int
		var err error
		if d[2] == "." {
			v = 0
		} else {
			v, err = strconv.Atoi(d[2])
			if err != nil {
				panic(errors.New("currency initialization failure"))
			}

		}
		currencies = append(currencies, &Currency{
			Symbol:        CurrencySymbol(d[0]),
			DecimalPlaces: v,
			Name:          d[3],
		})
	}
	var symbolMap = make(map[string]*Currency)
	for _, c := range currencies {
		symbolMap[string(c.Symbol)] = c
	}

	return &CurrencyRepository{
		currencies: currencies,
		symbolMap:  symbolMap,
	}
}

func (cm *CurrencyRepository) GetBySymbol(cc string) (Currency, error) {
	c, ok := cm.symbolMap[cc]
	if !ok {
		return Currency{}, errors.New("invalid currency symbol")
	}

	return *c, nil
}

func (cm *CurrencyRepository) CurrencyList() []Currency {
	m := make([]Currency, len(cm.currencies))
	for i, c := range cm.currencies {
		m[i] = *c
	}

	return m
}
