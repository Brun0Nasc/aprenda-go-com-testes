package maps

import "errors"

type Dicionario map[string]string

var (
	ErrNaoEncontrado    = errors.New("não foi possivel encontrar a palavra que voce procura")
	ErrPalavraNaoExistente = errors.New("não é possível adicionar essa palavra pois ela já existe")
)

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra] //* o segundo valor recebe um boolean que informa se a palavra informada existe ou não
	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraNaoExistente
	default:
		return err
	}

	return nil
}
