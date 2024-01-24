package main

import "fmt"

func check(contacts map[string]string, name string) bool {
	for i, _ := range contacts {
		if i == name {
			return true
		}
	}
	return false
}
func addContact(contacts map[string]string, name string, number string) map[string]string {
	isExist := check(contacts, name)
	if isExist {
		fmt.Println("DATA IS EXIST")
		return contacts
	}
	contacts[name] = number
	return contacts
}
func updateContactNumber(contacts map[string]string, name string, number string) map[string]string {
	isExist := check(contacts, name)
	if isExist {
		contacts[name] = number
		return contacts
	}
	fmt.Println("DATA NOT EXIST")
	return contacts
}
func updateContactName(contacts map[string]string, oldName string, newName string) map[string]string {
	isExist := check(contacts, oldName)
	if isExist {
		number := contacts[oldName]
		deleteContact(contacts, oldName)
		addContact(contacts, newName, number)
		return contacts
	}
	fmt.Println("DATA NOT EXIST")
	return contacts
}
func deleteContact(contacts map[string]string, name string) map[string]string {
	isExist := check(contacts, name)
	if isExist {
		delete(contacts, name) //it shown error but it works
		return contacts
	}
	fmt.Println("DATA NOT EXIST")
	return contacts
}

func main() {
	// contact management system using hash table (map)

	contacts := map[string]string{}
	contacts["kira"] = "222-001"
	contacts["antok"] = "222-002"
	contacts["budi"] = "222-003"

	fmt.Println(contacts)
	res := addContact(contacts, "mitoma", "222-004")
	fmt.Println(res)

	res = updateContactNumber(contacts, "antok", "222-005")
	fmt.Println(res)

	res = updateContactName(contacts, "kira", "new-kira")
	fmt.Println(res)

	res = deleteContact(contacts, "antok")
	fmt.Println(res)
}
