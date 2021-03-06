package cache

import "sync"

var (
	FIO          sync.Map
	Address      sync.Map
	Organization sync.Map
)

func LoadCache() {
	FIO.Store("И", []string{"Ирина", "Иван", "Игорь", "Иванова", "Иванов"})
	FIO.Store("Ив", []string{"Ивар", "Иван", "Иваненко", "Иванова", "Иванов"})
	FIO.Store("Ива", []string{"Ивар", "Иван", "Иваненко", "Иванова", "Иванов"})
	FIO.Store("Иван", []string{"Иванидзе", "Иван", "Иваненко", "Иванова", "Иванов"})
	FIO.Store("Ивано", []string{"Иванович", "Ивановна", "Иванонидзе", "Иванова", "Иванов"})
	FIO.Store("Иванов", []string{"Иванович", "Ивановна", "Иванонидзе", "Иванова", "Иванов"})
	Address.Store("К", []string{"Коровия вал 1", "Коровия вал 2", "Коровия вал 3", "Коровия вал 4", "Коровия вал 5"})
	Address.Store("Ко", []string{"Коровия вал 1", "Коровия вал 2", "Коровия вал 3", "Коровия вал 4", "Коровия вал 5"})
	Address.Store("Кор", []string{"Коровия вал 1", "Коровия вал 2", "Коровия вал 3", "Коровия вал 4", "Коровия вал 5"})
	Address.Store("Коро", []string{"Коровия вал 1", "Коровия вал 2", "Коровия вал 3", "Коровия вал 4", "Коровия вал 5"})
	Address.Store("Коров", []string{"Коровия вал 1", "Коровия вал 2", "Коровия вал 3", "Коровия вал 4", "Коровия вал 5"})
	Address.Store("Корови", []string{"Коровия вал 1", "Коровия вал 2", "Коровия вал 3", "Коровия вал 4", "Коровия вал 5"})
	Address.Store("Коровий", []string{"Коровия вал 1", "Коровия вал 2", "Коровия вал 3", "Коровия вал 4", "Коровия вал 5"})
	Organization.Store("Г", []string{"Газпром", "Газпромбанк", "Газпром Медия", "Газпром нефть", "Газпром европа"})
	Organization.Store("Га", []string{"Газпром", "Газпромбанк", "Газпром Медия", "Газпром нефть", "Газпром европа"})
	Organization.Store("Газ", []string{"Газпром", "Газпромбанк", "Газпром Медия", "Газпром нефть", "Газпром европа"})
	Organization.Store("Газп", []string{"Газпром", "Газпромбанк", "Газпром Медия", "Газпром нефть", "Газпром европа"})
	Organization.Store("Газпр", []string{"Газпром", "Газпромбанк", "Газпром Медия", "Газпром нефть", "Газпром европа"})
	Organization.Store("Газпро", []string{"Газпром", "Газпромбанк", "Газпром Медия", "Газпром нефть", "Газпром европа"})
	Organization.Store("Газпром", []string{"Газпром", "Газпромбанк", "Газпром Медия", "Газпром нефть", "Газпром европа"})
}
