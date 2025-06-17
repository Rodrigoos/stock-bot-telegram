package models

import "gorm.io/gorm"

type Portfolio struct {
  gorm.Model
  Name   string
  Assets []Asset
}

func (p *Portfolio) TotalValue() float64 {
  var total float64

  for _, asset := range p.Assets {
    total += float64(asset.Quantity) * asset.Price
  }

  return total
}
