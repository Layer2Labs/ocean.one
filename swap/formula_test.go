package swap

import (
	"testing"

	"github.com/MixinNetwork/go-number"
	"github.com/stretchr/testify/assert"
)

func TestConstantProductFormula(t *testing.T) {
	assert := assert.New(t)

	cpf := &ConstantProductFormula{}

	xin := number.FromString("100")
	btc := number.FromString("2100")
	pool := &Pool{X: xin, Y: btc}

	in := &Input{number.FromString("210"), true}
	out, err := cpf.Swap(pool, in)
	assert.Nil(err)
	assert.Equal("9.09090909", out.Amount.String())
	assert.Equal("0.04761904", out.PriceInitial.String())
	assert.Equal("0.03935458", out.PriceFinal.String())
	assert.Equal("0.17355368", out.PriceSlippage.String())
	assert.Equal("90.90909091", pool.X.String())
	assert.Equal("2310", pool.Y.String())
	assert.Equal("210000.0000021", pool.X.Mul(pool.Y).String())

	in = &Input{number.FromString("10"), false}
	out, err = cpf.Swap(pool, in)
	assert.Equal("228.91891891", out.Amount.String())
	assert.Equal("100.90909091", pool.X.String())
	assert.Equal("2081.08108109", pool.Y.String())
	assert.Equal("210000.0000027918918919", pool.X.Mul(pool.Y).String())

	pool = &Pool{X: xin, Y: btc}
	out, err = cpf.Swap(pool, in)
	assert.Equal("190.9090909", out.Amount.String())
	assert.Equal("110", pool.X.String())
	assert.Equal("1909.0909091", pool.Y.String())
	assert.Equal("210000.000001", pool.X.Mul(pool.Y).String())
}
