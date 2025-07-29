package translator

import "testing" //paquete estaándar para pruebas

func TestRegisterTranslation(t *testing.T){
	// tests
	type RegisterTest struct {
		riscv string
		want string
	}

	// casos de prueba
	var testsCases = []RegisterTest{
		{"a0", "%rdi"},
		{"t1", "%r11"},
		{"sp", "%rsp"},
        {"ra", "%rax"},
        {"zero", "$0"},
        {"s5", ""},    // no mapeado
        {"foo", ""},   // clave invalida
	}

	for _, testCase := range testsCases {
		actual := GetX86Register(testCase.riscv)
		if actual != testCase.want {
			t.Errorf("GetX86Register(%q) = %q; se esperaba %q", testCase.riscv,
			actual, testCase.want)
		}
	}
}