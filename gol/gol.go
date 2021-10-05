package main

func mod(a, b int) int {
	return (a % b + b) % b
}


func checkSurrounding(i int, z int, neww [][]byte) int {
	x := 0
	if neww[mod(i-1,16)][z] == 255 {x++}
	if neww[mod(i+1,16)][z] == 255 {x++}
	if neww[i][mod(z+1,16)] == 255 {x++}
	if neww[i][mod(z-1,16)] == 255 {x++}
	if neww[mod(i-1,16)][mod(z+1,16)] == 255 {x++}
	if neww[mod(i-1,16)][mod(z-1,16)] == 255 {x++}
	if neww[mod(i+1,16)][mod(z+1,16)] == 255 {x++}
	if neww[mod(i+1,16)][mod(z-1,16)] == 255 {x++}
	return x
}


func calculateNextState(p golParams, world [][]byte) [][]byte {
	neww := make([][]byte, p.imageHeight)
	for i := range neww {
		neww[i] = make([]byte, p.imageWidth)
		copy(neww[i], world[i][:])
	}
	h := p.imageHeight
	w := p.imageWidth

	for i:=0; i<h; i++ {
		for z:=0; z<w; z++ {
			alive := checkSurrounding(i,z,world)
			if world[i][z] == 0 && alive==3 {neww[i][z] = 255
			} else {
				if world[i][z] == 255 && (alive<2 || alive>3) {neww[i][z] = 0}
			}
		}
	}
	return neww
}


func calculateAliveCells(p golParams, world [][]byte) []cell {
	alive := make([]cell,0)
	height := p.imageHeight
	width := p.imageWidth

	for i:=0; i<height; i++ {
		for z:=0; z<width; z++ {
			if world[i][z]==255 {
				var x cell
				x.x = z
				x.y = i
				alive = append(alive, x)
			}
		}
	}
	return alive
}