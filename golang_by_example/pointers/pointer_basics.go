// pointers are special varibles, because they dont store the values, but they store the memory adress of those values

// For instance, when we have x := 1, here we mean that the value of x is 1

// However when we say write it this way var p *int = x
// here we are saying that the variable p is pointer, because it is not storing a value but the memory address of that value

// lets have a full explanantion

// we have:

// var p *int

// i := 42

// p = &i

// this means,

// i holds value 42
// however p hold the memory address of i that holds the value 42

// we have something called dereferencing:
// we use the word called *p, when we say we are dereferencing, we mean we are getting the value
// of the object *p.

// We do something like this,

// fmt.Println(*p) we are getting the value of i

// we can also change it through here

// *p = 20

// now the i will be 20 if we
// fmt.Println(i)