package reflexao

func percorre(x interface{}, fn func(entrada string)){
	fn("eu gosto de pastel de feira")
}