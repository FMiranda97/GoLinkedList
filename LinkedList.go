package main

import (
	"errors"
	"fmt"
	"reflect"
)

func comparator(this interface{}, target interface{}) (int8, error) {
	if reflect.TypeOf(this) != reflect.TypeOf(target) {
		return 0, errors.New("compared parameters implement different interfaces")
	}
	switch this.(type) {
	case linkedListCargoString:
		a := this.(linkedListCargoString)
		b := target.(linkedListCargoString)
		if a.getKey() < b.getKey() {
			return -1, nil
		} else if a.getKey() > b.getKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case linkedListCargoInt:
		a := this.(linkedListCargoInt)
		b := target.(linkedListCargoInt)
		if a.getKey() < b.getKey() {
			return -1, nil
		} else if a.getKey() > b.getKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case linkedListCargoInt64:
		a := this.(linkedListCargoInt)
		b := target.(linkedListCargoInt)
		if a.getKey() < b.getKey() {
			return -1, nil
		} else if a.getKey() > b.getKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New("compared parameters implement unsupported interface")
	}
}

func implementsLinkedListCargo(cargo interface{}) bool {
	_, ok := cargo.(linkedListCargoString)
	_, ok2 := cargo.(linkedListCargoInt)
	return ok || ok2
}

func insertList(list *linkedList, cargo interface{}) error {
	// check viability
	if *list != nil && reflect.TypeOf((*list).cargo) != reflect.TypeOf(cargo) {
		return errors.New("cargo is not of same type as contents previously inserted")
	}
	if !implementsLinkedListCargo(cargo) {
		return errors.New("cargo does not implement a supported interface")
	}
	var newNode node
	newNode.cargo = cargo
	//if insert at beginning
	if *list == nil {
		*list = &newNode
		return nil
	} else if comp, err := comparator((*list).cargo, newNode.cargo); comp > 0 {
		if err != nil {
			newNode.next = *list
			*list = &newNode
			return nil
		} else {
			return err
		}
	}
	//insert in order
	var i linkedList
	for i = *list; i.next != nil; i = i.next {
		comp, err := comparator(i.next.cargo, newNode.cargo)
		if err != nil {
			return err
		}
		if comp > 0 {
			break
		}
	}
	newNode.next = i.next
	i.next = &newNode
	return nil
}

func printList(list linkedList) {
	fmt.Println("----------------------")
	for i := list; i != nil; i = i.next {
		fmt.Println(*i)
	}
	fmt.Println("----------------------")
}