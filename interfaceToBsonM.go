func toBsonM(v interface{}) (bson.M, error) {
	var doc bson.M
	data, err := bson.Marshal(v)

	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(data, &doc)

	return doc, nil
}
