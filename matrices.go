package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func check(e error) {
	if e != nil {
	panic(e)
	}
}


func main() {
	//On stock dans la variable "name" le nom du fichier passé en argument
	name := os.Args[1]
	//Ouverture du fichier
	file, err := os.Open(name)
	check(err)

	//Création d'un reader
	rd := bufio.NewReader(file)
	var lien []string

	//On lit le fichier ligne à ligne auxquelles on enlève les caractères invisibles
	//On stocke chaque valeur dans un tableau lien
	for {
		lines, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		line := strings.TrimSuffix(lines, "\r\n")
		line = strings.TrimSuffix(line, "\n")

		lien = append(lien, strings.Split(line, " ")...)
	}

	//Conversion des valeurs "string" vers des "int"
	var tab []int
	for i:=0 ; i<len(lien) ; i++ {
		y,_ := strconv.Atoi(lien[i])
		tab = append(tab,y)
	}

	fmt.Printf("%v Tableau tab : \n" , tab)

	//x et y récupère le nombre de lignes et de colonnes des matrices
	x1 := tab[0]
	y1 := tab[1]
	x2 := tab[4+(x1*y1)]
	y2 := tab[5+(x1*y1)]

	//Condition pour multiplier deux matrices entre elles : le nombre de colonnes de A doit être égale au nombre de ligne de B
	if y1 != x2 {
		fmt.Println("ERREUR : Ces deux matrices ne sont pas multipliables entre elles")
	} else {
		//On remplit deux tableaux aves les valeurs correspondantes à chaque matrice
		var tabA []int
		for i := 3; i < 3+(x1*y1); i++ {
			tabA = append(tabA, tab[i])
		}
		var tabB []int
		for i := 7 + (x1 * y1); i < len(tab); i++ {
			tabB = append(tabB, tab[i])
		}

		fmt.Printf("%v Tableau tabA : \n", tabA)
		fmt.Printf("%v Tableau tabB : \n", tabB)

		//Création de la matrice A
		var matA [][]int
		matA = make([][]int, x1)
		for i := 0; i < x1; i++ {
			matA[i] = make([]int, y1)
		}
		//Création de la matrice B
		var matB [][]int
		matB = make([][]int, x2)
		for i := 0; i < x2; i++ {
			matB[i] = make([]int, y2)
		}

		//Remplissage de la matrice A
		cpt1 := 0
		cpt2 := 0
		for i := 0; i < x1; i++ {
			for j := 0; j < y1; j++ {
				matA[i][j] = tabA[cpt1]
				cpt1++
			}
		}
		//Remplissage de la matrice B
		for i := 0; i < x2; i++ {
			for j := 0; j < y2; j++ {
				matB[i][j] = tabB[cpt2]
				cpt2++
			}
		}

		fmt.Printf("%v Matrice A : \n", matA)
		fmt.Printf("%v Matrice B : \n", matB)

		//Déclaration de la matrice de résulat
		var res [][]int
		res = make([][]int, x1)
		for i := 0; i < x1; i++ {
			res[i] = make([]int, y2)
		}

		//Temps de calcul avec les goroutines
		debut := time.Now()
		time.Sleep(1 * time.Second)

		//Application des goroutines
		for i := 0; i < x1; i++ {
			for j := 0; j < y2; j++ {
				wg.Add(1)
				go mult(matA, matB, res, i, j)
			}
		}
		wg.Wait()
		fin := time.Since(debut)
		avecgr := fin
		fmt.Println("Matrice de résultat via les goroutines: ", res)
		fmt.Println("Temps de calcul avec les goroutines", fin.Nanoseconds())

		//Temps de calcul sans utilisation des goroutines
		debut = time.Now()
		time.Sleep(1 * time.Second)
		fmt.Println("Matrice de résultat sans goroutine : ", multMat(matA, matB, res, x1, y2, x2))
		fin = time.Since(debut)
		sansgr := fin
		fmt.Println("Temps de calcul sans les goroutines", fin.Nanoseconds())

		fmt.Println("Le temps gagné est de : ", sansgr-avecgr)
	}
}


//Fonction pour multiplier 2 matrices entre elles
func multMat (MatA [][]int, MatB [][]int,res [][]int, lig int, col int, mutu int) [][]int{
	for i:=0 ; i <lig ; i++{
		for j :=0 ; j <col ; j++{
			res[i][j] = 0
			for k :=0 ; k<mutu ; k++{
				res[i][j] = res[i][j] + MatA[i][k] * MatB[k][j]
			}
		}
	}
	return res
}


//Fonction pour multiplier ligne à colonne deux matrices -> Stratégie de goroutines
func mult(matA[][]int, matB[][] int, resu[][] int, i int, j int)  {
	defer wg.Done()
	res := 0
	counter := 0
	for counter < len(matB) {

		res = res + matA[i][counter]*matB[counter][j]

		counter++
	}
	resu[i][j] = res

}
