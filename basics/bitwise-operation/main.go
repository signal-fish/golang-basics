package main

import "fmt"

/*
（1）二进制的最高位是符号位，0表示正数，1表示负数；
（2）正数的原码，反码，补码都一样；
     1 —> 原码[0000 0001]  反码[0000 0001]  补码[0000 0001]
（3）负数的反码：它的原码符号位不变，其它位取反（0 —> 1；1 —> 0）；
     1 —> 原码[1000 0001]  反码[1111 1110]  补码[1111 1111]
（4）负数的补码：它的反码加1；
（5）0的反码，补码都是0；
（6）在计算机运算的时候，都是以补码的方式来运算的。
     1 + 1；1 - 1 = 1 + (-1)


*/
func main() {
	fmt.Println("2&3 =", 2&3)
	fmt.Println("2|3 =", 2|3)
	fmt.Println("2^3 =", 2^3)
	/*
		正数的原码，反码，补码都一样，计算机运算以补码方式进行运算
		2的补码：0000 0010
		3的补码：0000 0011
		    2&3：0000 0010  —> 2
		    2|3: 0000 0011  —> 3
		    2^3: 0000 0001  —> 1
	*/

	fmt.Println("-2^2 =", -2^2)
	/*
		负数的反码等于原码符号位不变，其它位取反；负数的补码等于反码加1
		    -2的原码：1000 0010
		    -2的反码：1111 1101
		    -2的补码：1111 1110
		    2的补码： 0000 0010
		-2^2（补码）: 1111 1100
		-2^2（反码）: 1111 1011（补码减1）
		-2^2（原码）: 1000 0100   —> -4
	*/

	/*
			右移运算(>>)：低位溢出，符号位不变，并用符号位补溢出的高位
			               1 >>2：0000 0001
			            低位溢出：**00 0000
		            	  符号位不变：0*00 0000
			用符号位补溢出的高位：0000 0000  --> 0
			左移运算(<<)：符号位不变，低位补0
			      1<<2：0000 0001
			符号位不变：0000 01**
			   低位补0：0000 0100  --> 4
	*/
	fmt.Println("1>>2 =", 1>>2)
	fmt.Println("1<<2 =", 1<<2)
}
