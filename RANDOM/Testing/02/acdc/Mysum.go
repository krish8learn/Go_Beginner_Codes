//Package acdc here we write details of code for documentation
package acdc

//Sum this perform sum of variadic parameters
func Sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
