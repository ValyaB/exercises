package mycompression

func mycompression(layers ...string) []string {

	for i := range layers {
		layers[i], layers[0] = layers[0], layers[i]
	}
	return layers
}
