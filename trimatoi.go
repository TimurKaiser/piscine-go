package piscine

func TrimAtoi(s string) int {
    x := false // x pour le signe négatif
    y := true  // chiffre ou pas
    r := 0
    for i, char := range s {
        if i == 0 && char == '-' {
            x = true // détecteur de -
        } else if char >= '0' && char <= '9' {
            r = r * 10           // la multiplication par 10 permet de déplacer avec un 0
            r = r + int(char-48) // -48 pour transformer un caractère en int dans l'ASCII
            y = false            // Si un chiffre est trouvé, y est faux (non vide)
        }
    }
    if y {
        return 0
    } else {
        if x {
            return -r // retourne le résultat négatif
        } else {
            return r // retourne le résultat positif
        }
    }
}
