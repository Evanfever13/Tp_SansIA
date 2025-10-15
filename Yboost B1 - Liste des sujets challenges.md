# 🧩 Challenges de programmation sans IA


## 🎯 Objectif général
Apprendre à **raisonner, structurer et coder** des solutions sans l’aide d’un outil d’IA.  
Ces exercices visent à :
- Développer la rigueur algorithmique.  
- Mettre en pratique les concepts fondamentaux du langage Go.  
- Favoriser l’autonomie avant la collaboration avec l’IA.  


## 👨🏼‍💻 Liste des challenges

### Exercice 1 – Two Sum : Trouver deux nombres pour une somme cible

#### Instruction
On vous donne un tableau d’entiers et une valeur cible `target`.  
Votre mission est de trouver deux indices distincts `i` et `j` tels que :

```
nums[i] + nums[j] == target
```

- Si plusieurs paires existent, renvoyez la première trouvée.  
- Si aucune paire ne correspond, renvoyez `(-1, -1)`.

#### Signature de la fonction
```go
func TwoSum(nums []int, target int) (int, int)
```

#### Exemple de test
```go
func main(){
    nums := []int{2, 7, 11, 15}
    target := 9
    i, j := TwoSum(nums, target)
    fmt.Println(i, j) // Résultat attendu : 0 1
}
```
---

### Exercice 2 – Compression RLE : Réduire une chaîne de caractères

#### Instruction
On souhaite compresser une chaîne en appliquant une encodage RLE (Run-Length Encoding). Chaque caractère répété plusieurs fois est remplacé par le caractère suivi du nombre de répétitions.

Exemples :
```
"aaabbc" → "a3b2c1"
"ab"     → "a1b1"
```

Cela permet de réduire la taille d’une chaîne lorsqu’il y a beaucoup de répétitions successives.

#### Signature de la fonction
```go
func CompressRLE(s string) string
```

#### Exemple de test
```go
func main(){
    fmt.Println(CompressRLE("aaabbc")) // a3b2c1
    fmt.Println(CompressRLE("ab"))     // a1b1
}
```
---

### Exercice 3 – Kadane : Somme maximale d’un sous-tableau contigu

#### Instruction
On vous donne un tableau d’entiers pouvant contenir des valeurs positives et négatives. Votre objectif est de trouver la somme maximale d’un sous-tableau contigu.

Un sous-tableau contigu signifie, une portion de tableau composée d’éléments consécutifs, sans interruption.  

> Exemple pour `[5, -2, 3, 4, -1]`  
> - `[5, -2, 3]` est contigu  
> - `[5, 3, -1]` ne l’est pas (on saute des éléments)

Autrement dit, on cherche la portion continue qui donne la plus grande somme possible.

#### Signature de la fonction
```go
func MaxSubArray(nums []int) int
```

#### Exemple de test
```go
func main(){
    nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    fmt.Println(MaxSubArray(nums)) // 6 (le sous-tableau [4, -1, 2, 1])
}
```

#### Exemple de raisonnement étape par étape
| Étape | Élément | Somme courante | Somme max |
|--------|----------|----------------|------------|
| -2     | négatif  | -2             | -2         |
| +1     | repartir | 1              | 1          |
| -3     | somme    | -2             | 1          |
| +4     | repartir | 4              | 4          |
| +(-1)  | continuer| 3              | 4          |
| +2     | continuer| 5              | 5          |
| +1     | continuer| 6              | 6          |
...

---

### Exercice 4 – Pyramide centrée en `*`

#### Instruction
Vous devez afficher une **pyramide centrée** composée d’astérisques `*`, dont la hauteur est déterminée par un entier `n`.

Chaque étage de la pyramide doit être **aligné au centre**, de façon symétrique.

#### Signature de la fonction
```go
func PrintCenteredPyramid(n int)
```

####  Exemple de test
```go
func main(){
    PrintCenteredPyramid(4)
}

   *
  ***
 *****
*******
```
---

### Exercice 5 – Losange creux en `*`

#### Instruction
On souhaite afficher un losange creux composé d’astérisques `*`, centré et symétrique. Le rayon `r` détermine la hauteur totale du losange : `2*r + 1` lignes.

#### Signature de la fonction
```go
func PrintHollowDiamond(r int)
```

#### Exemple de test
```go
func main(){
    PrintHollowDiamond(2)
}

  *
 * *
*   *
 * *
  *
```
---

## Exercice 6 – Trouver le colis perdu durant le chargement

#### Instruction
Une plateforme logistique a perdu un colis confidentiel lors du chargement d’un train.  
Chaque wagon contient plusieurs colis, et les wagons sont reliés entre eux par un pointeur `NextWagon`.  

Vous devez parcourir cette structure chaînée pour retrouver le colis à partir de son numéro unique.

#### Structures de données
```go
type Package struct {
    Number    int
    Recipient string
    Address   string
    Content   string
    Weight    float64
}

type Wagon struct {
    Number    string
    Loading   []Package
    NextWagon *Wagon
}

type ResultSearch struct {
    NumberWagon     string
    PositionPackage int
    PackageFind     Package
}
```

#### Signature de la fonction
```go
func FindPackage(number int, startWagon *Wagon) ResultSearch
```

#### Exemple de test
```go
func main(){
    // Exemple de données
    w1 := &Wagon{
        Number: "A",
        Loading: []Package{{Number: 10}, {Number: 11}},
    }
    w2 := &Wagon{
        Number: "B",
        Loading: []Package{{Number: 12}, {Number: 13}, {Number: 42}},
    }

    w1.NextWagon = w2
    
    fmt.Println(FindPackage(42, w1))
    // Résultat attendu : Wagon B, position 2, package 42
}
```

---

## Exercice 7 – Trier les wagons du train par poids total

#### Instruction
Le train a été chargé sans ordre, créant un déséquilibre dangereux. Votre mission : réorganiser les wagons du plus léger au plus lourd, selon la somme des poids des colis qu’ils contiennent.

#### Structures de données
```go
type Package struct {
    Number    int
    Recipient string
    Address   string
    Content   string
    Weight    float64
}

type Wagon struct {
    Number    string
    Loading   []Package
    NextWagon *Wagon
}
```

#### Signature de la fonction
```go
func SortWagonsByTotalWeight(startWagon *Wagon) *Wagon
```

#### Exemple de teste
```go
func main(){
    // Exemple de données
    w1 := &Wagon{Number: "A", Loading: []Package{{Weight: 100}}}
    w2 := &Wagon{Number: "B", Loading: []Package{{Weight: 50}}}
    w3 := &Wagon{Number: "C", Loading: []Package{{Weight: 150}}}
    w1.NextWagon = w2
    w2.NextWagon = w3
    
    sorted := SortWagonsByTotalWeight(w1)
    // Résultat attendu : B -> A -> C
}
```
---

## 🧠 Synthèse globale

| Axe | Description |
|------|--------------|
| Algorithmique | Réfléchir avant d’écrire, structurer ses idées et évaluer la complexité. |
| Langage Go | Manipuler structs, slices, pointeurs, chaînes et fonctions. |
| Autonomie & posture | Travailler sans IA, comprendre ses erreurs, apprendre à reformuler. |
| Préparation IA | Ces exercices seront repris pour la séance "Développement collaboratif avec l’IA" (debug, refactor, documentation). |

## 🧭 Citation de clôture

*L’IA ne pense pas à ta place, mais elle t’aide à penser plus loin.*
