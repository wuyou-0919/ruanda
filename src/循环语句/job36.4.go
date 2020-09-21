package main

func main() {
	var a,i,j,k int
		for a=100; a < 1000; a++{
			i=a/100
			j=a/10%10
			k=a%10
			if(i*100+j*10+k == i*i*i+j*j*j+k*k*k){
				println(a)
			}
		}
	}
