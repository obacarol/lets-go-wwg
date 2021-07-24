package calculadora

import "testing"

func TestAdicione(t *testing.T) {
	t.Run("a soma de 3 e 4 é igual a 7", func(t *testing.T) {
		obtive := Adicione(3, 4)
		espero := 7
		if obtive != espero {
			t.Errorf("espero '%d', mas obtive '%d'", espero, obtive)
		}
	})
	t.Run("a soma de 27 e 113 é igual a 140", func(t *testing.T) {
		obtive := Adicione(27, 113)
		espero := 140
		if obtive != espero {
			t.Errorf("espero '%d', mas obtive '%d'", espero, obtive)
		}
	})
}

func TestMultiplique(t *testing.T) {
	testes := []struct {
		nome     string
		x        int
		y        int
		esperado int
	}{
		{
			nome:     "7 x 5 = 35",
			x:        7,
			y:        5,
			esperado: 35,
		},
		{
			nome:     "3 x 20 = 60",
			x:        3,
			y:        20,
			esperado: 60,
		},
	}

	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			obtive := Multiplique(teste.x, teste.y)

			if obtive != teste.esperado {
				t.Errorf("espero '%d', mas obtive '%d'", teste.esperado, obtive)
			}
		})
	}

	/*t.Run("7 x 5 = 35", func(t *testing.T) {
		obtive := Multiplique(7, 5)
		espero := 35

		if obtive != espero {
			t.Errorf("espero '%d', mas obtive '%d'", espero, obtive)
		}
	})

	t.Run("3 x 20 = 60", func(t *testing.T) {
		obtive := Multiplique(3, 20)
		espero := 60

		if obtive != espero {
			t.Errorf("espero '%d', mas obtive '%d'", espero, obtive)
		}
	})*/
}