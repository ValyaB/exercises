package mycompression

func mycompression(layers ...interface{}) []interface{} {

	for i := range layers {
		layers[i], layers[0] = layers[0], layers[i]
	}
	return layers
}
