package math

import (
	"github.com/chfenger/goNum"
	"math/rand"
)

func les_ECPE(a [][]float64, b []float64) ([]float64, bool) {
	/*
		线性代数方程组的列主元消去法
		输入   :
		    a       a x = b线性代数方程组的系数矩阵
		    b       a x = b线性代数方程组的右侧常数列向量
		输出   :
		    sol     解值
		    err     解出标志：false-未解出或达到步数上限；
		                     true-全部解出
	*/
	//方程个数为n
	var err bool = false
	atemp := a
	btemp := b
	n := len(btemp)
	sol := make([]float64, n)
	temp0 := make([]float64, n)
	var temp1 float64

	// 输入判断
	if len(atemp) != n {
		return sol, err
	}

	//求解
	//消去，求得上三角矩阵
	for true {
		for i := 0; i < n-1; i++ {
			//求第i列的主元素并调整顺序
			acol := make([]float64, n-i)
			for icol := i; icol < n; icol++ {
				acol[icol-i] = atemp[icol][i]
			}
			_, ii, _ := goNum.MaxAbs(acol)
			if ii+i != i {
				temp0 = atemp[ii+i]
				atemp[ii+i] = atemp[i]
				atemp[i] = temp0
				temp1 = btemp[ii+i]
				btemp[ii+i] = btemp[i]
				btemp[i] = temp1
			}

			//列消去
			for j := i + 1; j < n; j++ {
				mul := atemp[j][i] / atemp[i][i]
				for k := i; k < n; k++ {
					atemp[j][k] = atemp[j][k] - atemp[i][k]*mul
				}
				btemp[j] = btemp[j] - btemp[i]*mul
			}
		}

		//回代
		sol[n-1] = btemp[n-1] / atemp[n-1][n-1]
		for i := n - 2; i >= 0; i-- {
			temp1 = 0.0
			for j := i + 1; j < n; j++ {
				temp1 = temp1 + atemp[i][j]*sol[j]
			}
			sol[i] = (btemp[i] - temp1) / atemp[i][i]
		}
	}
	//返回结果
	err = true
	return sol, err
}

func Gaos() (err bool) {
	const M = 1000
	const N = 1000
	var (
		ma [][]float64
		mb []float64
	)
	for i := 0; i < M; i++ {
		tmp := make([]float64, N)
		for j := 0; j < N; j++ {
			tmp = append(tmp, rand.Float64())
		}
		ma = append(ma, tmp)
		mb = append(mb, rand.Float64())
	}
	mx, err := les_ECPE(ma, mb)
	if mx != nil {
	}
	return err
}
