package main


//retunr the min in the variadics and return an error if we have an issue
import (
	"fmt"
)

func Min(values ...int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Error occured: The internal slice (values) is empty")
	}

	mini := values[0]
	
	for _, v := range values[1:] {
		if v < mini {
			mini = v
		}
	}
	return  mini, nil
}


func main() {
	res, err := Min(3, 4, 5, 2, 65, 3, 6, 3, 66, 78, 3 , 64, 46)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}
	fmt.Println(res)
}

