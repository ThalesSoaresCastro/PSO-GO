package main

import (
  "fmt"
  "math/rand"
  "time"
)

//função fitness
func fitness(input []float64) []float64{

  var arrayFitness []float64
  for i := 0; i < len(input); i++{
    arrayFitness = append(arrayFitness, 1+2*input[i]-(input[i]*input[i]))
  }

  return arrayFitness
}

//Acha a melhor posição até o momento da partícula
func pBestFunction(arr []float64, newPos []float64) []float64{

  var fitNewPos = fitness(newPos)
  var fitPBest = fitness(arr)

  var newPBest []float64

  fmt.Println("Arr Func: ", arr)

  fmt.Println("NewPos Func: ", newPos)

  //Se caso o valor do fitness da da posição da partícula calculada for melhor
  //do que o pBest, então o valor do pBest da partícula é atualizado para o novo
  //valor calculado da partícula.
  for i := 0; i< len(arr); i++ {
    if fitPBest[i] < fitNewPos[i]  {
        newPBest = append(newPBest ,newPos[i])
    }else{
      newPBest = append(newPBest ,arr[i])
    }
  }
  fmt.Println("pBest Func: ", newPBest)

  return newPBest
}

//Acha o melhor indivíduo do bando.
func gBestFunction(x []float64,fit []float64) float64{
  max := 0
  for i := 0; i<len(x) ; i++{
      if fit[max] < fit[i]{
        max = i
      }
  }
  return x[max]
}

/*
fi_1 = random(0.0 - 1.0) * constante_1 (c1)
fi_2 = random(0.0 - 1.0) * constante_2 (c2)

novaVelocidade  = (pBest - pAtual)* fi_1 + (gBest - pAtual) * fi_2
obs: pode ser somado a velocidade antiga junto com a nova ficando:


novaVelocidade = w*velocidadeAntiga +  (pBest - pAtual)* fi_1 + (gBest - pAtual) * fi_2
novaVelocidade = w*velocidadeAntiga + random_1(0.0 - 1.0)*constante_1*(pBest - pAtual) + random_2(0.0 - 1.0)*constante_2*(gBest - pAtual)

sendo: velocidadeAntiga = v_ant, constante_1 = c1, constante_2 = c2, random_1 = r1 e random_2 = r2, temos

novaVelocidade = w * v_ant + r1*c1*(pBest-pAtual) + r2 *c2*(gBest - pAtual)

posicaoNova = posicaoAntiga + novaVelocidade
*/

//Função para o cálculo das novas posições das partículas
func newPosition(oldPosition []float64, Velocity []float64) []float64{
    var newPos  []float64
    for i := 0;  i < len(oldPosition); i++ {
      newPos = append(newPos, oldPosition[i]+ Velocity[i])
    }
    return newPos
}


//Cálculo da nova posição do indivíduo
func newVelocity(w float64, c1 float64, c2 float64, gBest float64,
                 pAtual []float64, pBest []float64, oldVelocity []float64) []float64{

  var newVelocity []float64

  for i := 0;  i < len(pAtual); i++ {
    r1 := rand.Float64()
    r2 := rand.Float64()

    fi_1 := r1*c1
    fi_2 := r2*c2

    //Calcula a nova velocidade para cada indivíduo do array de Individuos
    newVelocity = append(newVelocity, w*oldVelocity[i] + (pBest[i] - pAtual[i])*fi_1 + (gBest - pAtual[i])*fi_2 )
  }

  return newVelocity
}

func main() {

  rand.Seed(time.Now().UTC().UnixNano())

  var n int
  var w, c1, c2, iteration float64

  var population []float64
  var velocity []float64
  var pBestArray []float64
  var newPBest []float64
  var velocityAux []float64
  var populationArray []float64

  gBestValue := 0.0

  fmt.Println("n, c1, c2, w, iteration: ")
  fmt.Scanln(&n, &c1, &c2 ,&w, &iteration)

  //primeira iteração da quantidade de épocas...
  for i :=0; i<n; i++{
    population = append(population, rand.Float64())
    velocity = append(velocity, rand.Float64())
  }

  //Na primeira iteração, o melhor valor para cada partícula é o valor Atual
  pBestArray = population

  //Cálculo do valor do gBest
  //indivíduo do bando que tenha o melhor valor de fitness da população
  gBestValue = gBestFunction(population,fitness(population))

  fmt.Println("População: ", population, "\nVelocidade:  ",velocity, "\npBest: ",pBestArray,"\ngBest: ", gBestValue,"\n\n")

  //Faz as próximas iterações
  for iteration > 0 {

    //Calculando a nova velocidade das partículas
    velocityAux = newVelocity(w, c1,c2,gBestValue, population, pBestArray, velocity)

    for i := 0; i<len(velocity); i++{
        velocity[i] = velocityAux[i]
    }

    //Calculando as novas posições das partículas
    populationArray = newPosition(population, velocity)
    for i:= 0; i < len(populationArray); i++ {
      population[i] = populationArray[i]
    }


    //Calculando o novo pBest
    newPBest = pBestFunction(pBestArray,population)
    for i:= 0;i< len(pBestArray);i++ {
        pBestArray[i] = newPBest[i]
    }


    //Calculo do novo gBest:
    //achando o novo valor do pBest com os novos valores de fitness
    gBestValue = gBestFunction(population,fitness(population))

    //--------------------------------------------------------------------------
    // Esvaziando os arrays para atualizar os dados...
    newPBest = append(newPBest[:0], newPBest[n:]...)
    velocityAux = append(velocityAux[:0], velocityAux[n:]...)
    populationArray = append(populationArray[:0], populationArray[n:]...)
    //--------------------------------------------------------------------------
    fmt.Println("População: ", population, "\nVelocidade:  ",velocity, "\npBest: ",pBestArray,"\ngBest: ", gBestValue,"\n\n")

    iteration--
  }

}
