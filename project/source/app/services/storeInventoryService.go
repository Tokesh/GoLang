package services

import "projectGoLang/source/domain/entity"

func (s Service) AddStoreInventoryService(inventory entity.StoreInventory) error {
	return s.AddStoreInventoryRepo(inventory)
}

func (s Service) SelectStoreInventoryService(storeId int, prodId int) (entity.StoreInventory, error) {
	return s.SelectStoreInventoryRepo(storeId, prodId)
}

func (s Service) DeleteStoreInventoryService(storeId int, prodId int) error {
	return s.DeleteStoryInventoryRepo(storeId, prodId)
}
func (s Service) UpdateStoreInventoryService(inventory entity.StoreInventory) error {
	return s.UpdateStoreInventoryRepo(inventory)
}
