package piscine

/*tri a bulle*/
func SortIntegerTable(table []int) {
	for i := len(table) - 1; i >= 1; i-- {
		for j := 0; j <= i-1; j++ {
			if table[j+1] < table[j] {
				table[j+1], table[j] = table[j], table[j+1]
			}
		}
	}
}

func SortStringTable(table []string) {
	for i := len(table) - 1; i >= 1; i-- {
		for j := 0; j <= i-1; j++ {
			if table[j+1] < table[j] {
				table[j+1], table[j] = table[j], table[j+1]
			}
		}
	}
}

/*
func SortIntegerTable2(table []int) {
	min:=table[0]
	tri:=[]int{}
	for i:= range table{
	if table[i+1]<min{
		min=table[i+1]
	}
	tri=append(tri,min)
	table=table[]
	}



}
*/
