package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main()  {
	ctx := context.Background()


	ctx2 := context.WithValue(ctx, "Clave", "Valor")

	//Function(ctx2, 10)

	//fmt.Printf("%+v\n", ctx2)

	/*ctx3, _ := context.WithDeadline(ctx2, time.Now().Add(5 * time.Second))

	//esta esperando la lectura del canal
	<- ctx3.Done()

	fmt.Println(ctx3.Err().Error())*/

	//WithTimeout es la simplificacion de WithDeadline
	//ctx3, _ := context.WithTimeout(ctx2, 5 * time.Second)



	Function(ctx2, 10)

	<- ctx2.Done()
	

}

func Function(ctx context.Context, dato int)  {
	valor := ctx.Value("Clave")

	str, ok := valor.(string)
	if !ok {
		fmt.Println("No es string")
		os.Exit(1)
	}

	fmt.Printf("%T\n", str)

	_, cancel :=  context.WithTimeout(ctx, 5 * time.Second)

	// se dispara si se completa la operacion antes de tiempo y cancela el timeout
	defer cancel()
}