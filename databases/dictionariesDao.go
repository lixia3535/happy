package databases

import "happy/models"

func AddDictionaries(dictionaries models.Dictionaries) error {
	_, err := DBClient.Insert(dictionaries)
	if err != nil {
		return err
	}
	return nil
}

func deleteDictionaries(address models.Dictionaries)  {

}

func editDictionaries(address models.Dictionaries)  {

}

func findAllDictionaries(userID int)  {

}

func findDictionariesById(id int)  {

}
