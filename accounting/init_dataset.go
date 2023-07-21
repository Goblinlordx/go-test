package accounting

import (
	"github.com/google/uuid"
)

func GetInitDatasetCategories() []Category {
	return []Category{
		{
			Id:    uuid.MustParse("47d2d766-fb85-4d08-876f-ba3596521de5"),
			Value: "Uncategorized",
		},
	}
}

type initDataset struct {
	Categories []Category
}

func GetInitDataset() initDataset {
	return initDataset{
		Categories: GetInitDatasetCategories(),
	}
}
