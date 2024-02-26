package handlers

import (
	"strings"
)

func IsValidScore(minScore float64) bool {
	if minScore >= 0 && minScore <= 10 {
		return true
	}
	return false
}

func IsValidSportClass(requestedClass string, existingClasses []string) bool {
	// Convert both lists to lowercase for case-insensitive comparison
	requestedClassLower := strings.ToLower(requestedClass)
	existingClassesLower := []string{}
	for _, class := range existingClasses {
		existingClassesLower = append(existingClassesLower, strings.ToLower(class))
	}

	// Check if the user sport is present in the valid sports list
	return contains(existingClassesLower, requestedClassLower)
}

func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
