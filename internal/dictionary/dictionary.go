package dictionary

import "errors"

func GetDefinition(word string) (string, error) {
	definitions := map[string]string{
		"apple":  "A fruit that grows on trees and is typically red or green",
		"banana": "A long curved fruit that grows in clusters",
		"orange": "A citrus fruit with a tough bright reddish-yellow rind",
	}

	definition, ok := definitions[word]
	if !ok {
		return "", errors.New("Definition not found")
	}

	return definition, nil
}
