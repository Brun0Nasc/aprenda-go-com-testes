package maps

import "errors"

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (string, error){
	definicao, existe := d[palavra] //* o segundo valor recebe um boolean que informa se a palavra informada existe ou não
	if !existe {
		return "", errors.New("não foi possivel encontrar a palavra que voce procura")
	}

	return definicao, nil
}