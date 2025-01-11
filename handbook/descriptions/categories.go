package descriptions

import _const "github.com/TrashPony/veliri-lib/const"

var CategoriesDescription = map[string]map[string]DescriptionItem{
	_const.EN: {
		"weapon":     {Name: "weapon", Description: ""},
		"ammo":       {Name: "ammunition", Description: ""},
		"equip":      {Name: "equipment", Description: ""},
		"body":       {Name: "body", Description: ""},
		"resource":   {Name: "mineral resource", Description: ""},
		"recycle":    {Name: "raw material", Description: ""},
		"detail":     {Name: "detail", Description: ""},
		"boxes":      {Name: "box", Description: ""},
		"trash":      {Name: "trash", Description: ""},
		"skill":      {Name: "-", Description: ""},
		"structure":  {Name: "structure", Description: ""},
		"blueprints": {Name: "blueprint", Description: ""},
		"sector":     {Name: "sector", Description: ""},
		"product":    {Name: "product", Description: ""},
		"fuel":       {Name: "fuel", Description: ""},
		"mine_drone": {Name: "mine drone", Description: ""},
	},
	_const.RU: {
		"weapon":     {Name: "оружие", Description: ""},
		"ammo":       {Name: "боеприпас", Description: ""},
		"equip":      {Name: "снаряжение", Description: ""},
		"body":       {Name: "корпус", Description: ""},
		"resource":   {Name: "полезное ископаемое", Description: ""},
		"recycle":    {Name: "сырье", Description: ""},
		"detail":     {Name: "деталь", Description: ""},
		"boxes":      {Name: "ящик", Description: ""},
		"trash":      {Name: "хлам", Description: ""},
		"skill":      {Name: "-", Description: ""},
		"structure":  {Name: "строение", Description: ""},
		"blueprints": {Name: "чертеж", Description: ""},
		"sector":     {Name: "сектор", Description: ""},
		"product":    {Name: "товар", Description: ""},
		"fuel":       {Name: "топливо", Description: ""},
		"mine_drone": {Name: "дрон шахтер", Description: ""},
	},
}
