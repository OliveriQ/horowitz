package main

import (
	"fmt"
)

func main() {
	fmt.Print("Go chess engine\n\n")
	initialize_lookup_tables()

	parse_fen(tricky_position, 0)

	print_board()
	var tt bool = is_square_attacked(SQ_B5, WHITE)
	fmt.Println(tt)
}